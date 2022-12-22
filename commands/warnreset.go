package commands

import (
	"Dreamstride/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
)

func resetWarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user = i.ApplicationCommandData().Options[0].UserValue(s)
	warnedUserMap[user.ID] = 0
	warnedUserCache.Set("warnedUserMap", warnedUserMap, cache.NoExpiration)
}
func WarnResetCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		resetWarn(s, i)
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Reset Warn",
						Description: fmt.Sprintf("User %s has been reset to 0 warns", user.Username),
					},
				},
			},
		})
		utils.Log(i.Member.User.Username, i.Member.User.ID, "warn command called")
	}
}
