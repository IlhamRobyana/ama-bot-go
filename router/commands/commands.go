package commands

import (
	"github.com/ilhamrobyana/ama-bot-go/configs"
	"github.com/lus/dgc"
)

func InitCommands(router *dgc.Router, cfg *configs.Config) {
	jokeCommands(router)
	animalCommands(router, cfg)
}
