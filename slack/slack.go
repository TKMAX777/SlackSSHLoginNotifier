package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SlackHandler struct {
	webhook string
}

func New(webhook string) SlackHandler {
	return SlackHandler{webhook}
}

// Send message to slack
func (s SlackHandler) Send(web Webhook) error {
	if s.webhook == "" {
		return fmt.Errorf("Webhook not setuped")
	}
	bData, err := json.MarshalIndent(web, "", "    ")
	if err != nil {
		return err
	}

	_, err = http.Post(s.webhook, "application/json", bytes.NewBuffer(bData))
	return err
}
