package commands

import (
	"github.com/Lukaesebrot/dgc"
	"github.com/ilhamrobyana/ama-bot-go/configs"
)

func InitCommands(router *dgc.Router, cfg *configs.Config) {
	jokeCommands(router)
	animalCommands(router, cfg)
}
