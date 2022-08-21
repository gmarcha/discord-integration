package cmd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v1/types"
)

var (
	BasicFile       *discordgo.ApplicationCommand
	BasicFileHandle types.CmdHandle
)

func init() {

	BasicFile = &discordgo.ApplicationCommand{
		Name:        "basic-file",
		Description: "Basic command with files",
	}

	BasicFileHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hey there! Congratulations, you just executed your first slash command with a file in the response",
				Files: []*discordgo.File{
					{
						ContentType: "text/plain",
						Name:        "test.txt",
						Reader:      strings.NewReader("Hello Discord!!"),
					},
				},
			},
		})
	}
}
