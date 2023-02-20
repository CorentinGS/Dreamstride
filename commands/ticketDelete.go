package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func getName(channel *discordgo.Channel) string {
	return channel.Name
}

func TicketDeleteCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		targetChannel := i.ChannelID
		tmp, _ := s.State.Channel(targetChannel)
		channelName := getName(tmp)
		if strings.Contains(channelName, "ticket-") {
			_, _ = s.ChannelDelete(targetChannel)
		}
	}
}
