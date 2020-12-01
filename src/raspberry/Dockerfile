FROM gcc:9

RUN apt-get update \
 && apt-get -y install \
    sudo \
    cmake \
    git

WORKDIR /app
COPY . .

WORKDIR /app/scripts
RUN chmod +x ./install-deps.sh

RUN sudo ./install-deps.sh

WORKDIR /app
RUN gcc main.c -o mqtt_exemplo -lpaho-mqtt3c -I/usr/local/include -L/usr/local/lib -lwiringPi -pthread -lpthread -Wall

CMD [ "./rasp" ]
