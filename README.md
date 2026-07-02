# Prueba técnica — Desarrollador de Soporte Técnico

Esta prueba consta de dos casos inspirados en situaciones reales del trabajo de soporte en
Naowee. El objetivo es evaluar cómo analizas un problema y llegas a una solución, no que
memorices sintaxis ni construyas una aplicación completa.

## Tiempo esperado

**~2 horas en total**, aprox. 1h por caso. Valoramos más la calidad del razonamiento que la velocidad.

## Reglas

- **Puedes usar IA (ChatGPT, Claude, Copilot, etc.), Google y cualquier recurso.** No evaluamos
  sintaxis de memoria, sino criterio: cómo entiendes el problema, la calidad de la solución, y
  el uso apropiado de la IA (saber qué pedirle, validar lo que devuelve y no aceptarlo a ciegas).
- Habrá una **conversación presencial** donde te pediremos explicar tu entrega y hacer un ajuste
  en vivo. Entrega solo lo que puedas sostener y explicar.
- Si el enunciado no es claro en algún punto, asume lo razonable y déjalo documentado.

## Los dos casos

| Caso | Carpeta | Entrega |
|---|---|---|
| **1. Ajuste de datos** | [`caso_datos/`](caso_datos/) | Un `SOLUCION.md` con diagnóstico y SQL |
| **2. Corrección de un bug** | [`caso_bug/`](caso_bug/) | Un **Pull Request** con el fix |

Lee el `enunciado.md` dentro de cada carpeta.

## Qué evaluamos (en ambos casos)

1. **Análisis:** identificar la causa raíz, no solo el síntoma.
2. **Solución:** que la corrección resuelva el problema de fondo.
3. **Prudencia y verificación:** cómo compruebas que el cambio es correcto y no rompe nada. En
   soporte se tocan datos y sistemas reales; el cuidado importa tanto como acertar.
4. **Comunicación:** explicar el problema y la solución con claridad.
5. **Manejo básico de Git:** crear el repositorio, commits con mensajes claros y un Pull Request bien armado.

## Entrega

La entrega la haces en **un repositorio que tú creas** (GitHub, GitLab, etc.). No nos envíes un
comprimido.

1. Crea un repositorio con la estructura de la prueba (`caso_datos/` y `caso_bug/`).
2. **Caso 1:** agrega `caso_datos/SOLUCION.md` con tu diagnóstico y SQL.
3. **Caso 2:** aplica el fix en el módulo de Go y ábrelo como un **Pull Request** dentro de tu repositorio.
4. Compártenos el enlace a tu repositorio y al Pull Request.
