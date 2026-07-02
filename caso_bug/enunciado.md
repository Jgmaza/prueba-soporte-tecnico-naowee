# Caso 2 — Corregir un bug: "no hay grupos disponibles al cambiar de deporte"

## Contexto del dominio (lo mínimo para entender)

En un evento, cada deportista pertenece a una **categoría** que depende de su **año de nacimiento**:

| Categoría | Años de nacimiento |
|---|---|
| `INF` (Infantil) | 2015–2017 |
| `PRE_JUV` (Pre-juvenil) | 2012–2014 |
| `JUV` (Juvenil) | 2009–2011 |

Las inscripciones se hacen en **grupos**. Cada grupo pertenece a un **deporte** y a una
**categoría** (p.ej. "grupo de BALONCESTO categoría JUV").

Existe una función para **cambiar a un deportista de deporte** (`sport-swap`): dado el deporte
destino, busca en qué grupos de ese deporte puede quedar inscrito.

## El ticket

> *"Tenemos una deportista que quedó mal, está en una categoría que no le corresponde para la
> edad (nació en 2011 pero aparece en PRE_JUV). Queremos moverla de deporte pero el sistema dice
> que 'no hay grupos disponibles', aunque sí existe el grupo correcto para su edad en el deporte
> al que la queremos pasar. ¿Pueden revisar?"*

Este es un **bug real**: hay deportistas cuya categoría guardada quedó **inconsistente** con su
edad (por ejemplo, por migraciones de datos). Cuando se intenta cambiarlos de deporte, el sistema
no encuentra grupos aunque el grupo correcto para su edad exista.

## Tu tarea

En este directorio hay un módulo de Go con sus pruebas.

1. Corre las pruebas: `go test ./...`
2. **Encuentra la causa raíz del bug** y **arréglalo** para que todas las pruebas pasen.
3. Abre un **Pull Request** con tu fix.

## Qué esperamos en el Pull Request

La descripción del PR es tan importante como el código. Incluye:
- **Causa raíz:** ¿por qué fallaba? (en tus palabras)
- **El fix:** qué cambiaste y por qué resuelve el problema de fondo.
- **Verificación:** cómo confirmaste que funciona y que no rompiste los demás casos.

Tiempo esperado: ~1 hora.
