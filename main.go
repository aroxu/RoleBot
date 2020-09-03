package main

import (
	"B1ackAnge1/RoleBot/events"
	"B1ackAnge1/RoleBot/initializer"
	"B1ackAnge1/RoleBot/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Hello World!")
	rawConfig, errFindConfigFile := ioutil.ReadFile("config.toml") // just pass the file name
	if errFindConfigFile != nil {
		fmt.Println("Error while load config file: " + errFindConfigFile.Error())
		return
	}
	errLoadConfigData, token := utils.GetToken(string(rawConfig))
	if  errLoadConfigData != nil {
		fmt.Println("Error while load config data: " + errLoadConfigData.Error())
	}

	var bot, err = discordgo.New("Bot " + token)
	// register events
	bot.AddHandler(events.Ready)
	bot.AddHandler(events.MessageCreate)

	initializer.InitCommands()

	err = bot.Open()

	if err != nil {
		log.Fatalln("Error opening Discord session: ", err)
	}

	fmt.Println("Bot is now running.")

	// wait forever
	select {}
}
