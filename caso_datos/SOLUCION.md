# Caso 1 — Ajuste de datos: "El documento ya está registrado"

## Diagnóstico (causa raíz)

El error **no es un bug** del sistema de unicidad: el constraint `UNIQUE (document_type, document_number)` está funcionando correctamente.

La causa raíz es un **registro duplicado de la misma persona** en la tabla `users`:

| Código   | Documento      | Nombre        | Creado     | Rol en el incidente |
|----------|----------------|---------------|------------|---------------------|
| `U-1001` | CC 1052000123  | Mariana Gomez | 2021-03-10 | Usuario histórico con el documento **correcto** |
| `U-1002` | CC 10520001230 | Mariana Gomez | 2026-05-28 | Registro nuevo del profe con un **typo** (cero extra al final) |

El profesor intentó corregir el documento de `U-1002` a `CC 1052000123`, pero ese documento **ya pertenece a `U-1001`**. Por eso el servicio responde `DOCUMENT_NUMBER_ALREADY_REGISTERED (409)`.

Además, la inscripción al evento (`P-5001` en `EVT-JIN-2026`) quedó asociada a `U-1002` (el usuario con el documento incorrecto), no a `U-1001` (el usuario con el documento correcto).

El profesor tiene razón desde su perspectiva ("solo la registré una vez"): él creó `U-1002` en 2026, pero desconoce el registro histórico `U-1001` de 2021.

---

## Consultas SQL de investigación

### 1. Usuario del error (del log)

```sql
SELECT code, document_type, document_number, full_name, email, created_at
FROM users
WHERE code = 'U-1002';
```

**Resultado esperado:** `U-1002` con documento `CC 10520001230` (typo).

### 2. ¿Quién ya tiene el documento correcto?

```sql
SELECT code, document_type, document_number, full_name, email, created_at
FROM users
WHERE document_type = 'CC'
  AND document_number = '1052000123';
```

**Resultado esperado:** `U-1001` con el documento que el profe intenta asignar.

### 3. ¿Hay más registros de Mariana Gomez?

```sql
SELECT code, document_type, document_number, full_name, email, created_at
FROM users
WHERE full_name LIKE '%Mariana%';
```

**Resultado esperado:** dos filas — `U-1001` y `U-1002`.

### 4. ¿A qué usuario está ligada la inscripción?

```sql
SELECT p.code, p.user_code, p.event_code, p.status, p.created_at,
       u.full_name, u.document_type, u.document_number
FROM participants p
JOIN users u ON p.user_code = u.code
WHERE u.full_name LIKE '%Mariana%';
```

**Resultado esperado:** `P-5001` apunta a `U-1002` (documento con typo), no a `U-1001`.

### 5. ¿`U-1002` tiene otras referencias?

```sql
SELECT *
FROM participants
WHERE user_code = 'U-1002';
```

**Resultado esperado:** solo `P-5001`. Necesario confirmar en producción que no hay referencias en otros microservicios.

---

## SQL de corrección propuesto

**Enfoque:** reasignar la inscripción al usuario canónico (`U-1001`, que ya tiene el documento correcto) y eliminar el duplicado (`U-1002`).

> Ejecutar dentro de una transacción. En producción, hacer backup o snapshot antes.

```sql
BEGIN;

-- 1. Reasignar la inscripción al usuario con el documento correcto
UPDATE participants
SET user_code = 'U-1001'
WHERE code = 'P-5001'
  AND user_code = 'U-1002';

-- 2. Eliminar el usuario duplicado (solo si no quedan referencias)
DELETE FROM users
WHERE code = 'U-1002';

COMMIT;
```

### Consultas de verificación post-corrección

```sql
-- Solo debe quedar un Mariana Gomez
SELECT code, document_type, document_number, full_name
FROM users
WHERE full_name LIKE '%Mariana%';

-- La inscripción debe apuntar a U-1001 con documento correcto
SELECT p.code, p.user_code, p.status, u.document_number
FROM participants p
JOIN users u ON p.user_code = u.code
WHERE p.code = 'P-5001';

-- No debe haber documentos duplicados
SELECT document_type, document_number, COUNT(*) AS cnt
FROM users
GROUP BY document_type, document_number
HAVING COUNT(*) > 1;
```

**Resultado esperado tras la corrección:**
- Un solo usuario Mariana (`U-1001`) con `CC 1052000123`
- `P-5001` inscrita bajo `U-1001` con estado `ENROLLED`
- Sin violaciones de unicidad

---

## Verificaciones previas a producción

Antes de ejecutar en producción, verificaría:

1. **Identidad:** confirmar con el profe o Servicio al Cliente que `U-1001` y `U-1002` son la misma persona (mismo nombre, contexto del evento, posible registro previo en 2021).
2. **Referencias cruzadas:** en otros microservicios (pagos, historial deportivo, etc.) buscar referencias a `U-1002` que no estén en `participants`.
3. **Historial de `U-1001`:** revisar si tiene inscripciones activas en otros eventos que pudieran verse afectadas.
4. **Backup:** snapshot o export de las filas afectadas (`users` U-1001/U-1002, `participants` P-5001) antes del cambio.
5. **Estado antes/después:** ejecutar las consultas de investigación y verificación, guardar resultados.

---

## Riesgos

| Riesgo | Mitigación |
|--------|------------|
| `U-1001` y `U-1002` son personas distintas con el mismo nombre | Confirmar identidad antes de fusionar; no ejecutar sin certeza |
| Otras tablas/servicios referencian `U-1002` | Auditar referencias en todos los microservicios antes del DELETE |
| Pérdida de trazabilidad del registro duplicado | Guardar backup del registro `U-1002` antes de eliminarlo |
| Condición de carrera si el profe reintenta mientras se corrige | Coordinar ventana de mantenimiento o bloquear el registro temporalmente |

---

## Cuándo NO ejecutaría solo y escalaría al líder

- Si no puedo confirmar con certeza que `U-1001` y `U-1002` son la misma persona.
- Si `U-1002` tiene referencias en otros sistemas que no puedo actualizar desde soporte.
- Si `U-1001` tiene historial activo relevante (otros eventos, pagos, certificados) y la fusión podría afectarlo.
- Si el volumen de duplicados es mayor al esperado (posible problema sistémico, no un caso aislado).
- Si no tengo permisos para DELETE en `users` o para modificar inscripciones en producción.

---

## Nota sobre el enfoque descartado

**No recomiendo** cambiar el documento de `U-1001` para "liberar" el número y luego corregir `U-1002`: implicaría alterar el registro que **ya tiene el dato correcto**, con más riesgo y sin beneficio. La corrección natural es consolidar la inscripción en el usuario canónico.
