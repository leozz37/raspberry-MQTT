#!/bin/bash

# Installing wiringpi
git clone https://github.com/WiringPi/WiringPi.git
cd WiringPi
./build
ldconfig
cd ..
rm -rf WiringPi

# Installing paho.mqtt.cpp
git clone https://github.com/eclipse/paho.mqtt.c.git
cd paho.mqtt.c
git checkout v1.3.1
cmake -Bbuild -H. -DPAHO_WITH_SSL=ON -DPAHO_ENABLE_TESTING=OFF
cmake --build build/ --target install
ldconfig
cd ..
rm -rf paho.mqtt.c