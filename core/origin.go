package core

import "net/url"

const (
	Region     = "reg"
	Zone       = "az"
	SubZone    = "sz"
	Host       = "host"
	InstanceId = "id"
)

// Origin - location
type Origin struct {
	Region     string `json:"region"`
	Zone       string `json:"zone"`
	SubZone    string `json:"sub-zone"`
	Host       string `json:"host"`
	InstanceId string `json:"instance-id"`
}

func NewValues(o Origin) url.Values {
	values := make(url.Values)
	if o.Region != "" {
		values.Add(Region, o.Region)
	}
	if o.Zone != "" {
		values.Add(Zone, o.Zone)
	}
	if o.SubZone != "" {
		values.Add(SubZone, o.SubZone)
	}
	if o.Host != "" {
		values.Add(Host, o.Host)
	}
	return values
}

func NewOrigin(values url.Values) Origin {
	o := Origin{}
	if values != nil {
		o.Region = values.Get(Region)
		o.Zone = values.Get(Zone)
		o.SubZone = values.Get(SubZone)
		o.Host = values.Get(Host)
	}
	return o
}
