-- Up migration script
BEGIN;
INSERT INTO sale_types (id, description, kind, signal, created_at, updated_at) VALUES
    (1, 'Venda produtor', 'Entrada', '+', NOW(), NOW()),
    (2, 'Venda afiliado', 'Entrada', '+', NOW(), NOW()),
    (3, 'Comissão paga', 'Saída', '-', NOW(), NOW()),
    (4, 'Comissão recebida', 'Entrada', '+', NOW(), NOW());
COMMIT;
