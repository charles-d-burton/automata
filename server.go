package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "automata/routes"
    "log"
)


func main() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/", routes.HandleIndex).Methods("GET")
    rtr.HandleFunc("/css/{rest:.*}", routes.HandleCSS).Methods("GET")
    rtr.HandleFunc("/font-awesome/{rest:.*}", routes.HandleFontAwesome).Methods("GET")
    rtr.HandleFunc("/js/{rest:.*}", routes.HandleJS).Methods("GET")
    rtr.HandleFunc("/fonts/{rest:.*}", routes.HandleFonts).Methods("GET")
    rtr.HandleFunc("/devices/smoker", routes.HandleSmoker).Methods("POST", "GET")
    rtr.HandleFunc("/devices/temp-sensor", routes.HandleTempSensor).Methods("POST", "GET")
    
    http.Handle("/", rtr)

    // Bind to a port and pass our router in
    log.Println("Listening on port 8000")
    log.Fatal(http.ListenAndServe(":8000", rtr))
}


