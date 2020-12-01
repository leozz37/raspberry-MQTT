# import serial
      
class Config:
    def __init__(self):
        self.get_inputs()

    def get_inputs(self):
        self.host = input("Hostname: ")
        self.led_topic = input("Topic Led: ")
        self.sensor_topic = input("Sensor Led: ")
        self.refresh_time = input("Refresh time: ")

    def send_values(self):
        payload = self.host + ";" + self.led_topic + ";" + self.sensor_topic + ";" + self.refresh_time
        print(payload)
        print("Press enter to setup...")
        # ser = serial.Serial("/dev/ttyS0", 9600)
        # ser.write(payload)

if __name__ == "__main__":
    config = Config()
    while True:
        config.send_values()
        input()
        config.get_inputs()