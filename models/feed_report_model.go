package models

import(
	"time"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type FeedReport struct {
	Id				bson.ObjectId        `json:"id"`
	DeviceId 		string 		         `json:"device_id"`
	Timestamp		time.Time        	 `json:"timestamp"`
	Success 		bool          		 `json:"success"`
}

func NewFeedReport(success bool, deviceId string) *FeedReport {
	fr := new(FeedReport)
	fr.DeviceId = deviceId
	fr.Timestamp = time.Now()
	fr.Id = bson.NewObjectId()
	fr.Success = success

	return fr
}

func (fr *FeedReport) MarshalJson() ([]byte, error) {
	return json.Marshal(*fr)
}

func (fr *FeedReport) UnmarshalJson(report_json []byte) error {
	return json.Unmarshal(report_json, &fr)
}

