package main

import (
	"B1ackAnge1/RoleBot/events"
	"B1ackAnge1/RoleBot/handler"
	"B1ackAnge1/RoleBot/initializer"
	"B1ackAnge1/RoleBot/utils"
	"io/ioutil"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.Println("Initializing Database...")
	errInitDB := initializer.InitDB()
	if errInitDB != nil {
		log.Fatalln("Error while load config file: " + errInitDB.Error())
		return
	}
	rawConfig, errFindConfigFile := ioutil.ReadFile("config.toml") // just pass the file name
	if errFindConfigFile != nil {
		log.Fatalln("Error while load config file: " + errFindConfigFile.Error())
		return
	}
	errLoadConfigData, token := utils.GetToken(string(rawConfig))
	if errLoadConfigData != nil {
		log.Fatalln("Error while load config data: " + errLoadConfigData.Error())
	}

	var bot, err = discordgo.New("Bot " + token)
	// register events
	bot.AddHandler(events.Ready)
	bot.AddHandler(events.MessageCreate)
	bot.AddHandler(events.MessageReactionAdd)
	bot.AddHandler(events.MessageReactionRemove)

	initializer.InitCommands()

	err = bot.Open()
	go checkTimeForVote(bot)

	if err != nil {
		log.Fatalln("Error opening Discord session: ", err)
	}

	log.Println("Bot is now running.")

	// wait forever
	select {}
}

func checkTimeForVote(session *discordgo.Session) {
	for {
		time.Sleep(time.Second)
		go handler.VoteTimeOverHandler(session)
	}
}
