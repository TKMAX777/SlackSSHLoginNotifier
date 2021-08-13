package main

import (
	"os"

	"local.packages/slack"
)

var Slack slack.SlackHandler

var HostName string
var UserName string
var ClientIP string

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultValue
}

func init() {
	var hook = os.Getenv("SSHSlackHookURL")
	if hook == "" {
		panic("Hook URL not set!")
	}
	Slack = slack.New(hook)

	UserName = os.Getenv("USER")
	ClientIP = os.Getenv("SSH_CLIENT")
	HostName = os.Getenv("HOSTNAME")
}

func main() {
	SendLoginMessage()
}
