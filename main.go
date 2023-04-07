package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	err := os.Setenv("SLACK_BOT_TOKEN", "xoxb-5089486479328-5055692711335-EMCgMVlbWyUNxNsv3iKsvzuP")
	if err != nil {
		return
	}
	err1 := os.Setenv("CHANNEL_ID", "C051Z01B9DY")
	if err1 != nil {
		return
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Assignment3_2023.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
