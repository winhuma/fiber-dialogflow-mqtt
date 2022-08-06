package models

type BodyDialogFlow struct {
	Intent Payload `json:"originalDetectIntentRequest"`
}

type Payload struct {
	Data DialogData `json:"payload"`
}

type DialogData struct {
	Data Dialog `json:"data"`
}

type Dialog struct {
	ReplyToken string `json:"replyToken"`
	Message    MSG    `json:"Message"`
}

type MSG struct {
	Text string `json:"text"`
}
