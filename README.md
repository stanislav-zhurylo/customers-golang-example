## Preparing the database

In order to start the software, ensure that postgres service is up and running.

Database connection setup can be prepared via environmental variables:

* DB_USER - default value is "postgres"
* DB_PASS - default value is "changeme"
* DB_NAME - default value is "golang"
* DB_HOST - default value is "localhost"
* DB_PORT - default value is "5432"

## Migration

Database migration will be performed during deployment automatically. By default I have inserted 21 demo customers for demo purposes.

## 1. Updating dependencies

In order to fetch all dependencies, please execute the following command:

```
go get -u
```

## 2. Starting the web service

In order to start the web service, please execute the following command:

```
go run app.go
```