# hello-world-go
This project is a 'Hello, World!' in Golang Web Applications to build a docker image we can test behind an API Gateway

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
curl http://IP:8080/
```