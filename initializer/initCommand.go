package initializer

import (
	"github.com/B1ackAnge1/RoleBot/commands/general/ping"
	"github.com/B1ackAnge1/RoleBot/handler"
)

//InitCommands inits command structures
func InitCommands() {
	// initializer command map
	handler.InitCommands()
	// register commands
	ping.Initialize()
}