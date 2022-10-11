package commands

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ilhamrobyana/ama-bot-go/configs"
	"github.com/ilhamrobyana/ama-bot-go/logger"
	"github.com/lus/dgc"
)

var (
	NovelAIKey       string
	NovelAIAccessKey string
)

func novelAICommands(router *dgc.Router, cfg *configs.Config) {
	NovelAIKey = cfg.NovelAI.Key
	accessKey, err := novelAILogin(NovelAIKey)
	if err != nil {
		logger.ErrorWithStack(err)
	} else {
		NovelAIAccessKey = accessKey
	}
	generateImage(router)
}

func novelAILogin(key string) (token string, err error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	body := map[string]string{
		"key": key,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	payload := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, "https://api.novelai.net/user/login", payload)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	accessToken := map[string]string{
		"accessToken": "",
	}
	err = json.Unmarshal(resp, &accessToken)
	if err != nil {
		logger.ErrorWithStack(err)
		return
	}
	token = accessToken["accessToken"]
	return
}

func generateImage(router *dgc.Router) {
	router.RegisterCmd(&dgc.Command{
		Name: "novelai",
		Aliases: []string{
			"novelai",
		},
		Description: "Generates an image based on a prompt",
		Usage:       "novelai",
		Example:     "novelai",
		Flags:       []string{},
		IgnoreCase:  true,
		SubCommands: []*dgc.Command{},
		RateLimiter: dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
			ctx.RespondText("You are being rate limited!")
		}),
		Handler: generateImageHandler,
	})
}

func generateImageHandler(ctx *dgc.Ctx) {
	ctx.RespondText("generating image...")
	client := http.Client{
		Timeout: time.Second * 40,
	}

	body := map[string]interface{}{
		"input": "masterpiece, best quality, " + ctx.Event.Message.Content,
		"model": "safe-diffusion",
		"parameters": map[string]interface{}{
			"height":    768,
			"n_samples": 1,
			"sampler":   "k_euler_ancestral",
			"scale":     11,
			"seed":      rand.Intn(4294967295),
			"steps":     28,
			"uc":        "lowres, bad anatomy, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry",
			"ucPreset":  0,
			"width":     512,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		ctx.RespondText("failed to generate image")
		logger.ErrorWithStack(err)
		return
	}

	payload := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, "https://api.novelai.net/ai/generate-image", payload)
	if err != nil {
		ctx.RespondText("failed to generate image")
		logger.ErrorWithStack(err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+NovelAIAccessKey)
	res, err := client.Do(req)
	if err != nil {
		ctx.RespondText("failed to generate image")
		logger.ErrorWithStack(err)
		return
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ctx.RespondText("failed to generate image")
		logger.ErrorWithStack(err)
		return
	}
	splitted := strings.Split(string(resp), "data:")
	if len(splitted) < 2 {
		ctx.RespondText("failed to generate image")
		fmt.Println(splitted)
		return
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(splitted[1]))
	_, err = ctx.Session.ChannelMessageSendComplex(ctx.Event.ChannelID, &discordgo.MessageSend{
		Content: ctx.Event.Author.Mention(),
		File:    &discordgo.File{Name: "image.png", ContentType: "image/png", Reader: reader},
	})
	if err != nil {
		ctx.RespondText("failed to generate image")
		logger.ErrorWithStack(err)
		return
	}
}
