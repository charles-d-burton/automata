package routes

import (
	"net/http"
    "github.com/gorilla/mux"
	"automata/handlers"
	"io/ioutil"
	"bytes"
	"log"
	
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/")
}

func HandleCSS(w http.ResponseWriter, r *http.Request) {
	path := mungeString("html/css/", mux.Vars(r)["rest"])
	http.ServeFile(w, r, path)
}

func HandleFontAwesome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/font-awesome/")
}

func HandleJS(w http.ResponseWriter, r *http.Request) {
	lead := "html/js/"
	tail := mux.Vars(r)["rest"]
	path := mungeString(lead, tail)
	http.ServeFile(w,r, path)

}

func HandleFonts(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/fonts/")
}

func HandleSmoker(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    go handlers.SmokerHandler(body)
    w.Write([]byte("OK\n"))
}

func HandleTempSensor(w http.ResponseWriter, r *http.Request) {
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