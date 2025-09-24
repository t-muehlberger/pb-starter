# PB-Starter

My personal [PocketBase](https://pocketbase.io/) starter project

**Features:**

- PocketBase boilerplate
- Basic client template
- Docker setup
- GitHub Actions build
- Live reload via [air](https://github.com/air-verse/air)

## Run

```shell
go run ./cmd/server serve
```

Run with live reload

```shell
go run github.com/air-verse/air@latest serve --dev
```

## GitHub Actions

`Settings` => `Actions` => `General` => `Workflow permissions` => Set to `Read and write`

## Before using

**Update dependencies**

```shell
go get -u ./...
go mod tidy
```

Make sure [`modernc.org/libc`](https://pkg.go.dev/modernc.org/libc) version matches [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite) if *modernc* Sqlite-Driver is used.

## Client

Basic hello world template at `client/` using Vite + PocketBase SDK.

**Development:**
```shell
cd client && npm install
npm run start    # dev server at :5173
```

**Production:**
```shell
cd client && npm run build    # builds to client/dist/
# pb_public/ symlinks to client/dist/ for PocketBase serving
```
