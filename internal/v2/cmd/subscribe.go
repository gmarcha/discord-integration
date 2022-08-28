package cmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v2/types"
)

var (
	Subscribe       *discordgo.ApplicationCommand
	SubscribeHandle types.CmdHandle
)

func init() {

	Subscribe = &discordgo.ApplicationCommand{
		Name:        "subscribe",
		Description: "Subscribe to an event",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Name",
				Required:    true,
			},
		},
	}

	SubscribeHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		margs := make([]interface{}, 0, len(options))
		msgformat := "You learned how to use command options! " +
			"Take a look at the value(s) you entered:\n"

		if option, ok := optionMap["name"]; ok {
			margs = append(margs, option.StringValue())
			msgformat += "> string-option: %s\n"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					msgformat,
					margs...,
				),
			},
		})
	}
}
