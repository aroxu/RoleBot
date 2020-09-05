package rankvote

import (
	"B1ackAnge1/RoleBot/handler"
	"strings"
	"time"
)

func Initialize() {
	handler.AddCommand(
		handler.Command{
			Run:                  run,
			Names:                []string{"rankvote"},
			RequiredArgumentType: []string{roles},
			Usage:                map[string]string{"필요한 권한": "**``없음``**", "설명": "``명령어를 신청한 사람에게 역할 추가 투표를 개최합니다. 만료 기한은 1일 입니다.``", "사용법": "```css\n?!rankvote rolename1 rolename2 rolename3 ...```"},
		},
	)
}

const (
	roles = "role1, role2, role3, ..."
)

func run(ctx handler.CommandContext) error {
	if len(ctx.Arguments[roles]) == 0 {
		ctx.Message.Reply("❌ 요청할 역할을 기재하시고 다시 시도해주세요.")
		return nil
	}

	requester, _ := ctx.Message.AuthorMember()
	guild := ctx.Message.GuildID
	serverRoles, _ := ctx.Session.GuildRoles(guild)
	role := strings.Fields(ctx.Arguments[roles])
	var resultRoles []string
	var resultRolesToString = ""
	var requestedRoles []string

	ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	for _, tempForBlankIncludedRoleNames := range role {
		requestedRoles = append(requestedRoles, strings.ReplaceAll(tempForBlankIncludedRoleNames, "_", " "))
	}

	findingRoleMsg, _ := ctx.Message.Reply("🔎 요청한 역할(들)을 검증하는 중입니다...")

	for _, requestedRole := range requestedRoles {
		for _, currentRole := range serverRoles {
			if currentRole.Name == requestedRole {
				resultRoles = append(resultRoles, currentRole.ID)
				resultRolesToString = resultRolesToString + "**``" + currentRole.Name + "``**, "
			}
		}
	}

	resultRolesToString = strings.TrimSuffix(resultRolesToString, ", ")

	if resultRolesToString == "" {
		ctx.Session.ChannelMessageEdit(findingRoleMsg.ChannelID, findingRoleMsg.ID, ">>> ❌ 신청한 역할중 찾을 수 있는 역할이 없습니다. 요청을 기각합니다.\n만약 역할 이름에 공백이 포함되어 있다면, 공백은 ``_``로 변경하고 다시 시도해주세요.")
		return nil
	} else {
		confirmMsg, _ := ctx.Session.ChannelMessageEdit(findingRoleMsg.ChannelID, findingRoleMsg.ID, ">>> ✅ 다음 역할들을 찾았습니다: "+resultRolesToString+"\n⚠️ 혹시 추가되지 않은 역할이 있다면, 공백은 ``_``로 변경하고 다시 시도해주세요.\nℹ️ 계속 진행하려면 ⭕, 요청을 취하하라면 ❌ 이모티콘을 추가해주세요.\n이 메세지의 반응은 신청자에게만 유효합니다.\n\n||신청자: "+requester.User.ID+"\n"+"신청한 역할: ឵"+resultRolesToString+"||")
		ctx.Session.MessageReactionAdd(confirmMsg.ChannelID, confirmMsg.ID, "⭕")
		ctx.Session.MessageReactionAdd(confirmMsg.ChannelID, confirmMsg.ID, "❌")
		time.Sleep(time.Second * 30)
		ctx.Session.MessageReactionsRemoveAll(confirmMsg.ChannelID, confirmMsg.ID)
		ctx.Session.ChannelMessageEdit(confirmMsg.ChannelID, confirmMsg.ID, "❌ 요청시간이 만료되어 취하되었습니다.")
		time.Sleep(time.Second * 10)
		_ = ctx.Session.ChannelMessageDelete(confirmMsg.ChannelID, confirmMsg.ID)
	}

	return nil
}
