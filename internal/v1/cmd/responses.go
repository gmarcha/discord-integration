package cmd

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v2/types"
)

var (
	Responses       *discordgo.ApplicationCommand
	ResponsesHandle types.CmdHandle
)

func init() {

	Responses = &discordgo.ApplicationCommand{
		Name:        "responses",
		Description: "Interaction responses testing initiative",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "resp-type",
				Description: "Response type",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Channel message with source",
						Value: 4,
					},
					{
						Name:  "Deferred response With Source",
						Value: 5,
					},
				},
				Required: true,
			},
		},
	}

	ResponsesHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Responses to a command are very important.
		// First of all, because you need to react to the interaction
		// by sending the response in 3 seconds after receiving, otherwise
		// interaction will be considered invalid and you can no longer
		// use the interaction token and ID for responding to the user's request

		content := ""
		// As you can see, the response type names used here are pretty self-explanatory,
		// but for those who want more information see the official documentation
		switch i.ApplicationCommandData().Options[0].IntValue() {
		case int64(discordgo.InteractionResponseChannelMessageWithSource):
			content =
				"You just responded to an interaction, sent a message and showed the original one. " +
					"Congratulations!"
			content +=
				"\nAlso... you can edit your response, wait 5 seconds and this message will be changed"
		default:
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseType(i.ApplicationCommandData().Options[0].IntValue()),
			})
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
			}
			return
		}

		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseType(i.ApplicationCommandData().Options[0].IntValue()),
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
		if err != nil {
			s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Something went wrong",
			})
			return
		}
		time.AfterFunc(time.Second*5, func() {
			content := content + "\n\nWell, now you know how to create and edit responses. " +
				"But you still don't know how to delete them... so... wait 10 seconds and this " +
				"message will be deleted."
			_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &content,
			})
			if err != nil {
				s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})
				return
			}
			time.Sleep(time.Second * 10)
			s.InteractionResponseDelete(i.Interaction)
		})
	}
}
