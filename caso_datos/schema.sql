-- Esquema simplificado para el Caso 1.
-- Compatible con PostgreSQL y SQLite.

-- === "Base de datos" de usuarios (microservicio de identidad) ===
CREATE TABLE users (
    code            TEXT PRIMARY KEY,          -- código interno del usuario
    document_type   TEXT NOT NULL,             -- p.ej. 'CC', 'TI', 'CE'
    document_number TEXT NOT NULL,
    full_name       TEXT NOT NULL,
    email           TEXT NOT NULL,
    created_at      TEXT NOT NULL,
    -- La regla de negocio clave: el documento es único globalmente.
    UNIQUE (document_type, document_number)
);

-- === "Base de datos" de inscripciones (microservicio de eventos) ===
-- OJO: se relaciona con users por user_code, pero NO hay FOREIGN KEY:
-- son servicios/BD distintos en la realidad.
CREATE TABLE participants (
    code        TEXT PRIMARY KEY,
    user_code   TEXT NOT NULL,                 -- -> users.code (sin FK real)
    event_code  TEXT NOT NULL,
    status      TEXT NOT NULL,                 -- 'ENROLLED', 'WITHDRAWN', etc.
    created_at  TEXT NOT NULL
);
