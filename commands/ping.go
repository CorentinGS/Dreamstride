package commands

import (
	"github.com/bwmarrin/discordgo"
)

//return the latency of the bot
func PingCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: s.HeartbeatLatency().String(),
			},
		})
	}
}
