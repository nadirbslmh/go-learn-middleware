# go-learn-middleware

REST API Application with logging, authentication and rate limiter middleware. Written in Go with Echo Framework.

## How to Use

1. Clone this repository.

2. Create a new database.

```sql
CREATE DATABASE learn_middleware;
```

3. Create `.env` file to store the database configurations.

```sh
cp .env.example .env
```

4. Fill the database configuration in `.env` file.

5. Run the application.

```sh
go run main.go
```
