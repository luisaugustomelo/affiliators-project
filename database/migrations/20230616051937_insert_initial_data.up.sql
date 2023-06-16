-- Up migration script
BEGIN;
INSERT INTO hubla.sales_type (id, description, kind, signal) VALUES
    (1, 'Venda produtor', 'Entrada', '+'),
    (2, 'Venda afiliado', 'Entrada', '+'),
    (3, 'Comissão paga', 'Saída', '-'),
    (4, 'Comissão recebida', 'Entrada', '+');
COMMIT;
