package handlers

import (
   "encoding/json"
   "automata/devices"
   "io"
)

func SmokerHandler(message io.Reader) bool {
   decoder := json.NewDecoder(message)

   //Initialize the struct
   var smokerUpdate devices.SmokerRead
   err := decoder.Decode(&smokerUpdate)
   if err != nil {
      //panic(err)
      return false
   }

   return true
}
