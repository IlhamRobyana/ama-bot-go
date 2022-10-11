package commands

import (
	"log"
	"time"

	saucenao "github.com/GenDoNL/saucenao-go"
	"github.com/ilhamrobyana/ama-bot-go/configs"
	"github.com/lus/dgc"
)

var (
	SaucenNaoAPIKey string
)

func sauceCommands(router *dgc.Router, cfg *configs.Config) {
	SaucenNaoAPIKey = cfg.Saucenao.APIKey
	getSauce(router)
}

func getSauce(router *dgc.Router) {
	router.RegisterCmd(&dgc.Command{
		Name: "sauce",
		Aliases: []string{
			"sauce",
		},
		Description: "Responds the source of the image",
		Usage:       "sauce",
		Example:     "sauce",
		Flags:       []string{},
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{},
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("You are being rate limited!")
		}),
		Handler: getSauceHandler,
	})
}

func getSauceHandler(ctx *dgc.Ctx) {
	client := saucenao.New(SaucenNaoAPIKey)
	if len(ctx.Event.Attachments) == 0 {
		ctx.RespondText("Please include an image with the command")
	}
	result, err := client.FromURL(ctx.Event.Attachments[0].URL)
	if err != nil {
		log.Fatal(err)
	}

	if len(result.Data) == 0 {
		ctx.RespondText("no sauce found")
	}
	ctx.RespondText(result.Data[0].Data.ExtUrls[0])
}
