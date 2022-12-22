package commands

import (
	"Dreamstride/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/patrickmn/go-cache"
)

var (
	user            *discordgo.User
	reason          string
	warnedUserMap   map[*discordgo.User]int
	warnedUserCache *cache.Cache
)

func SetWarnedUserMap(m map[*discordgo.User]int) {
	warnedUserMap = m
	warnedUserCache = cache.New(cache.NoExpiration, cache.NoExpiration)
	warnedUserCache.Set("warnedUserMap", warnedUserMap, cache.NoExpiration)
}
func addWarn(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user = i.ApplicationCommandData().Options[0].UserValue(s)
	if len(i.ApplicationCommandData().Options) > 1 {
		reason = i.ApplicationCommandData().Options[1].StringValue()
	}
	warnedUserMap[user]++
	warnedUserCache.Set("warnedUserMap", warnedUserMap, cache.NoExpiration)
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
						Description: fmt.Sprintf("Warned %s for %s \n total warns: %d", user.Username, reason, (*WarnedUser)[user]),
					},
				},
			},
		})
		utils.Log(i.Member.User.Username, i.Member.User.ID, "warn command called")
	}
}
