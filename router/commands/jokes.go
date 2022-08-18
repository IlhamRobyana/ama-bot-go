package commands

import (
	"strconv"
	"time"

	"github.com/Lukaesebrot/dgc"
)

func jokeCommands(router *dgc.Router) {
	nice(router)
	blazeIt(router)
}

func nice(router *dgc.Router) {
	router.RegisterCmd(&dgc.Command{
		Name: "nice",
		Aliases: []string{
			"nice",
		},
		Description: "Responds with a nice joke",
		Usage:       "nice",
		Example:     "nice",
		Flags:       []string{},
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{},
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("You are being rate limited!")
		}),
		Handler: niceHandler,
	})
}

func blazeIt(router *dgc.Router) {
	router.RegisterCmd(&dgc.Command{
		Name: "blazeit",
		Aliases: []string{
			"blaze_it",
		},
		Description: "Responds with a lit joke",
		Usage:       "blazeit",
		Example:     "blazeit",
		Flags:       []string{},
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{},
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("You are being rate limited!")
		}),
		Handler: blazeitHandler,
	})
}

func niceHandler(ctx *dgc.Ctx) {
	// Respond with the just set custom object
	ctx.RespondText(strconv.Itoa(ctx.CustomObjects.MustGet("niceObject").(int)))
}

func blazeitHandler(ctx *dgc.Ctx) {
	// Respond with the just set custom object
	ctx.RespondText(strconv.Itoa(ctx.CustomObjects.MustGet("blazeitObject").(int)))
}
