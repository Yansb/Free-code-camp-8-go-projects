package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	api := slack.New(goDotEnvVariable("SLACK_BOT_TOKEN"))
	channelArr := []string{goDotEnvVariable("SLACK_CHANNEL_ID")}
	//name of the file \/
	fileArr := []string{"comprovante.pdf"}

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
		fmt.Printf("Name: %s, Url: %s\n", file.Name, file.URLPrivate)
	}
}
