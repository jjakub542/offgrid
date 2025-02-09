# Project offgrid

Off-Grid project

## How to run locally

1. Install Go 1.22
2. Install and configure PostgreSQL (offgrid_admin, offgrid_db - grant all permissions to offgrid_admin)
3. Clone repo
4. Add .env with PORT, APP_ENV, DB_USERNAME=offgrid_admin, DB_PASSWORD, DB_PORT=5432, DB_HOST=localhost, DB_NAME=offgrid_db
5. Run with commands below:

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```

## Useful commands:

Connect to psql as postgres:
sudo -u postgres psql

Connect to offgrid_db as offgrid_admin:
psql -h localhost -d offgrid_db -U offgrid_admin -p 5432
Password for user offgrid_admin: offgrid123
