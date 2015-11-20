package devices

/*
Device ... top level struct.  Contains all the other possible JSON message types
*/

type Device struct {
	ID string `json:"id"`
	Key string `json:"key"`
	Notify bool `json:"notify"`
	SmokerRead *SmokerRead
	SmokerWrite *SmokerWrite
	TemperatureReade *TemperatureRead
	
}
