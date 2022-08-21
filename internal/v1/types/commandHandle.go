package types

import "github.com/bwmarrin/discordgo"

type CmdHandle = func(s *discordgo.Session, i *discordgo.InteractionCreate)
type MapStrCmdHandle = map[string]CmdHandle
