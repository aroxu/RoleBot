package handler

import (
	"B1ackAnge1/RoleBot/model"
	"B1ackAnge1/RoleBot/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
	"time"
)

func VoteTimeOverHandler(session *discordgo.Session) {
	type Result struct {
		ID         string
		Requester  string
		ChannelID  string
		GuildID    string
		Agree      int
		Disagree   int
		VoteType   string
		Data       string
		StartDate  string
		EndDate    string
	}
	results := &[]Result{}
	currentTime := time.Now().UTC()
	utils.GetDB().Raw("SELECT votes.guild_id, votes.requester, votes.start_date, votes.end_date, votes.id, votes.channel_id, votes.agree, votes.disagree, votes.vote_type, votes.data FROM votes").Scan(&results)

	for _, result := range *results {
		if result.VoteType == "rank" {
			voteExpireDate, _ := time.Parse("02-Jan-2006 15:04:05", result.EndDate)
			if currentTime.Unix() > voteExpireDate.UTC().Unix() {
				targetMsgId := result.ID
				targetMsgChannelId := result.ChannelID

				requester, _ := session.User(result.Requester)

				roleCheckRequired := false

				var voteResultMsg string
				if result.Agree < result.Disagree {
					voteResultMsg = "반대표의 과반수 이상으로 해당 투표는 부결되었습니다.\n역할을 부여하지 않습니다.\n"
				} else if result.Agree > result.Disagree {
					voteResultMsg = "찬성표의 과반수 이상으로 해당 투표는 가결되었습니다.\n역할을 부여합니다.\n"
					roleCheckRequired = true
				} else if result.Agree == result.Disagree {
					voteResultMsg = "반대표와 찬성표가 같으므로 해당 투표는 부결되었습니다.\n역할을 부여하지 않습니다.\n"
				}

				embed := discordgo.MessageEmbed{
					Author: &discordgo.MessageEmbedAuthor{},
					Color:  39423,
					Title:  "🚫 역할 신청 투표 마감됨",
					Fields: []*discordgo.MessageEmbedField{},
					Footer: &discordgo.MessageEmbedFooter{
						Text: "개최일: " + result.StartDate + " | 마감일: " + result.EndDate,
					},
				}

				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   "역할 신청 투표가 마감되었습니다.",
					Value:  "신청자: " + requester.Mention() + "\n신청한 역할: **``" + strings.ReplaceAll(result.Data, "|", "") + "``**",
					Inline: true,
				})

				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:   "📊 투표 결과",
					Value:  "찬성: **``" + strconv.Itoa(result.Agree) + "표``**, 반대: **``" + strconv.Itoa(result.Disagree) +"표``**\n" + voteResultMsg,
					Inline: true,
				})

				server, errGuildNotFound := session.Guild(result.GuildID)
				if errGuildNotFound != nil {
					session.ChannelMessageSend(result.ChannelID, "❌ 서버를 찾을 수 없습니다.")
					log.Fatalln("서버를 찾을 수 없습니다: " + errGuildNotFound.Error())
					return
				}

				serverRoles := server.Roles

				foundRole := false
				for _, currentRole := range serverRoles {
					if currentRole.Name == strings.ReplaceAll(result.Data, "|", "") {
						foundRole = true

						if roleCheckRequired {
							errFailedAssignRoleToUser := session.GuildMemberRoleAdd(result.GuildID, result.Requester, currentRole.ID)
							if errFailedAssignRoleToUser != nil {
								session.ChannelMessageSend(targetMsgChannelId, "❌ 역할을 부여할 수 없습니다. 봇의 권한을 확인한 후 addrole 명령어로 다시 시도해주세요. addrole 명령어의 사용법은 도움말을 참조하세요.")
							}
						}
					}
				}

				if !foundRole && roleCheckRequired {
					session.ChannelMessageSend(targetMsgChannelId, "❌ 역할을 찾을 수 없습니다. 역할의 존재 유무와 봇의 권한을 확인한 후 addrole 명령어로 다시 시도해주세요. addrole 명령어의 사용법은 도움말을 참조하세요.")
				}

				if (currentTime.Unix() - voteExpireDate.UTC().Unix()) > 3 {
					session.ChannelMessageEdit(targetMsgChannelId, targetMsgId, "⚠️ 주의 : 해당 투표가 진행중이었을때 봇의 오프라인으로 전환 또는 Rate Limit을 감지했습니다. 투표결과가 정확하지 않을 수 있으니, 역할의 갯수를 직접 계산하길 권장드립니다.")
				}

				session.ChannelMessageEditEmbed(targetMsgChannelId, targetMsgId, &embed)
				utils.GetDB().Model(model.Vote{}).Delete(model.Vote{}, "id = ?", targetMsgId)
			}
		}
		if result.VoteType == "normal" {
			// Handler for normal vote
		}
	}
}