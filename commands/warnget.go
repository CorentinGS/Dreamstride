package commands

import (
	"Dreamstride/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func getWarn(s *discordgo.Session, i *discordgo.InteractionCreate) int {
	user = i.ApplicationCommandData().Options[0].UserValue(s)
	return warnedUserMap[user.ID]
}
func WarnGetCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		number := getWarn(s, i)
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Total Warns",
						Description: fmt.Sprintf("User %s has %d warns", user.Username, number),
					},
				},
			},
		})
		utils.Log(i.Member.User.Username, i.Member.User.ID, "warn command called")
	}
}
