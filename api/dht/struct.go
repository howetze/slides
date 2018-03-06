package dht

type Result []struct {
	Location struct {
		Country   string `json:"country"`
		Longitude string `json:"longitude"`
		Altitude  string `json:"altitude"`
		Latitude  string `json:"latitude"`
		ID        int    `json:"id"`
	} `json:"location"`
	Timestamp        string `json:"timestamp"`
	Sensordatavalues []struct {
		Value     string `json:"value"` // HL
		ValueType string `json:"value_type"`
		ID        int    `json:"id"`
	} `json:"sensordatavalues"`
	Sensor struct {
		SensorType struct {
			Manufacturer string `json:"manufacturer"`
			Name         string `json:"name"`
			ID           int    `json:"id"`
		} `json:"sensor_type"`
		Pin string `json:"pin"`
		ID  int    `json:"id"`
	} `json:"sensor"`
	SamplingRate interface{} `json:"sampling_rate"`
	ID           int         `json:"id"`
}
