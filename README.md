# Estimation Sheet

This microservice is a proof of concept for estimation sheet creation

# Configuration files

Create .env file

```bash
# Server
PORT=9000

# Database
DB_CONNECTION=postgres://postgres:postgres@db:5432/postgres?sslmode=disable&search_path=estimation
```

Create .env.test

```bash
# Server
PORT=9000

# Database
DB_CONNECTION=postgres://test:test@db:5432/test?sslmode=disable&search_path=estimation

```

# End-to-end testing

Using the command `make build` you can build the binary within the folder `bin/`.
Create a .env file within this folder as well, with test configuration.

```bash
# Server
PORT=9000

# Database
DB_CONNECTION=postgres://test:test@db:5432/test?sslmode=disable&search_path=estimation

```
After that, you can the command `make test-e2e`.