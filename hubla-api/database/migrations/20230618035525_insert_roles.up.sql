-- Up migration script
BEGIN;
INSERT INTO roles (id, description, created_at, updated_at) VALUES
        (1, 'creator', NOW(), NOW()),
        (2, 'affiliator', NOW(), NOW());
COMMIT;