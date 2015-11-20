package routes
import(
    "net/http"
    "github.com/gorilla/websocket"
    "log"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

/*
HandleWebSocket ... This function handles an initiated Websocket.  It's not currently used
but will eventually provide the mechanism for real time data
stats from connected devices.
*/
func HandleWebSocket (w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    for {
        message := conn.ReadJSON
        log.Println(message)        
    }
}