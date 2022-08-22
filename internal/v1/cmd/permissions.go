package cmd

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v1/types"
)

var (
	Permissions       *discordgo.ApplicationCommand
	PermissionsHandle types.CmdHandle
)

func init() {

	defaultMemberPermissions := int64(discordgo.PermissionManageServer)
	dmPermission := false

	Permissions = &discordgo.ApplicationCommand{
		Name:                     "permissions",
		Description:              "Command for demonstration of default command permissions",
		DefaultMemberPermissions: &defaultMemberPermissions,
		DMPermission:             &dmPermission,
	}

	PermissionsHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		perms, err := s.ApplicationCommandPermissions(s.State.User.ID, i.GuildID, i.ApplicationCommandData().ID)

		var restError *discordgo.RESTError
		if errors.As(err, &restError) && restError.Message != nil && restError.Message.Code == discordgo.ErrCodeUnknownApplicationCommandPermissions {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: ":x: No permission overwrites",
				},
			})
			return
		} else if err != nil {
			panic(err)
		}

		if err != nil {
			panic(err)
		}
		format := "- %s %s\n"

		channels := ""
		users := ""
		roles := ""

		for _, o := range perms.Permissions {
			emoji := "❌"
			if o.Permission {
				emoji = "☑"
			}

			switch o.Type {
			case discordgo.ApplicationCommandPermissionTypeUser:
				users += fmt.Sprintf(format, emoji, "<@!"+o.ID+">")
			case discordgo.ApplicationCommandPermissionTypeChannel:
				allChannels, _ := discordgo.GuildAllChannelsID(i.GuildID)

				if o.ID == allChannels {
					channels += fmt.Sprintf(format, emoji, "All channels")
				} else {
					channels += fmt.Sprintf(format, emoji, "<#"+o.ID+">")
				}
			case discordgo.ApplicationCommandPermissionTypeRole:
				if o.ID == i.GuildID {
					roles += fmt.Sprintf(format, emoji, "@everyone")
				} else {
					roles += fmt.Sprintf(format, emoji, "<@&"+o.ID+">")
				}
			}
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Permissions overview",
						Description: "Overview of permissions for this command",
						Fields: []*discordgo.MessageEmbedField{
							{
								Name:  "Users",
								Value: users,
							},
							{
								Name:  "Channels",
								Value: channels,
							},
							{
								Name:  "Roles",
								Value: roles,
							},
						},
					},
				},
				AllowedMentions: &discordgo.MessageAllowedMentions{},
			},
		})
	}
}
