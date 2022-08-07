package myenv

import (
	"fmt"
	"os"
)

var SERVICE_PORT string

var LINEBOT_ENDPOINT string
var LINEBOT_SECRET string
var LINEBOT_ACCESS_TOKEN string
var DIALOGFOW_WEBHOOK string

var MQTT_BROKER string
var MQTT_PORT string
var MQTT_PROTOCAL string
var MQTT_USERNAME string
var MQTT_PASSWORD string
var MQTT_TOPIC string

func GetMyENV(EnvKey string, defaultValue ...string) string {
	value := os.Getenv(EnvKey)
	fmt.Printf("[ENV] %s: %s\n", EnvKey, value)
	if len(value) == 0 && len(defaultValue) != 0 {
		value = defaultValue[0]
	} else if len(value) == 0 && len(defaultValue) == 0 {
		value = ""
	}
	return value
}

func SetEnv() {
	SERVICE_PORT = GetMyENV("PORT")

	LINEBOT_ENDPOINT = GetMyENV("LINEBOT_ENDPOINT")
	LINEBOT_SECRET = GetMyENV("LINEBOT_SECRET")
	LINEBOT_ACCESS_TOKEN = GetMyENV("LINEBOT_ACCESS_TOKEN")
	DIALOGFOW_WEBHOOK = GetMyENV("DIALOGFOW_WEBHOOK")

	MQTT_BROKER = GetMyENV("MQTT_BROKER")
	MQTT_PORT = GetMyENV("MQTT_PORT")
	MQTT_PROTOCAL = GetMyENV("MQTT_PROTOCAL")
	MQTT_USERNAME = GetMyENV("MQTT_USERNAME")
	MQTT_PASSWORD = GetMyENV("MQTT_PASSWORD")
	MQTT_TOPIC = GetMyENV("MQTT_TOPIC")
}
