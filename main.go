package main

import (
	"fiber-dialogflow-mqtt/libs/mymqtt"
	"fiber-dialogflow-mqtt/myenv"
	"fiber-dialogflow-mqtt/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("[WARNING] .env file not found")
	}
	myenv.SetEnv()
	mymqtt.ConnectMQTT()

	app := routes.NewFiber()
	routes.GetRoutes(app)

	app.Listen(":3000")
}
