package ping

import "B1ackAnge1/RoleBot/handler"

func Initialize() {
	handler.AddCommand(
		handler.Command{
			Run:                  run,
			Names:                []string{"ping"},
			RequiredArgumentType: []string{commandArg},
			Usage:                map[string]string{"설명":"봇이 살아있나 확인합니다.", "사용법": "```css\n?!ping```"},
		},
	)
}

const (
	commandArg = "없음"
)

func run(ctx handler.CommandContext) error {
	var _, err = ctx.Message.Reply("Pong!")
	return err
}