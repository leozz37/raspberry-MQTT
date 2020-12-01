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
RUN gcc main.c -o rasp -lpaho-mqtt3c -I/usr/local/include -L/usr/local/lib -lwiringPi -Wall

CMD [ "./rasp" ]
