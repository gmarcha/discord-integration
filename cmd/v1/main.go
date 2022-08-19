package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/gmarcha/discord-integration/internal/v1/commands"
	"github.com/joho/godotenv"
)

var (
	guildID  string
	botToken string
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	guildID = os.Getenv("GUILD_ID")
	botToken = os.Getenv("BOT_TOKEN")
}

func main() {

	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Error creating Discord session, ", err)
		return
	}

	dg.AddHandler(loggedIn)
	dg.AddHandler(messageCreate)
	dg.AddHandler(launchCommand)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection, ", err)
		return
	}
	defer dg.Close()

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.AppCommands))
	for i, v := range commands.AppCommands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, v)
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
		err := dg.ApplicationCommandDelete(dg.State.User.ID, guildID, v.ID)
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
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

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
	if h, ok := commands.AppCommandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}
