package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	var err = session.UpdateStatus(0, "golang")
	if err != nil {
		fmt.Println("Error updating status: ", err)
	}
	fmt.Println("Logged in as user " + session.State.User.ID)
}
