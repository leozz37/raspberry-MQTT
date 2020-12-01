# Raspberry MQTT Project

Raspberry MQTT example

## Controle

### Running

To run do `Go` code, install the third party dependencies:

```shell
$ go mod download
```

To run it, do the following:

```shell
$ go run control.go
```

### Docker

To run the `Dockerfile`, build the image:

```shell
$ docker build . -t control
```

To run it set a environment variable `PORT` with the port you wanna use and, do the following:

```shell
$ export PORT=8000

$ docker run -p $PORT:$PORT -e PORT=$PORT control
```
