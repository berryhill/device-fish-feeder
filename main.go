package main

import (
	"os"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var MqttClient *MQTT.Client

var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("TOPIC: %s\n", msg.Topic())
	//fmt.Printf("MSG: %s\n", msg.Payload())
	if string(msg.Payload()) == "feed_fish" {
		go HandleMessage(msg)
	} else {
		fmt.Println("Don't Understand Message")
	}
}

func main() {
  	fmt.Println("Starting Fish Feeder Device")

	e := echo.New()
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
	if token := MqttClient.Subscribe("test", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}

func HandleMessage (msg MQTT.Message) (string, error) {
	FeedFish()
	return "Feed", nil
}

func SendMessage(message []byte) error {
	token := MqttClient.Publish("web_to_bot", 0, false, message)
	token.Wait()
	fmt.Println("Sending Message")

	return nil
}

func FeedFish() {
	fmt.Println("IMPLEMENT FEEDING FISH HERE")
	fmt.Println("")
}
