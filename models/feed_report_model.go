package models

import(
	"encoding/json"
)

type FeedReport struct {
	Id			string        `json:"id"`
	Timestamp	string        `json:"timestamp"`
	Success 	bool          `json:"success"`
}

func NewFeedReport(id string, success bool) *Device {
	d := new(FeedReport)
	d.Id = id
	d.Success = success

	return d
}

func (d *FeedReport) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(*d, "", "		")
}

func (d *FeedReport) UnmarshalJson(report_json []byte) error {
	return json.Unmarshal(report_json, &d)
}

