package handler

import (
	"B1ackAnge1/RoleBot/model"
	"B1ackAnge1/RoleBot/utils"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
	"time"
)

func ConfirmHandler(session *discordgo.Session, message *discordgo.Message, event *discordgo.MessageReactionAdd) {

	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name != event.Emoji.Name {
			if event.Emoji.Name == "❌" {
				session.MessageReactionsRemoveAll(message.ChannelID, message.ID)
				cancelMessage, _ := session.ChannelMessageEdit(message.ChannelID, message.ID, "❌ 사용자에 의해 취소되었습니다.")
				time.Sleep(time.Second * 10)
				session.ChannelMessageDelete(cancelMessage.ChannelID, cancelMessage.ID)
				return
			}
		}
	}

	originalMessage := message.Content
	userMsgPattern := regexp.MustCompile(`([0-9])\w+`)

	targetMessage := userMsgPattern.FindString(originalMessage)
	rolesMsgString := strings.Split(originalMessage, "឵")[1]

	tempStr := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(rolesMsgString, "|", ""), "*", ""), "`", "")
	roles := strings.Split(tempStr, ", ")

	for _, currentRole := range roles {

		requester, errFailedFindUser := session.User(targetMessage)
		if errFailedFindUser != nil {
			session.ChannelMessageSend(message.ChannelID, errFailedFindUser.Error())
			return
		}

		session.ChannelMessageDelete(message.ChannelID, message.ID)

		embed := discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  39423,
			Title:  "✅ 역할 신청 투표 개최됨",
			Fields: []*discordgo.MessageEmbedField{},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "개최일 ",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		}

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "해당 사용자에게 해당 태그를 받을 최소한의 자격이 된다고 판단된다면 ⭕, 아니라면 ❌ 이모티콘을 추가해주세요.",
			Value:  "신청자: " + requester.Mention() + "\n신청한 역할: **``" + strings.ReplaceAll(currentRole, "|", "") + "``**",
			Inline: true,
		})

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "📊 투표 현황",
			Value:  "찬성: **``0표``**, 반대: **``0표``**",
			Inline: true,
		})

		confirmMessage, err := session.ChannelMessageSendEmbed(message.ChannelID, &embed)
		if err != nil {
			session.ChannelMessageSend(confirmMessage.ChannelID, err.Error())
			return
		}

		rankVoteData := model.Vote{
			ID: message.ID,
			StartDate: time.Now().Format(time.RFC3339),
			EndDate: time.Now().Add(24*time.Hour).Format(time.RFC3339),
			Agree: 0,
			Disagree: 0,
			VoteType: "rank",
			Data: strings.ReplaceAll(currentRole, "|", ""),
		}

		utils.GetDB().Create(&rankVoteData)


		session.MessageReactionAdd(confirmMessage.ChannelID, confirmMessage.ID, "⭕")
		session.MessageReactionAdd(confirmMessage.ChannelID, confirmMessage.ID, "❌")
	}
}
