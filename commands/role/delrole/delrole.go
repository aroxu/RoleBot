package delrole

import (
	"B1ackAnge1/RoleBot/extensions/permissions"
	"B1ackAnge1/RoleBot/handler"
	"B1ackAnge1/RoleBot/utils"
	"strings"
)

func Initialize() {
	handler.AddCommand(
		handler.Command{
			Run:                  run,
			Names:                []string{"delrole"},
			RequiredArgumentType: []string{target, roles},
			Usage:                map[string]string{"필요한 권한":"**``역할 관리``**", "설명":"``멘션된 유저로부터 역할을 박탈합니다.``", "사용법": "```css\n?!delrole <@mention> rolename1 rolename2 rolename3 ...```"},
		},
	)
}

const (
	target = "멘션"
	roles = "role1, role2, role3, ..."
)

func run(ctx handler.CommandContext) error {
	checkPermissionResult, _ := utils.MemberHasPermission(ctx.Session, ctx.Message.GuildID, ctx.Message.Author.ID, permissions.ADMINISTRATOR)

	if !checkPermissionResult {
		ctx.Message.Reply("❌ 이 명령어를 실행하기 위해서는 관리자 권한이 필요합니다.")
		return nil
	}

	if len(ctx.Arguments[target]) == 0 {
		ctx.Message.Reply("❌ 박탈할 대상을 기재하시고 다시 시도해주세요.")
		return nil
	}

	if len(ctx.Arguments[roles]) == 0 {
		ctx.Message.Reply("❌ 박탈할 역할을 기재하시고 다시 시도해주세요.")
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

	ctx.Message.Reply("🔎 다음 역할(들)을 찾고 박탈하는 중입니다: " + resultRolesToString)

	for _, roleToAdd := range resultRoles {
		errAddRole := ctx.Session.GuildMemberRoleRemove(guild, target, roleToAdd)
		if errAddRole != nil {
			ctx.Message.Reply("❌ 다음 역할을 박탈할 수 없습니다: ```json\n" + errAddRole.Error() + "```")
		}
	}

	ctx.Message.Reply("ℹ️ 혹시 박탈되지 않은 역할이 있다면, 공백은 ``_``로 변경하고 다시 시도해주세요.")
	return nil
}