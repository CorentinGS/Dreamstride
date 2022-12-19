package main

import (
	"Dreamstride/commands"
	_ "Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

var (
	TOKEN string
)

func getVar() {
	TOKEN = os.Getenv("DISCORD_TOKEN")
	if TOKEN == "" {
		log.Fatal("No token found")
	}
}

func main() {
	getVar()
	discord, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatal("Error while creation the session ", err)
	}
	discord.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name: "Dreamstride",
			Type: discordgo.ActivityTypeWatching,
		},
	}
	err = discord.Open()
	if err != nil {
		log.Fatal("Error while opening the session ", err)
	}
	commandHandlers := commands.GetCommandHandlers()
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}

	})

	appCommands := commands.GetCommands()
	_, err = discord.ApplicationCommandBulkOverwrite(discord.State.User.ID, discord.State.Guilds[0].ID, appCommands)
	if err != nil {
		log.Panicf("Error overwriting commands: %v", err)
	}
	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			log.Fatal("Error while closing the session ", err)
		}
	}(discord)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Connected as ", discord.State.User.Username)
	log.Println("----------------------")
	log.Println("Starting logs")
	log.Println("Press CTRL+C to exit")
	<-stop

	if true {
		log.Println("Removing commands...")
		_, err := discord.ApplicationCommandBulkOverwrite(discord.State.User.ID, "955192188332023819", nil)
		if err != nil {
			log.Panicf("Cannot delete a command: %v", err)
		}
		log.Println("Removed commands")
	}
	log.Println("Gracefully shutting down.")

}
