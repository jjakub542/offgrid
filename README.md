# Project offgrid

Off-Grid project

## How to run locally

1. Install Go 1.22
2. Install and configure PostgreSQL (create user and db for project, grant permissions to user)
3. Clone repository
4. Add .env with PORT, APP_ENV, DB_USERNAME, DB_PASSWORD, DB_PORT, DB_HOST, DB_NAME
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
```bash
sudo -u postgres psql
```

Then:
```sql
CREATE DATABASE offgrid_db;
CREATE DATABASE offgrid_db_test;
CREATE USER offgrid_admin WITH PASSWORD 'offgrid123';
GRANT ALL PRIVILEGES ON DATABASE offgrid_db TO offgrid_admin;
GRANT ALL PRIVILEGES ON DATABASE offgrid_db_test TO offgrid_admin;
```


Connect to offgrid_db as offgrid_admin:
```bash
psql -h localhost -d offgrid_db -U offgrid_admin -p 5432
```
