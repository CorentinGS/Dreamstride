package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
)

func WelcomeImageCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		utils.WELCOME_LINK = i.ApplicationCommandData().Options[0].StringValue()
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Welcome Image",
						Description: "The welcome image has been set to: " + utils.WELCOME_LINK,
						Color:       utils.GREEN,
					},
				},
			},
		})
	}
}
