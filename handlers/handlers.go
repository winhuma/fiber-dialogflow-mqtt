package handlers

import (
	"encoding/json"
	"fiber-dialogflow-mqtt/helpers"
	"fiber-dialogflow-mqtt/libs/linebot"
	"fiber-dialogflow-mqtt/models"
	"fiber-dialogflow-mqtt/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.Status(200).JSON(models.ResponseSuccess("hello!"))
}

func LineBotCallback(c *fiber.Ctx) error {
	var err error
	var dialogflowBody models.BodyDialogFlow
	var myMSG string
	json.Unmarshal(c.Body(), &dialogflowBody)
	getToken := dialogflowBody.Intent.Data.Data.ReplyToken
	getTextPayload := dialogflowBody.Intent.Data.Data.Message.Text
	match := helpers.CheckFormatFuncCallback(getTextPayload)
	if !match {
		myMSG = "funcking matching bitch!"
		err = linebot.SendMsgLine(myMSG, getToken)
		if err != nil {
			return err
		}
		return c.Status(400).JSON(models.ResponseFail("Text not match"))
	}
	text := strings.Split(strings.ToLower(getTextPayload), " ")
	var action string
	var device string
	if len(test) > 1 {
		action = text[0]
		device = text[1]
	} else {
		action = text[0]
	}
	switch action {
	case "1":
		err = services.PublishDevice(action)
		myMSG = "ฮา"
	case "open":
		err = services.PublishOpen(device)
		myMSG = "open device success"
	case "close":
		err = services.PublishClose(device)
		myMSG = "close device success"
	case "test":
		myMSG = "Fuck duck"
	default:
		myMSG = "not thing"
	}
	if err != nil {
		err = linebot.SendMsgLine(err.Error(), getToken)
		return err
	}
	err = linebot.SendMsgLine(myMSG, getToken)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(models.ResponseSuccess("line bot callback"))
}
