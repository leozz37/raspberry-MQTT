#include <iostream>
#include <wiringPi.h>

int main() {
    wiringPiSetup();
    pinMode(0, OUTPUT);
    pinMode(1, INPUT);

    while(1) {
        if(digitalRead(1) == 1) {
            digitalWrite(0, !digitalRead(0));
            delay(500);
        }
    }
    return 0;
}