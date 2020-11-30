package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lus/dgc"
)

type Cat struct {
	URL string `json:"url"`
}

func animalCommands(router *dgc.Router) {
	randomCat(router)
}

func randomCat(router *dgc.Router) {
	router.RegisterCmd(&dgc.Command{
		Name: "cat",
		Aliases: []string{
			"cat",
		},
		Description: "Responds with a random cat",
		Usage:       "cat",
		Example:     "cat",
		Flags:       []string{},
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{},
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("You are being rate limited!")
		}),
		Handler: randomCatHandler,
	})
}

func randomCatHandler(ctx *dgc.Ctx) {
	url := "https://api.thecatapi.com/v1/images/search"
	apiKey := os.Getenv("CAT_API")
	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("x-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	cat := []Cat{}
	err = json.Unmarshal(body, &cat)
	if err != nil {
		log.Fatal(err)
	}
	ctx.RespondText(cat[0].URL)
}
