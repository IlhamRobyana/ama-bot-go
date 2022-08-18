package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilhamrobyana/ama-bot-go/configs"
	"github.com/ilhamrobyana/ama-bot-go/logger"
	"github.com/ilhamrobyana/ama-bot-go/router"

	"github.com/bwmarrin/discordgo"
)

var (
	config *configs.Config
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize config
	config = configs.Get()

	// Set desired log level
	logger.SetLogLevel(config)

	bot, err := discordgo.New("Bot " + config.Discord.Token)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	err = bot.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening connection, %v", err))
	}

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc
		bot.Close()
	}()

	router.Init(bot, config)
}
