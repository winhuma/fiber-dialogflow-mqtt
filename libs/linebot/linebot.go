package linebot

import (
	"bytes"
	"encoding/json"
	"fiber-dialogflow-mqtt/myenv"
	"fmt"
	"net/http"
)

type BodyLineBot struct {
	ReplyToken string `json:"replyToken"`
	Message    []MSG  `json:"messages"`
}

type MSG struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func SendMsgLine(msg string, replyToken string) error {
	var err error
	var SendLine BodyLineBot

	SendLine.ReplyToken = replyToken
	SendLine.Message = append(SendLine.Message, MSG{
		Type: "text",
		Text: msg,
	})

	Client := http.Client{}
	uri := myenv.LINEBOT_ENDPOINT
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(SendLine)
	req, err := http.NewRequest(http.MethodPost, uri, payloadBuf)
	if err != nil {
		return err
	}
	req.Header = map[string][]string{}
	req.Header["Authorization"] = []string{"Bearer " + myenv.LINEBOT_ACCESS_TOKEN}
	req.Header["Content-Type"] = []string{"application/json"}
	resp, err := Client.Do(req)
	if err != nil {
		fmt.Println("###", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		return fmt.Errorf(
			`{"replyToken": %s, "message": %s, "status_code": %d}`,
			replyToken, buf.String(), resp.StatusCode,
		)
	}
	return nil
}
