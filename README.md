# hello-world-go
This project is a 'Hello, World!' in Golang Web Applications to build a docker image we can test behind an API Gateway

## Supported tags and respective Dockerfile links
- [`latest` `1.0.1` `1.0.1-alpine`](https://github.com/zazayaya/hello-world-go)
- [`1.0.0`](https://github.com/zazayaya/hello-world-go)

## Build hello-world
```shell
SET GOOS=linux
go build -o build/package/hello-world src/hello-world.go
chmod +x hello-world
```

## Build Dockerfile
```shell
docker build -t zazayaya/hello-world-go:latest .
docker build -t zazayaya/hello-world-go:1.0.0 .
```

## Test
```shell
docker run --name go-hello-world -d -p 0.0.0.0:8080:8080 zazayaya/go-hello-world
docker ps | grep go-hello-world
curl http://localhost:8080
```

## Quick reference
- Where to file issues: https://github.com/zazayaya/hello-world-go