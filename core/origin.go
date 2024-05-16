package core

// Origin - location
type Origin struct {
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	App        string `json:"app"`
	InstanceId string `json:"instance-id"`
}
