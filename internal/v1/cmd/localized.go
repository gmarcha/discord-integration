package cmd

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v2/types"
)

var (
	Localized       *discordgo.ApplicationCommand
	LocalizedHandle types.CmdHandle
)

func init() {

	Localized = &discordgo.ApplicationCommand{
		Name:        "localized",
		Description: "Localized command. Description and name may vary depending on the Language setting",
		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishUS: "hello-world",
			discordgo.French:    "bonjour-monde",
			discordgo.ChineseCN: "本地化的命令",
		},
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.EnglishUS: "Hello World!",
			discordgo.French:    "Bonjour Monde!",
			discordgo.ChineseCN: "这是一个本地化的命令",
		},
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "localized-option",
				Description: "Localized option. Description and name may vary depending on the Language setting",
				NameLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishUS: "an-option",
					discordgo.French:    "une-option",
					discordgo.ChineseCN: "一个本地化的选项",
				},
				DescriptionLocalizations: map[discordgo.Locale]string{
					discordgo.EnglishUS: "An Option!",
					discordgo.French:    "Une Option!",
					discordgo.ChineseCN: "这是一个本地化的选项",
				},
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name: "First",
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishUS: "First!",
							discordgo.French:    "Premier!",
							discordgo.ChineseCN: "一的",
						},
						Value: 1,
					},
					{
						Name: "Second",
						NameLocalizations: map[discordgo.Locale]string{
							discordgo.EnglishUS: "Second!",
							discordgo.French:    "Second!",
							discordgo.ChineseCN: "二的",
						},
						Value: 2,
					},
				},
			},
		},
	}

	LocalizedHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		responses := map[discordgo.Locale]string{
			discordgo.ChineseCN: "你好！ 这是一个本地化的命令",
		}
		response := "Hi! This is a localized message"
		if r, ok := responses[i.Locale]; ok {
			response = r
		}
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		})
		if err != nil {
			panic(err)
		}
	}
}
