package models

import(
	"encoding/json"
)

type FeedReport struct {
	Id			string        `json:"id"`
	Timestamp	string        `json:"timestamp"`
	Success 	bool          `json:"success"`
}

func NewFeedReport(id string, success bool) *FeedReport {
	fr := new(FeedReport)
	fr.Id = id
	fr.Success = success

	return fr
}

func (fr *FeedReport) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(*fr, "", "		")
}

func (fr *FeedReport) UnmarshalJson(report_json []byte) error {
	return json.Unmarshal(report_json, &fr)
}

