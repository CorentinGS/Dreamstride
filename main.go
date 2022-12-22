package main

import (
	"Dreamstride/commands"
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"os/signal"
	"strings"
)

var (
	TOKEN      string
	WarnedUser = make(map[string]int)
	c          *cache.Cache
)

const supportID = "1055259004726685807" //"986013751276884038" original testiing id

func getVar() {
	TOKEN = os.Getenv("DISCORD_TOKEN")
	if TOKEN == "" {
		log.Fatal("No token found")
	}
}

func cacheSetup() {
	c = cache.New(cache.NoExpiration, cache.NoExpiration)

	if v, found := c.Get("warnedUserMap"); found {
		WarnedUser = v.(map[string]int)
	} else {
		WarnedUser = make(map[string]int)
	}
}

/*
func sendSupportEmbed(s *discordgo.Session) string {
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
}*/
func main() {
	getVar()
	cacheSetup()

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
	commands.SetWarnedUserMap(WarnedUser)
	commandHandlers := commands.GetCommandHandlers()
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}

	})
	//supportID := sendSupportEmbed(discord)

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
		if r.Emoji.Name == "📩" && r.Member.User.ID != s.State.User.ID && r.MessageID == supportID {
			checkIfTicketExists := utils.CheckIfTicketExists(s, "ticket-"+r.Member.User.Username)
			if checkIfTicketExists {
				channel, _ := s.UserChannelCreate(r.UserID)
				_, _ = s.ChannelMessageSend(channel.ID, "You already have a ticket open!")
			} else {
				name := strings.ToLower("ticket-" + r.Member.User.Username)
				st, err := s.GuildChannelCreateComplex(utils.SERVER_ID, discordgo.GuildChannelCreateData{
					Name:     name,
					Type:     discordgo.ChannelTypeGuildText,
					ParentID: "1055265595697930290", // Support category ID
					PermissionOverwrites: []*discordgo.PermissionOverwrite{
						{
							ID:   "988183933768327238",
							Type: discordgo.PermissionOverwriteTypeRole,
							Deny: discordgo.PermissionViewChannel,
						},
						{
							ID:    "955192188764045400",
							Type:  discordgo.PermissionOverwriteTypeRole,
							Allow: discordgo.PermissionViewChannel | discordgo.PermissionSendMessages | discordgo.PermissionReadMessageHistory | discordgo.PermissionManageChannels,
						},

						{
							ID:    r.Member.User.ID,
							Type:  discordgo.PermissionOverwriteTypeMember,
							Allow: discordgo.PermissionSendMessages | discordgo.PermissionViewChannel | discordgo.PermissionReadMessageHistory,
						},
					},
				})
				if err != nil {
					log.Println("Error while creating channel ", err)
				}
				_, _ = s.ChannelMessageSend(st.ID, "Welcome to your ticket, <@"+r.Member.User.ID+">. Please describe your issue here. A staff member will be with you shortly.")
				channel, _ := s.UserChannelCreate(r.UserID)
				_, _ = s.ChannelMessageSend(channel.ID, "Your ticket has been created. Look for the channel in the server.")
			}
		}
	})

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		embed := &discordgo.MessageEmbed{
			Title: "Welcome to Dreamstride ・Anime ・Social ・Gaming (Revamp) !",
			Description: "Make sure to check these channels out!\n" +
				"<a:DS_watch:1053113846346289192>】 <#955192188776616081>\n" +
				"<a:DS_playingwithhair:1053115074010697800>】 <#955192188776616082>\n" +
				"<a:DS_hug:1053133920637554819>】 <#955192188776616085>\n\n" +
				"<a:DS_glad:1053134867271012362>】 We hope you have an enjoyable experience here at Dreamstride !\n",
			Color: 0xDF73F5,
			Image: &discordgo.MessageEmbedImage{
				URL: utils.WELCOME_LINK,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: utils.WELCOME_ICON,
			},
		}
		_, err = s.ChannelMessageSend(utils.WELCOME_CHAN, "Hey <@"+m.User.ID+">")
		if err != nil {
			log.Println("Error while sending message ", err)
		}
		_, err = s.ChannelMessageSendEmbed(utils.WELCOME_CHAN, embed)
		if err != nil {
			log.Panicln("Error while sending embed ", err)
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
