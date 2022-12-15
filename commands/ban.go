package commands

import (
	"Dreamstride/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func BanCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Get the user ID from the interaction data
		userID := i.ApplicationCommandData().Options[0].UserValue(s).ID
		reason := i.ApplicationCommandData().Options[1].StringValue()

		fmt.Println(userID)
		if reason == "" {
			err := s.GuildBanCreate(utils.SERVER_ID, userID, 0)
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
			err := s.GuildBanCreateWithReason(i.GuildID, userID, reason, 0)
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
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "User has been banned",
			},
		})
	}
}
