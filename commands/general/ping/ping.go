package ping

import "github.com/B1ackAnge1/RoleBot/handler"

func Initialize() {
	handler.AddCommand(
		handler.Command{
			Run:   run,
			Names: []string{"ping"},
		},
	)
}

func run(ctx handler.CommandContext) error {
	var _, err = ctx.Message.Reply("Pong!")
	return err
}