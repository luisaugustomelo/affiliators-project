# hubla-challenge

### Makefile

Sure, here are the descriptions for these commands in English:

- `make build` will compile the Go project and create an executable named myapp.
- `make test` will run all tests in the project.
- `make clean` will clean up compiled files from the project.
- `make run` will compile and run the application.
- `make up` will start up all services defined in the Docker Compose file.
- `make down` will stop and remove the containers defined in the Docker Compose file.
- `make migration-up` will start up all tables defined in the Docker Compose file.
- `make migration-down` will stop and remove the tables defined in the Docker Compose file.

## Migrations

The database access credentials for local access are defined inside the Makefile in the following variables. If it is necessary to modify these variables, make the changes inside the file.

```bash
DB_HOST=localhost
DB_PORT=5432
DB_NAME=hubla
MIGRATIONS_PATH=database/migrations
```

All application migrations are contained in database/migrations and should be executed using the `make migrate-up` and `make migrate-down` commands.
