# Golang Basic REST API

Golang basic REST API is a web application structure using an onion architecture approach.

It implements controller, service, database layers, that makes above layers independent of which database (or any other
bottom layer type e.g. other API downstream) type is being used.
There are ``cmd`` implemented examples of the application, for each database type: ``mongo``, ``mysql`` and ``postgres``.

It also could be a useful structure for Golang development, having model, config loading, implemented GitHub workflows,
Makefile, that are often requirements of a Golang application.

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

For a REST application with MySQL database layer use:

``
go run ./cmd/mysql/main.go
``

For a REST application with MongoDB database layer use:

``
go run ./cmd/mongo/main.go
``

## 3. Swagger specifications


### 3.1. Install swagger in Unix

In order to install ``swagger`` binary in a Debian Linux distribution, use the following code:

```
sudo apt update
sudo apt install -y apt-transport-https gnupg curl
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | sudo apt-key add -
curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' | sudo tee /etc/apt/sources.list.d/go-swagger-go-swagger.list
sudo apt update 
sudo apt install swagger
```

Source: 🔗 [Go Swagger](https://goswagger.io/install.html)

## 3.2. Run swagger command in Windows

If you're using windows use the following command before swagger code:

``docker run --rm -it --env GOPATH=/go -v %CD%:/go/src -w /go/src quay.io/goswagger/swagger <swagger code here>``
