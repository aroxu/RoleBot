package initializer

import (
	"B1ackAnge1/RoleBot/commands/general/addrole"
	"B1ackAnge1/RoleBot/commands/general/delrole"
	"B1ackAnge1/RoleBot/commands/general/help"
	"B1ackAnge1/RoleBot/commands/general/ping"
	"B1ackAnge1/RoleBot/handler"
)

//InitCommands inits command structures
func InitCommands() {
	// initializer command map
	handler.InitCommands()
	// register commands
	ping.Initialize()
	help.Initialize()
	addrole.Initialize()
	delrole.Initialize()
}