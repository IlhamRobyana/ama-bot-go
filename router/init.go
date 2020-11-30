package router

import (
	"os"

	"github.com/ilhamrobyana/ama-bot-go/router/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
)

func Init(Bot *discordgo.Session) {
	prefix := os.Getenv("PREFIX")
	router := create(prefix)
	registerMiddleWare(router)
	commands.InitCommands(router)

	router.Initialize(Bot)
}

func create(prefix string) *dgc.Router {
	return dgc.Create(&dgc.Router{
		Prefixes: []string{
			prefix,
		},

		BotsAllowed: false,

		Commands: []*dgc.Command{},

		Middlewares: []dgc.Middleware{},

		PingHandler: func(ctx *dgc.Ctx) {
			ctx.RespondText("Pong!")
		},
	})
}

func registerMiddleWare(router *dgc.Router) {
	// Register a simple middleware that injects a custom object
	router.RegisterMiddleware(func(next dgc.ExecutionHandler) dgc.ExecutionHandler {
		return func(ctx *dgc.Ctx) {
			// Inject a custom object into the context
			ctx.CustomObjects.Set("niceObject", 69)
			ctx.CustomObjects.Set("blazeitObject", 420)
			// Call the next execution handler
			next(ctx)
		}
	})
}
