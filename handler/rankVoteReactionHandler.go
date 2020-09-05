package handler

import (
	"B1ackAnge1/RoleBot/model"
	"B1ackAnge1/RoleBot/utils"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

func RankVoteReactionAddHandler(session *discordgo.Session, message *discordgo.Message, event *discordgo.MessageReactionAdd) {
	type AgreeAndDisagreeData struct {
		Agree      int
		Disagree   int
	}

	var agree int
	var disagree int

	var agreeAndDisagreeData AgreeAndDisagreeData

	utils.GetDB().Raw("SELECT agree, disagree FROM votes WHERE id = ?", message.ID).Scan(&agreeAndDisagreeData)

	agree = agreeAndDisagreeData.Agree
	disagree = agreeAndDisagreeData.Disagree

	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name != event.Emoji.Name {
			var embed discordgo.MessageEmbed

			if event.Emoji.Name == "⭕" {
				agree++
				utils.GetDB().Model(&model.Vote{}).Where("id = ?", message.ID).Update("agree", agree)
				for _, origin := range message.Embeds {
					embed = discordgo.MessageEmbed{
						Author: &discordgo.MessageEmbedAuthor{},
						Color:  39423,
						Title:  "✅ 역할 신청 투표 개최됨",
						Fields: []*discordgo.MessageEmbedField{},
						Footer: &discordgo.MessageEmbedFooter{
							Text: origin.Footer.Text,
						},
					}
					for _, field := range origin.Fields {
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**

							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  "찬성: **``" + strconv.Itoa(agree) + "표``**, 반대: **``" + strconv.Itoa(disagree) + "표``**",
								Inline: field.Inline,
							})
						}
						if field.Name == "해당 사용자에게 해당 태그를 받을 최소한의 자격이 된다고 판단된다면 ⭕, 아니라면 ❌ 이모티콘을 추가해주세요." {
							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  field.Value,
								Inline: field.Inline,
							})
						}
					}
				}

				confirmMessage, err := session.ChannelMessageEditEmbed(message.ChannelID, message.ID, &embed)
				if err != nil {
					session.ChannelMessageSend(confirmMessage.ChannelID, err.Error())
					return
				}
				return
			} else if event.Emoji.Name == "❌" {
				disagree++
				utils.GetDB().Model(&model.Vote{}).Where("id = ?", message.ID).Update("disagree", disagree)
				for _, origin := range message.Embeds {
					embed = discordgo.MessageEmbed{
						Author: &discordgo.MessageEmbedAuthor{},
						Color:  39423,
						Title:  "✅ 역할 신청 투표 개최됨",
						Fields: []*discordgo.MessageEmbedField{},
						Footer: &discordgo.MessageEmbedFooter{
							Text: origin.Footer.Text,
						},
					}
					for _, field := range origin.Fields {
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**

							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  "찬성: **``" + strconv.Itoa(agree) + "표``**, 반대: **``" + strconv.Itoa(disagree) + "표``**",
								Inline: field.Inline,
							})
						}
						if field.Name == "해당 사용자에게 해당 태그를 받을 최소한의 자격이 된다고 판단된다면 ⭕, 아니라면 ❌ 이모티콘을 추가해주세요." {
							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  field.Value,
								Inline: field.Inline,
							})
						}
					}
				}

				confirmMessage, err := session.ChannelMessageEditEmbed(message.ChannelID, message.ID, &embed)
				if err != nil {
					session.ChannelMessageSend(confirmMessage.ChannelID, err.Error())
					return
				}
				return
			}
		}
	}
}

func RankVoteReactionRemoveHandler(session *discordgo.Session, message *discordgo.Message, event *discordgo.MessageReactionRemove) {
	type AgreeAndDisagreeData struct {
		Agree      int
		Disagree   int
	}

	var agree int
	var disagree int

	var agreeAndDisagreeData AgreeAndDisagreeData

	utils.GetDB().Raw("SELECT agree, disagree FROM votes WHERE id = ?", message.ID).Scan(&agreeAndDisagreeData)

	agree = agreeAndDisagreeData.Agree
	disagree = agreeAndDisagreeData.Disagree

	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name != event.Emoji.Name {
			var embed discordgo.MessageEmbed

			if event.Emoji.Name == "⭕" {
				agree--
				utils.GetDB().Model(&model.Vote{}).Where("id = ?", message.ID).Update("agree", agree)
				for _, origin := range message.Embeds {
					embed = discordgo.MessageEmbed{
						Author: &discordgo.MessageEmbedAuthor{},
						Color:  39423,
						Title:  "✅ 역할 신청 투표 개최됨",
						Fields: []*discordgo.MessageEmbedField{},
						Footer: &discordgo.MessageEmbedFooter{
							Text: origin.Footer.Text,
						},
					}
					for _, field := range origin.Fields {
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**

							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  "찬성: **``" + strconv.Itoa(agree) + "표``**, 반대: **``" + strconv.Itoa(disagree) + "표``**",
								Inline: field.Inline,
							})
						}
						if field.Name == "해당 사용자에게 해당 태그를 받을 최소한의 자격이 된다고 판단된다면 ⭕, 아니라면 ❌ 이모티콘을 추가해주세요." {
							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  field.Value,
								Inline: field.Inline,
							})
						}
					}
				}

				confirmMessage, err := session.ChannelMessageEditEmbed(message.ChannelID, message.ID, &embed)
				if err != nil {
					session.ChannelMessageSend(confirmMessage.ChannelID, err.Error())
					return
				}
				return
			} else if event.Emoji.Name == "❌" {
				disagree--
				utils.GetDB().Model(&model.Vote{}).Where("id = ?", message.ID).Update("disagree", disagree)
				for _, origin := range message.Embeds {
					embed = discordgo.MessageEmbed{
						Author: &discordgo.MessageEmbedAuthor{},
						Color:  39423,
						Title:  "✅ 역할 신청 투표 개최됨",
						Fields: []*discordgo.MessageEmbedField{},
						Footer: &discordgo.MessageEmbedFooter{
							Text: origin.Footer.Text,
						},
					}
					for _, field := range origin.Fields {
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**

							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  "찬성: **``" + strconv.Itoa(agree) + "표``**, 반대: **``" + strconv.Itoa(disagree) + "표``**",
								Inline: field.Inline,
							})
						}
						if field.Name == "해당 사용자에게 해당 태그를 받을 최소한의 자격이 된다고 판단된다면 ⭕, 아니라면 ❌ 이모티콘을 추가해주세요." {
							embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
								Name:   field.Name,
								Value:  field.Value,
								Inline: field.Inline,
							})
						}
					}
				}

				confirmMessage, err := session.ChannelMessageEditEmbed(message.ChannelID, message.ID, &embed)
				if err != nil {
					session.ChannelMessageSend(confirmMessage.ChannelID, err.Error())
					return
				}
				return
			}
		}
	}
}
