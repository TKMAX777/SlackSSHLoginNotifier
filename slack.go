package main

import (
	"fmt"
	"strings"
	"time"

	"local.packages/slack"
)

func SendLoginMessage() error {
	var message = slack.Webhook{
		IconEmoji: ":key:",
		UserName:  fmt.Sprintf("[%s] SSH Login Alert", strings.ToUpper(HostName)),
		Text:      fmt.Sprintf("%s has logined at %s\n from: %s", UserName, time.Now().Format(time.UnixDate), ClientIP),
	}

	return Slack.Send(message)
}
