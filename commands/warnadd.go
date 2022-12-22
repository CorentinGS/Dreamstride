package commands

import (
	"Dreamstride/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	WarnedUser = map[*discordgo.User]int{}
	user       *discordgo.User
	reason     string
)

func addWarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user = i.ApplicationCommandData().Options[0].UserValue(s)
	if len(i.ApplicationCommandData().Options) > 1 {
		reason = i.ApplicationCommandData().Options[1].StringValue()
	}
	WarnedUser[user]++
}
func WarnCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		addWarn(s, i)
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Warn",
						Description: fmt.Sprintf("Warned %s for %s \n total warns: %d", user.Username, reason, WarnedUser[user]),
					},
				},
			},
		})
		utils.Log(i.Member.User.Username, i.Member.User.ID, "warn command called")
	}
}
