package controllers

import (
	"fmt"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func HandleGetDevice() {
	device_json, err := Device.MarshalJson()
	if err != nil {
		fmt.Println(err)
	}

	jsonStr := string(device_json)
	message := Device.Id + " " + "device" + " " + "Pass" + " " + jsonStr
	SendMessage([]byte(message))
}

func HandleFeedFish (msg MQTT.Message) {
	fmt.Println("Task_Recieved: feed_fish")
	SendMessage([]byte("Task_Recieved: feed_fish"))

	report, err := FeedFish()
	report_json, errr := report.MarshalJson()
	if errr != nil {
		fmt.Println(errr)
	}
	if err == nil {
		jsonStr := string(report_json)
		message := Device.Id + " " + "report" + " " + "Pass" + " " + jsonStr
		SendMessage([]byte(message))
	} else {
		jsonStr := string(report_json)
		message := Device.Id + " " + "report" + " " + "Failed" + " " + jsonStr
		SendMessage([]byte(message))
	}
}
