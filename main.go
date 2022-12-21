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

func sendSupportEmbed(s *discordgo.Session) string {
	const supportChannel = "986013751276884038"
	embed := &discordgo.MessageEmbed{
		Title:       "Support",
		Description: "If you need help with the bot, or have any questions, react to this message with 📩 to open a ticket.",
		Color:       0x00ff00,
	}
	msg, err := s.ChannelMessageSendEmbed(supportChannel, embed)
	if err != nil {
		log.Println("Error while sending embed ", err)
	}

	err = s.MessageReactionAdd(supportChannel, msg.ID, "📩")
	if err != nil {
		log.Println("Error while adding reaction ", err)
	}
	return msg.ID
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
	supportID := sendSupportEmbed(discord)
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		if r.Emoji.Name == "📩" && r.Member.User.ID != s.State.User.ID && r.MessageID == supportID {
			log.Println("test test test")
			channel, err := s.UserChannelCreate(r.Member.User.ID)
			if err != nil {
				log.Println("Error while creating channel ", err)
			}
			_, err = s.ChannelMessageSend(channel.ID, "Test msg please ignore")
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
