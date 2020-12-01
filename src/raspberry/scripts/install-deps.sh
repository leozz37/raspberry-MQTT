#!/bin/bash

git clone https://github.com/WiringPi/WiringPi.git /tmp/WiringPi
cd /tmp/WiringPi
./build
sudo ldconfig