package handler

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
	"time"
)

func RankVoteReactionAddHandler(session *discordgo.Session, message *discordgo.Message, event *discordgo.MessageReactionAdd) {
	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name != event.Emoji.Name {

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

			if event.Emoji.Name == "⭕" {
				for _, origin := range message.Embeds {
					for _, field := range origin.Fields {
						log.Println(field.Value)
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**
							currentResultToString := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(field.Value, "찬성: ", ""), "반대: ", ""), "*", ""), "`", ""), "표", ""), " ", "")

							currentResult := strings.Split(currentResultToString, ",")
							agree, _ := strconv.Atoi(currentResult[0])
							disagree, _ := strconv.Atoi(currentResult[1])
							agree++

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
				for _, origin := range message.Embeds {
					for _, field := range origin.Fields {
						log.Println(field.Value)
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**
							currentResultToString := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(field.Value, "찬성: ", ""), "반대: ", ""), "*", ""), "`", ""), "표", ""), " ", "")

							currentResult := strings.Split(currentResultToString, ",")
							agree, _ := strconv.Atoi(currentResult[0])
							disagree, _ := strconv.Atoi(currentResult[1])
							disagree++

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
	for _, reaction := range message.Reactions {
		if reaction.Emoji.Name != event.Emoji.Name {

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

			if event.Emoji.Name == "⭕" {
				for _, origin := range message.Embeds {
					for _, field := range origin.Fields {
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**
							currentResultToString := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(field.Value, "찬성: ", ""), "반대: ", ""), "*", ""), "`", ""), "표", ""), " ", "")

							currentResult := strings.Split(currentResultToString, ",")
							agree, _ := strconv.Atoi(currentResult[0])
							disagree, _ := strconv.Atoi(currentResult[1])
							agree--

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
				for _, origin := range message.Embeds {
					for _, field := range origin.Fields {
						if field.Name == "📊 투표 현황" {
							//찬성: **``0표``**, 반대: **``0표``**
							currentResultToString := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(field.Value, "찬성: ", ""), "반대: ", ""), "*", ""), "`", ""), "표", ""), " ", "")

							currentResult := strings.Split(currentResultToString, ",")
							agree, _ := strconv.Atoi(currentResult[0])
							disagree, _ := strconv.Atoi(currentResult[1])
							disagree--

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
