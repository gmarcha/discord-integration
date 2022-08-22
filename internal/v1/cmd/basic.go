package cmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v2/types"
)

var (
	Basic       *discordgo.ApplicationCommand
	BasicHandle types.CmdHandle
)

func init() {

	Basic = &discordgo.ApplicationCommand{
		Name:        "basic",
		Description: "Basic command",
	}

	BasicHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command",
			},
		})
	}
}
