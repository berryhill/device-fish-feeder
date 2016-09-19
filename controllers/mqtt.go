package controllers

import (
	"fmt"
	"os"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var MqttClient *MQTT.Client
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	if string(msg.Payload()) == "feed_fish" {
		go HandleFeedFish(msg)
	} else {
		fmt.Println("ERROR: Don't Understand Message")
	}
}

func StartMqttClient() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://138.68.47.232:8883")
	opts.SetClientID("devic")
	//opts.SetClientID(Device)
	opts.SetDefaultPublishHandler(f)
	opts.SetUsername("mosquitto")
	opts.SetPassword("amh05055")
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