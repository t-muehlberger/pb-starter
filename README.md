# PB-Starter

My personal [PocketBase](https://pocketbase.io/) starter project

**Features:**

- PocketBase boilerplate
- Docker setup
- GitHub Actions build
- Live reload vi [air](https://github.com/air-verse/air)

## Run

```shell
go run ./cmd/server serve
```

## GitHub Actions

`Settings` => `Actions` => `General` => `Workflow permissions` => Set to `Read and write`

## Before using

**Update dependencies**

```shell
go get -u ./...
go mod tidy
```
