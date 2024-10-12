# hello-world-go
This project is a Golang web application that returns 'Hello, ${hostname}!' and builds a Docker image for testing behind an API Gateway.

## Supported tags and respective Dockerfile links
- [`latest` `1.0.1` `1.0.1-alpine`](https://github.com/zazayaya/hello-world-go)
- [`1.0.0`](https://github.com/zazayaya/hello-world-go)

## Test
```shell
docker run --name hello-world-go -d -p 0.0.0.0:8080:8080 zazayaya/hello-world-go
docker ps | grep hello-world-go
curl http://localhost:8080
```

## Quick reference
- Where to file issues: https://github.com/zazayaya/hello-world-go
