# GO Auth Server

## Description

This is a simple authentication server written in Go. It uses JWT for authentication and Postgres for storage.

## Usage

### Environment Variables

| Name          | Description                      | Default     |
| ------------- | -------------------------------- | ----------- |
| `PORT`        | Port to listen on                | `8080`      |
| `DB_HOST`     | Hostname of the Postgres server  | `localhost` |
| `DB_PORT`     | Port of the Postgres server      | `5432`      |
| `DB_USER`     | Username for the Postgres server | `postgres`  |
| `DB_PASSWORD` | Password for the Postgres server | `postgres`  |
| `DB_NAME`     | Name of the Postgres database    | `postgres`  |
