package core

import (
	"net/url"
	"strings"
)

const (
	RegionKey     = "reg"
	ZoneKey       = "az"
	SubZoneKey    = "sz"
	HostKey       = "host"
	InstanceIdKey = "id"
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
		values.Add(RegionKey, o.Region)
	}
	if o.Zone != "" {
		values.Add(ZoneKey, o.Zone)
	}
	if o.SubZone != "" {
		values.Add(SubZoneKey, o.SubZone)
	}
	if o.Host != "" {
		values.Add(HostKey, o.Host)
	}
	return values
}

func NewOrigin(values url.Values) Origin {
	o := Origin{}
	if values != nil {
		o.Region = values.Get(RegionKey)
		o.Zone = values.Get(ZoneKey)
		o.SubZone = values.Get(SubZoneKey)
		o.Host = values.Get(HostKey)
	}
	return o
}

func OriginMatch(target Origin, filter Origin) bool {
	if !StringMatch(target.Region, filter.Region) {
		return false
	}
	if !StringMatch(target.Zone, filter.Zone) {
		return false
	}
	if !StringMatch(target.SubZone, filter.SubZone) {
		return false
	}
	if !StringMatch(target.Host, filter.Host) {
		return false
	}
	return true
}

func StringMatch(target, filter string) bool {
	if filter == "" {
		return true
	}
	return strings.ToLower(target) == strings.ToLower(filter)
}
