package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/graytonio/house-bot/src/commands"
	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken         string
	DiscordApplicationID string
	TestGuildID          string
}

var s *discordgo.Session
var config *Config

func init() {
	godotenv.Load()
	config = &Config{
		DiscordToken:         os.Getenv("DISCORD_TOKEN"),
		DiscordApplicationID: os.Getenv("DISCORD_APPLICATION_ID"),
		TestGuildID:          os.Getenv("TEST_GUILD_ID"),
	}

	var err error
	s, err = discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	s.Identify.Intents = discordgo.IntentsGuildMessages

	err := s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	commands.LoadCommands(s, config.DiscordApplicationID, config.TestGuildID)

	fmt.Println("Bot is now running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	s.Close()
}
