-- 1. Primero eliminamos las conciliaciones que apuntan a esas facturas
DELETE FROM conciliations;

-- 2. Ahora sí podemos borrar las facturas sin violar el constraint
DELETE FROM invoices;

-- 3. Insertamos los nuevos seeds con soporte multimoneda (2026)
INSERT INTO invoices (id, number, provider, amount, currency, status, date, due_date) VALUES
(gen_random_uuid(), 'FAC-MXN-001', 'Proveedor Local MX', 5000, 'MXN', 'PENDING', NOW(), NOW()),
(gen_random_uuid(), 'FAC-USD-001', 'Servicios Globales', 150, 'USD', 'PENDING', NOW(), NOW()),
(gen_random_uuid(), 'FAC-TEST-003', 'Auditoría Externa', 888, 'COP', 'PENDING', NOW(), NOW());