package devices

/*
SmokerRead ... This struct is the next level device for a Smoker.  It possesses
a threshole and a current reading.
*/
type SmokerRead struct {
   Thresh int `json:"thresh"`
   Reading float64 `json:"reading"`
}

/*
SmokerWrite ... This struct sends a command to the Smoker.Device
*/
type SmokerWrite struct {
    Thresh int `json:"thresh"`
    RunTime string `json:"runtime"`   
}
