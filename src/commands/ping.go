package commands

import "github.com/bwmarrin/discordgo"

var ping = &Command{
	Definition: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the server",
	},
	Handler: func(s *discordgo.Session, ic *discordgo.InteractionCreate) {
		s.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	},
}
