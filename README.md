# Golang Startup

Golang basic web application structure using an onion architecture approach.

## 1. Config

Before to run application you may setup ``.env`` with ``Config`` values.
Check out the following tables to know all ``Config`` parameters detailed.

| Parameter | Description | Type | Default | Required |
|:---|:---|:---|:---|:---|
| ``DATABASE`` | Database URL address. | `string`  | ` ` | **YES** |
| ``HOST`` | HTTP server host number. | `string` | ` ` | **NO** |
| ``PORT`` | HTTP server port number. | `string` | ` ` | **YES** |

## 2. Run Application

After define ``env`` variables values, run the following commands to start web application server.

``
go run ./cmd/server/main.go
``
