package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilhamrobyana/ama-bot-go/router"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token      string
	AvatarFile string
	AvatarURL  string
	Bot        *discordgo.Session
	BotID      string
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env : %v", err))
	}

	Token = os.Getenv("TOKEN")
	Bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	// Bot.AddHandler(messageCreate)

	// Bot.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	err = Bot.Open()
	if err != nil {
		panic(fmt.Sprintf("error opening connection, %v", err))
	}

	defer func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
		Bot.Close()
	}()

	router.Init(Bot)
}

// func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if m.Author.ID == s.State.User.ID {
// 		return
// 	}

// 	if m.Content == "ping" {
// 		s.ChannelMessageSend(m.ChannelID, "Pong!")
// 	}

// 	if m.Content == "pong" {
// 		s.ChannelMessageSend(m.ChannelID, "Ping!")
// 	}
// }
