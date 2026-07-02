-- Datos de ejemplo para el Caso 1.
-- Todo es ficticio.

-- ------------------------------------------------------------------
-- USUARIOS
-- ------------------------------------------------------------------
INSERT INTO users (code, document_type, document_number, full_name, email, created_at) VALUES
  ('U-1001', 'CC', '1052000123',  'Mariana Gomez',        'mariana.gomez@correo.com',      '2021-03-10T09:00:00Z'),
  ('U-1002', 'CC', '10520001230', 'Mariana Gomez',        'mariana.gomez.dep@correo.com',  '2026-05-28T14:22:00Z'),
  ('U-1003', 'CC', '1052000124',  'Andres Torres',        'a.torres@correo.com',           '2024-01-15T08:30:00Z'),
  ('U-1004', 'TI', '1087752127',  'Carlos Burgos',        'c.burgos@correo.com',           '2025-02-01T10:05:00Z'),
  ('U-1005', 'CC', '43555111',    'Laura Marin',          'laura.marin@correo.com',        '2023-11-20T16:45:00Z'),
  ('U-1006', 'CC', '01052000123', 'Pedro Ramirez',        'pedro.ramirez@correo.com',      '2022-07-07T11:15:00Z'),
  ('U-1007', 'CC', '52001200',    'Sofia Herrera',        'sofia.herrera@correo.com',      '2024-09-02T13:00:00Z'),
  ('U-1008', 'CE', '99887766',    'Diego Castaño',        'diego.castano@correo.com',      '2025-06-18T09:40:00Z'),
  ('U-1009', 'CC', '1052999888',  'Valentina Rojas',      'v.rojas@correo.com',            '2023-04-11T17:20:00Z');

-- ------------------------------------------------------------------
-- INSCRIPCIONES (participants)
-- Recuerda: se cruza por user_code, NO hay FK.
-- ------------------------------------------------------------------
INSERT INTO participants (code, user_code, event_code, status, created_at) VALUES
  ('P-5001', 'U-1002', 'EVT-JIN-2026', 'ENROLLED',  '2026-05-28T14:25:00Z'),
  ('P-5002', 'U-1003', 'EVT-JIN-2026', 'ENROLLED',  '2026-05-29T09:10:00Z'),
  ('P-5003', 'U-1004', 'EVT-JIN-2026', 'ENROLLED',  '2026-05-30T10:00:00Z'),
  ('P-5004', 'U-1005', 'EVT-JIN-2026', 'ENROLLED',  '2026-05-30T11:30:00Z'),
  ('P-5005', 'U-1006', 'EVT-JIN-2026', 'ENROLLED',  '2026-06-01T08:00:00Z'),
  ('P-5006', 'U-1007', 'EVT-JIN-2026', 'WITHDRAWN', '2026-06-02T15:00:00Z');
