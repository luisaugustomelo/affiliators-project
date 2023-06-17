-- Down migration script
BEGIN;
DELETE FROM sale_types WHERE id IN (1, 2, 3, 4);
COMMIT;