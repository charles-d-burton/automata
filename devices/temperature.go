package devices

/*
TemperatureRead ...  Handles marshalling of temperature reading
*/

type TemperatureRead struct {
   Reading float32 `json:"reading"`
} `json:"temperature_read"`
