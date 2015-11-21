package devices

/*
Device ... top level struct.  Contains all the other possible JSON message types
*/

type DeviceReceive struct {
	ID string `json:"id"`
	Key string `json:"key"`
	Notify bool `json:"notify"`
	SmokerRead *SmokerRead
	TemperatureRead *TemperatureRead
}

type DeviceSend struct {
	ID string `json:"id"`
	Device interface{} `json:"device"`
}
