package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
)

func BanCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Get the user ID from the interaction data
		userID := i.ApplicationCommandData().Options[0].UserValue(s).ID
		var reason string
		if len(i.ApplicationCommandData().Options) >= utils.OptionsLenWithReason {
			reason = i.ApplicationCommandData().Options[1].StringValue()
		} else {
			reason = ""
		}

		if reason == "" {
			err := s.GuildBanCreate(utils.ServerID, userID, 0)
			if err != nil {
				_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Error banning user: " + err.Error(),
					},
				})
				return
			}
		} else {
			err := s.GuildBanCreateWithReason(utils.ServerID, userID, reason, 0)
			if err != nil {
				_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Error banning user: " + err.Error(),
					},
				})
				return
			}
		}
		utils.Log(i.Member.User.Username, i.Member.User.ID, "Banned user "+userID+" for reason: "+reason)
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "User has been banned",
			},
		})
	}
}
