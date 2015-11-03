package handlers

import (
   "encoding/json"
   "automata/devices"
   "log"
   "strings"
)

func TempSensorHandler(message []byte) bool {
   decoder := json.NewDecoder(strings.NewReader(string(message)))

   //Initialize the struct
   var smokerUpdate devices.SmokerRead
   err := decoder.Decode(&smokerUpdate)
   if err != nil {
      //panic(err)
      //return false
   }
   log.Println(smokerUpdate.Reading) 
   return true
}
