package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "automata/handlers"
    "io/ioutil"
    "bytes"
    "log"
)


func main() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/", handleIndex).Methods("GET")
    rtr.HandleFunc("/css/{rest:.*}", handleCSS).Methods("GET")
    rtr.HandleFunc("/font-awesome/{rest:.*}", handleFontAwesome).Methods("GET")
    rtr.HandleFunc("/js/{rest:.*}", handleJS).Methods("GET")
    rtr.HandleFunc("/fonts/{rest:.*}", handleFonts).Methods("GET")
    rtr.HandleFunc("/devices/smoker", handleSmoker).Methods("POST", "GET")
    rtr.HandleFunc("/devices/temp-sensor", handleTempSensor).Methods("POST", "GET")
    http.Handle("/", rtr)

    // Bind to a port and pass our router in
    log.Println("Listening on port 8000")
    log.Fatal(http.ListenAndServe(":8000", rtr))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/")
}

func handleCSS(w http.ResponseWriter, r *http.Request) {
	path := mungeString("html/css/", mux.Vars(r)["rest"])
	http.ServeFile(w, r, path)
}

func handleFontAwesome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/font-awesome/")
}

func handleJS(w http.ResponseWriter, r *http.Request) {
	lead := "html/js/"
	tail := mux.Vars(r)["rest"]
	path := mungeString(lead, tail)
	http.ServeFile(w,r, path)

}

func handleFonts(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/fonts/")
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

func mungeString(lead string, tail string) string{
	var buffer bytes.Buffer
	buffer.WriteString(lead)
	buffer.WriteString(tail)
	log.Println(buffer.String())
	return buffer.String()
}
