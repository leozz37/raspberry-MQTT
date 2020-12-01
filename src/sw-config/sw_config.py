import serial

class Config:
    def __init__(self):
        self.get_inputs()

    def get_inputs(self):
        self.host = input("Hostname: ")
        self.led_topic = input("Topic Led: ")
        self.sensor_topic = input("Sensor Led: ")
        self.refresh_time = input("Refresh time: ")

    def send_values(self):
        payload = self.host + ";" + self.led_topic + ";" + \
                  self.sensor_topic + ";" + self.refresh_time
        print(payload)
        print("Press enter to setup...")
        
        ser = serial.Serial("/dev/tty0")
        payload_bytes = str.encode(payload)
        ser.write(payload_bytes)


if __name__ == "__main__":
    print("SW-Config started")
    config = Config()
    while True:
        config.send_values()
        input()
        config.get_inputs()
