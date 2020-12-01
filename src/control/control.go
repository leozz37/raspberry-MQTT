package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// LedValue represents led on or off
type LedValue struct {
	State string `json:"state"`
}

func publish(client mqtt.Client, state string) {
	token := client.Publish("led", 0, false, state)
	token.Wait()
	time.Sleep(time.Second)
}

func sub(client mqtt.Client) {
	topic := "sensor"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

// SetupMQTT setup broker
func SetupMQTT() mqtt.Client {
	var broker = "localhost"
	var port = 1883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("control")
	opts.SetUsername("control")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

// SetLedValue set the led status
func SetLedValue(w http.ResponseWriter, r *http.Request) {
	log.Println("POST from /led")

	params := mux.Vars(r)
	var ledValue LedValue

	_ = json.NewDecoder(r.Body).Decode(&ledValue)
	ledValue.State = params["state"]

	fmt.Fprintln(w, ledValue.State)

	client := SetupMQTT()
	publish(client, ledValue.State)
	client.Disconnect(250)
}

// GetSensorValue returns the last sensor value
func GetSensorValue(w http.ResponseWriter, r *http.Request) {
	log.Println("GET from /sensor")
	fmt.Fprintln(w, rand.Intn(30))
}

// GetPort returns the port value from
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	return ":" + port
}

func main() {
	godotenv.Load()

	log.Println("Raspberry Control started!")

	router := mux.NewRouter()

	router.HandleFunc("/sensor", GetSensorValue)
	router.HandleFunc("/led/{state}", SetLedValue)

	port := GetPort()
	log.Println("Listening on port" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
