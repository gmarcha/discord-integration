package cmd

import (
	"context"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v2/notion"
	"github.com/gmarcha/discord-integration/internal/v2/types"
	"github.com/jomei/notionapi"
)

var (
	Event       *discordgo.ApplicationCommand
	EventHandle types.CmdHandle
)

func init() {

	integerOptionMinValue := 0.0

	Event = &discordgo.ApplicationCommand{
		Name:        "event",
		Description: "Manage events",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "list",
				Description: "List all events",
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "create",
				Description: "Create an event",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "category",
						Description: "Category",
						Required:    true,
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  "Exam",
								Value: "exam",
							},
							{
								Name:  "Defense",
								Value: "defense",
							},
							{
								Name:  "Meeting",
								Value: "meeting",
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "date",
						Description: "Date (dd/mm/yy)",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "start-time",
						Description: "Start time (hh:mm)",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "end-time",
						Description: "End time  (hh:mm)",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "end-date",
						Description: "End date (dd/mm/yy)",
					},
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "users",
						Description: "Number of users",
						MinValue:    &integerOptionMinValue,
						MaxValue:    100,
					},
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "wallets",
						Description: "Wallet reward",
						MinValue:    &integerOptionMinValue,
						MaxValue:    1000000,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "read",
				Description: "Read an event",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "id",
						Description: "Event ID",
						Required:    true,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "update",
				Description: "Update an event",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "id",
						Description: "Event ID",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "name",
						Description: "Name",
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "category",
						Description: "Category",
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  "Exam",
								Value: "exam",
							},
							{
								Name:  "Defense",
								Value: "defense",
							},
							{
								Name:  "Meeting",
								Value: "meeting",
							},
						},
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "start-date",
						Description: "Start date (dd/mm/yy)",
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "end-date",
						Description: "End date (dd/mm/yy)",
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "start-time",
						Description: "Start time (hh:mm)",
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "end-time",
						Description: "End time  (hh:mm)",
					},
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "users",
						Description: "Number of users",
						MinValue:    &integerOptionMinValue,
						MaxValue:    100,
					},
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "wallets",
						Description: "Wallet reward",
						MinValue:    &integerOptionMinValue,
						MaxValue:    1000000,
					},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Name:        "delete",
				Description: "Delete an event",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "id",
						Description: "Event ID",
						Required:    true,
					},
				},
			},
		},
	}

	eventDatabaseID := notionapi.DatabaseID(os.Getenv("NOTION_EVENTS_DB_ID"))

	EventHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		content := ""

		switch options[0].Name {
		case "list":
			content = "List subcommand."
			res, err := notion.Client.Database.Query(context.Background(), eventDatabaseID, nil)
			if err != nil {
				log.Println(err)
			} else {
				for _, page := range res.Results {
					log.Println(page.Properties["Name"])
				}
			}
		case "create":
			content = "Create subcommand."
		case "read":
			content = "Read subcommand."
		case "update":
			content = "Update subcommand."
		case "delete":
			content = "Delete subcommand."
		default:
			content = "Not implemented"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	}
}
