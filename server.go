package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "automata/handlers"
    "io/ioutil"
    "log"
)


func main() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/devices/smoker", handleSmoker).Methods("POST", "GET")
    rtr.HandleFunc("/devices/temp-sensor", handleTempSensor).Methods("POST", "GET")

    // Bind to a port and pass our router in
    log.Println("Listening on port 8000")
    log.Fatal(http.ListenAndServe(":8000", rtr))
}

func handleSmoker(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    go handlers.SmokerHandler(body)
    w.Write([]byte("OK\n"))
}

func handleTempSensor(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    go handlers.TempSensorHandler(body)
    w.Write([]byte("OK\n"))
}
