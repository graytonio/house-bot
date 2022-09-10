package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var echo = &Command{
	Definition: &discordgo.ApplicationCommand{
		Name:        "echo",
		Description: "Echo the message back",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "echo-text",
				Description: "Echo Text",
				Required:    true,
			},
		},
	},
	Handler: func(s *discordgo.Session, ic *discordgo.InteractionCreate) {
		options := ic.ApplicationCommandData().Options
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if option, ok := optionMap["echo-text"]; ok {
			s.InteractionRespond(ic.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("You said: %s", option.StringValue()),
				},
			})
		}
	},
}
