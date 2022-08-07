package services

import (
	"fiber-dialogflow-mqtt/libs/mymqtt"
	"fiber-dialogflow-mqtt/myenv"
	"fmt"
)

func PublishOpen(device string) error {
	var text = fmt.Sprintf("device:%s:open", device)
	mqttSession := mymqtt.GetInstantMQTT()
	token := mqttSession.Publish(myenv.MQTT_TOPIC, 2, false, text)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed publish to topic")
	}
	return nil
}

func PublishClose(device string) error {
	var text = fmt.Sprintf("device:%s:close", device)
	mqttSession := mymqtt.GetInstantMQTT()
	token := mqttSession.Publish(myenv.MQTT_TOPIC, 2, false, text)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed publish to topic")
	}
	return nil
}

func PublishDevice(device string) error {
	var text = fmt.Sprintf("device %s", device)
	mqttSession := mymqtt.GetInstantMQTT()
	token := mqttSession.Publish(myenv.MQTT_TOPIC, 2, false, text)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("failed publish to topic")
	}
	return nil
}

