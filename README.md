# Raspberry MQTT Project

![cover](resources/images/infra.png)

Raspberry MQTT example

## Dependencies

You need both [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/install/) installed on your Raspberry. If you are not using Docker, you need [Go](https://golang.org/doc/install) and [Mosquitto MQTT Broker](https://mosquitto.org/).

## Raspberry

Install these dependencies:

```shell
$ sudo apt install -y wiringpi cmake git
```

Build the binary:

```shell
$ gcc main.c -o rasp -lpaho-mqtt3c -I/usr/local/include -L/usr/local/lib -lwiringPi -Wall
```

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

### Docker compose

To run all the project services, do the following:

```shell
$ docker-compose up
```