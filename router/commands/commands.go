package commands

import (
	"github.com/lus/dgc"
)

func InitCommands(router *dgc.Router) {
	jokeCommands(router)
	animalCommands(router)
}
