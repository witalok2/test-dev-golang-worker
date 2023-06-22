# test-dev-golang-worker

## Structure

```
├── cmd
│   └── main.go
├── config
│   ├── .env
│   ├── .env-example
│   └── config.go
├── deployment
│   ├── cloud-build.yaml
│   └── Dockerfile.go
├── internal
│   ├── entity
│   │   └── entity.go
│   ├── gateway
│   │   └── rabbimq.go
│   ├── repository
|   │   ├── postgres.go
|   │   └── sql.go
│   ├── service
│   │   ├── client.go
│   │   └── service.go
├── go.mod
├── go.sum
```

run worker
```
make up
```
run test
```
make test
```