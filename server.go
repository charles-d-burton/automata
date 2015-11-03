package main

import (
    "net/http"
    "github.com/gorilla/mux"
    //"io/ioutil"
    "log"
    "encoding/json"
    "automata/devices"
)


func main() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/devices/smoker", handleSmoker).Methods("POST", "GET")

    // Bind to a port and pass our router in
    log.Println("Listening on port 8000")
    log.Fatal(http.ListenAndServe(":8000", rtr))
}

func handleSmoker(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    
    //Initialize the struct
    var smokerUpdate devices.SmokerRead
    err := decoder.Decode(&smokerUpdate)
    if err != nil {
       panic(err)    
    } 
    log.Println(smokerUpdate.ID)
    w.Write([]byte("OK\n"))
}
