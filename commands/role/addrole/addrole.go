package addrole

import (
	"B1ackAnge1/RoleBot/handler"
	"strings"
)

func Initialize() {
	handler.AddCommand(
		handler.Command{
			Run:                  run,
			Names:                []string{"addrole"},
			RequiredArgumentType: []string{target, roles},
			Usage:                map[string]string{"필요한 권한":"**``역할 관리``**", "설명":"``멘션된 유저에게 역할을 부여합니다.``", "사용법": "```css\n?!addrole <@mention> rolename1 rolename2 rolename3 ...```"},
		},
	)
}

const (
	target = "멘션"
	roles = "role1, role2, role3, ..."
)

func run(ctx handler.CommandContext) error {
	if len(ctx.Arguments[target]) == 0 {
		ctx.Message.Reply("부여할 대상을 기재하시고 다시 시도해주세요.")
		return nil
	}

	if len(ctx.Arguments[roles]) == 0 {
		ctx.Message.Reply("부여할 역할을 기재하시고 다시 시도해주세요.")
		return nil
	}

	guild := ctx.Message.GuildID
	target := strings.ReplaceAll(strings.ReplaceAll(ctx.Arguments[target], "<@!", ""), ">", "")
	role := strings.Fields(ctx.Arguments[roles])
	serverRoles, _ := ctx.Session.GuildRoles(guild)

	var resultRoles []string
	var resultRolesToString = ""
	var requestedRoles []string

	for _, tempForBlankIncludedRoleNames := range role {
		requestedRoles = append(requestedRoles, strings.ReplaceAll(tempForBlankIncludedRoleNames, "_", " "))
	}

	for _, requestedRole := range requestedRoles {
		for _, currentRole := range serverRoles {
			if currentRole.Name == requestedRole {
				resultRoles = append(resultRoles, currentRole.ID)
				resultRolesToString = resultRolesToString + "``" + currentRole.Name + "``, "
			}
		}
	}
	resultRolesToString = strings.TrimSuffix(resultRolesToString, ", ")

	ctx.Message.Reply("다음 역할(들)을 찾고 추가하는 중입니다: " + resultRolesToString)

	for _, roleToAdd := range resultRoles {
		errAddRole := ctx.Session.GuildMemberRoleAdd(guild, target, roleToAdd)
		if errAddRole != nil {
			ctx.Message.Reply("다음 역할을 추가할 수 없습니다: ```json\n" + errAddRole.Error() + "```")
		}
	}

	ctx.Message.Reply("혹시 추가되지 않은 역할이 있다면, 공백은 ``_``로 변경하고 다시 시도해주세요.")
	return nil
}