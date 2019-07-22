package main

import (
	"encoding/json"
	"fmt"
	"go-stateservice/com/nike/model"
	"net/http"
	"strconv"
)

func main() {
	/*	r := mux.NewRouter()

		r.HandleFunc("/", healthHandler).Methods("GET")
		r.HandleFunc("/states", statesHandler).Methods("GET")
		r.HandleFunc("/health" , HealthCheck).Methods("GET")

		fmt.Println("Starting the server at 7070")
		_ = http.ListenAndServe(":7070", r)*/
	model.DisplayState()

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("home handler")
}

func statesHandler(w http.ResponseWriter, r *http.Request) {
	data := []byte("River Handler called")
	writeJsonResponse(w, http.StatusOK, data)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Since we're here, we already know that HTTP service is up. Let's just check the state of the boltdb connection
	data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
	writeJsonResponse(w, http.StatusOK, data)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
