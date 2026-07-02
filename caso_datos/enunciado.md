# Caso 1 — Ajuste de datos: "el documento ya está registrado"

## Contexto

En nuestra plataforma, cada persona (deportista, entrenador, etc.) es un **usuario**. La
combinación `(tipo_documento, numero_documento)` es **única**: no pueden existir dos usuarios
con el mismo documento.

Por otro lado, un usuario puede estar **inscrito** en un evento deportivo o no. La inscripción
vive en una tabla aparte (`participantes`) que se relaciona con el usuario por su `codigo`.
**No hay llave foránea** entre ambas tablas: se cruzan por `usuario_codigo`.

> En la vida real son dos bases de datos de microservicios distintos. Para esta prueba las
> tenemos juntas en un mismo esquema para que sea fácil de correr, pero **trátalas como si no
> pudieras hacer un `JOIN` con integridad referencial garantizada**.

## El ticket

Nos llega esto de Servicio al Cliente (tal cual, sin pulir):

> *"El profe dice que se equivocó al registrar a un deportista y le puso mal el documento.
> Cuando intenta corregirlo la plataforma le tira un error de que el documento ya existe. Pero
> él dice que ese deportista solo lo registró una vez. Ayúdenme porfa que el torneo arranca ya."*

En los logs del servicio encontramos esto: [`error.log`](error.log).

El deportista afectado se llama **Mariana Gómez** y su documento **correcto** es
`CC 1052000123`.

## Tu tarea

1. **Levanta los datos.** Tienes [`schema.sql`](schema.sql) y [`seed.sql`](seed.sql). Puedes
   usar PostgreSQL, SQLite, o el motor que prefieras (el SQL es estándar).
2. **Diagnostica la causa raíz.** ¿Por qué la plataforma dice que el documento ya existe si el
   profe registró a Mariana una sola vez? Escribe las consultas que usaste para investigarlo.
3. **Propón la solución:** el/los comandos SQL para dejar el dato correcto, de modo que el profe
   pueda corregir el documento de Mariana.
4. **Antes de "ejecutar":** explica **qué verificarías antes** de correr tu solución en
   producción, qué riesgo tiene, y **en qué caso NO la ejecutarías tú solo** y escalarías a tu
   líder.

## Entrega

Un archivo **`SOLUCION.md`** en esta carpeta con:
- Tu diagnóstico (causa raíz, en tus palabras).
- Las consultas SQL de investigación.
- El SQL de corrección propuesto.
- Tus verificaciones previas y consideraciones de riesgo.

Tiempo esperado: ~1 hora.
