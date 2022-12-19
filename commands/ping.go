package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
)

//return the latency of the bot
func PingCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		utils.Log(i.Member.User.Username, i.Member.User.ID, "Ping command called")
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: s.HeartbeatLatency().String(),
			},
		})
	}
}
