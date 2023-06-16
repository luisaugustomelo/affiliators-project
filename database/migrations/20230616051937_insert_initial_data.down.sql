-- Down migration script
BEGIN;
DELETE FROM hubla.sales_type WHERE id IN (1, 2, 3, 4);
COMMIT;