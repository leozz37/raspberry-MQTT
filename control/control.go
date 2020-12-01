package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// LedValue represents led on or off
type LedValue struct {
	State string `json:"state"`
}

// SetLedValue set the led status
func SetLedValue(w http.ResponseWriter, r *http.Request) {
	log.Println("POST from /led")

	params := mux.Vars(r)
	var ledValue LedValue

	_ = json.NewDecoder(r.Body).Decode(&ledValue)
	ledValue.State = params["state"]

	fmt.Fprintln(w, ledValue.State)
	// TODO: Implement turning on LED
}

// GetSensorValue returns the last sensor value
func GetSensorValue(w http.ResponseWriter, r *http.Request) {
	// TODO: Get sensor value
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
