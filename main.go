package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"ubar-bot/discord"

	"github.com/bwmarrin/discordgo"
)

func main() {

	b, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the func as a callback for MessageCreate events.
	b.AddHandler(discord.GetWeather)

	err = b.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// In this example, we only care about receiving message events.
	b.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Wait here until CTRL-C or other term signal is received.
	log.Print("Discord bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	b.Close()
}
