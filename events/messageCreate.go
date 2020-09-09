package events

import (
	"B1ackAnge1/RoleBot/handler"
	"B1ackAnge1/RoleBot/utils"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	rawConfig, errFindConfigFile := ioutil.ReadFile("config.toml") // just pass the file name
	if errFindConfigFile != nil {
		fmt.Println("Error while load config file: " + errFindConfigFile.Error())
		return
	}
	errLoadConfigData, prefix := utils.GetPrefix(string(rawConfig))
	if errLoadConfigData != nil {
		fmt.Println("Error while load config data: " + errLoadConfigData.Error())
	}
	go handler.HandleCreatedMessage(session, event, prefix)
}
