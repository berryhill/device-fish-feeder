package main

import (
	"os"
	"fmt"
	"time"

	"github.com/device-fish-feeder/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var Device *models.Device

var MqttClient *MQTT.Client
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MSG: %s\n", msg.Payload())
	if string(msg.Payload()) == "feed_fish" {
		go HandleFeedFish(msg)
	} else {
		fmt.Println("ERROR: Don't Understand Message")
	}
}

func main() {
  	fmt.Println("Starting Fish Feeder Device")
	e := echo.New()

	Device = models.NewDevice("1234")

	StartMqttClient()
	fmt.Println("Running a Server on localhost:1323")
	e.Run(standard.New(":1323"))
}

func StartMqttClient() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://test.mosquitto.org:1883")
	opts.SetClientID("device")
	opts.SetDefaultPublishHandler(f)
	MqttClient = MQTT.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := MqttClient.Subscribe("to_device", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func SendMessage(message []byte) error {
	token := MqttClient.Publish("to_web", 0, false, message)
	token.Wait()
	return nil
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
		message := []byte(Device.Id + " " + "report" + " " + "Pass" + " ") + report_json
		SendMessage(message)
	} else {
		message := []byte(Device.Id + " " + "report" + " " + "Failed" + " ") + report_json
		SendMessage(message)
	}
}

func HandleGetDevice() {
	device_json, err := Device.MarshalJson()
	if err != nil {
		fmt.Println(err)
	}
	message := []byte(Device.Id + " " + "device" + " " + "Pass" + " ") + device_json
	SendMessage(message)
}

func FeedFish() (*models.FeedReport, error) {
	time.Sleep(4 * time.Second)
	success := true // place holders
	err := false	// ...
	feed_report := models.NewFeedReport("1324", success)

	return feed_report, err
}

