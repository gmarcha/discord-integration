package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v2/cmd"
	_ "github.com/gmarcha/discord-integration/internal/v2/env"
	"github.com/gmarcha/discord-integration/internal/v2/types"
)

var (
	botToken        string
	guildID         string
	s               *discordgo.Session
	commands        []*discordgo.ApplicationCommand
	commandHandlers types.MapStrCmdHandle
)

func init() {

	botToken = os.Getenv("DISCORD_BOT_TOKEN")
	guildID = os.Getenv("DISCORD_GUILD_ID")
}

func init() {

	var err error

	s, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatalln("Error creating Discord session")
	}
}

func init() {

	commands = []*discordgo.ApplicationCommand{
		cmd.Event,
	}

	commandHandlers = types.MapStrCmdHandle{
		"event": cmd.EventHandle,
	}
}

func main() {

	s.AddHandler(loggedIn)
	s.AddHandler(createMessage)
	s.AddHandler(launchCommand)

	err := s.Open()
	if err != nil {
		log.Println("Error opening connection, ", err)
		return
	}
	defer s.Close()

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc

	log.Println("Removing commands...")
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, guildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	log.Println("Bot is exiting.")
}

func loggedIn(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func createMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}

func launchCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
