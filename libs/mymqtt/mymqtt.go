package mymqtt

import (
	"fiber-dialogflow-mqtt/myenv"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var session mqtt.Client

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("[INFO] MQTT Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("[INFO] MQTT Connection lost: %v", err)
}

func ConnectMQTT() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", myenv.MQTT_PROTOCAL, myenv.MQTT_BROKER, myenv.MQTT_PORT))
	opts.SetClientID("golang-script")
	if myenv.MQTT_USERNAME != "" {
		opts.SetUsername(myenv.MQTT_USERNAME)
	}
	if myenv.MQTT_USERNAME != "" {
		opts.SetPassword(myenv.MQTT_PASSWORD)
	}

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	session = mqtt.NewClient(opts)
	if token := session.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func GetInstantMQTT() mqtt.Client {
	return session
}
