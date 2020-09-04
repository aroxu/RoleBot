package events

import (
	"B1ackAnge1/RoleBot/handler"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func MessageReactionAdd(session *discordgo.Session, event *discordgo.MessageReactionAdd) {
	botID := session.State.User.ID
	if event.UserID == botID {
		return
	}
	message, err := session.ChannelMessage(event.ChannelID, event.MessageID)
	if err != nil {
		return
	}

	if strings.HasPrefix(message.Content, ">>> ✅") {
		go handler.ConfirmHandler(session, message, event)
	}

	isPoll := false
	for _, embed := range message.Embeds {
		if strings.HasPrefix(embed.Title, "🗳") {
			isPoll = true
		} else if strings.HasPrefix(embed.Title, "✅") {
			go handler.RankVoteReactionAddHandler(session, message, event)
		}
	}

	if !isPoll || message.Author.ID != botID || message.Author.Bot {
		return
	}

	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name != event.Emoji.Name {
			err := session.MessageReactionRemove(event.ChannelID, event.MessageID, reaction.Emoji.Name, event.UserID)
			if err != nil {
				return
			}
		}
	}
}