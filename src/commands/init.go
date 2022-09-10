package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Definition *discordgo.ApplicationCommand
	Handler    func(*discordgo.Session, *discordgo.InteractionCreate)
}

var commands = map[string]*Command{
	"ping": ping,
	"echo": echo,
}

func LoadCommands(s *discordgo.Session, applicationID string, guildID string) {
	if len(guildID) > 0 {
		log.Println("Reloading Dev Commands")
		registeredCommands, err := s.ApplicationCommands(s.State.User.ID, guildID)
		if err != nil {
			log.Fatalf("Could not fetch registered commands: %v", err)
		}

		for _, v := range registeredCommands {
			if v.ApplicationID == applicationID {
				err := s.ApplicationCommandDelete(s.State.User.ID, guildID, v.ID)
				if err != nil {
					log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
				}
				log.Printf("Deleted '%v' command", v.Name)
			}
		}
	}

	log.Println("Loading Commands...")
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if c, ok := commands[i.ApplicationCommandData().Name]; ok {
			c.Handler(s, i)
		}
	})

	for _, c := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, c.Definition)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", c.Definition.Name, err)
		}
	}
}
