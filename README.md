# Golang Startup

Golang basic web application structure using an onion architecture approach.

## 1. Config

Before to run application you may setup ``.env`` with ``Config`` values.
Check out the following tables to know all ``Config`` parameters detailed.

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``DB_ADDRESS`` | Database URL address. | `string`  | ` ` | **YES** |
| ``DB_NAME`` | Name of the initial database to access. | `string`  | ` ` | **NO** |
| ``MIGRATIONS_PATH`` | Migrations local path if needed. | `string`  | ` ` | **NO** |
| ``HOST`` | HTTP server host number. | `string` | ` ` | **NO** |
| ``PORT`` | HTTP server port number. | `string` | ` ` | **YES** |

## 2. Run Application

After define ``env`` variables values, run the following commands to start web application server.

The different examples of running this application packages and features are available in ``cmd``.

For a REST application with a postgres database layer use:

``
go run ./cmd/postgres/main.go
``

For a mongoDB database layer use:

``
go run ./cmd/mongo/main.go
``
