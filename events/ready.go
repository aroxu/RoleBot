package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	var err = session.UpdateStatus(0, "golang")
	if err != nil {
		log.Println("Error updating status: ", err)
	}
	log.Println("Logged in as user " + session.State.User.ID)
}
