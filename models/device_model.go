package models

import (
	"encoding/json"
)

type Device struct {
	Id 		string        `json:"id"`
}

func NewDevice(id string) *Device {
	d := new(Device)
	d.Id = id

	return d
}

func (d *Device) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(*d, "", "		")
}

func (d *Device) UnmarshalJson(device_json []byte) error {
	return json.Unmarshal(device_json, &d)
}
