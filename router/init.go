package router

import (
	"github.com/ilhamrobyana/ama-bot-go/configs"
	"github.com/ilhamrobyana/ama-bot-go/router/commands"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

func Init(bot *discordgo.Session, cfg *configs.Config) {
	router := create(cfg.Discord.Prefix)
	registerMiddleWare(router)
	commands.InitCommands(router, cfg)

	router.Initialize(bot)
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
