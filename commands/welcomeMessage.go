package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
)

func WelcomeImageCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		utils.WelcomeLink = i.ApplicationCommandData().Options[0].StringValue()
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title:       "Welcome Image",
						Description: "The welcome image has been set to: ",
						Color:       utils.GREEN,
					},
				},
			},
		})
		embed := &discordgo.MessageEmbed{
			Title: "Welcome to the server !<a:DS_strawberry:1081335566265761842>",
			Description: "︶︶︶✿︶♡︶꒷꒦︶︶✿︶︶\n" +
				"<a:DS_heart:1081315143398461461> : <#955192188776616081>\n" +
				"<a:DS_heart:1081315143398461461> : <#955192188776616085>\n" +
				"<a:DS_heart:1081315143398461461> : <#955192188776616082>\n" +
				"<a:DS_heart:1081315143398461461> : <#995805735235620864>\n" +
				"<a:DS_heart:1081315143398461461> : <#995823153567768576>\n" +
				"︶︶︶✿︶♡︶꒷꒦︶︶✿︶︶\n" +
				"We hope you enjoy your stay !<a:DS_FloatingHeart:955580787909087322>\n",
			Color: utils.SALMON,
			Image: &discordgo.MessageEmbedImage{
				URL: utils.WelcomeLink,
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: utils.WelcomeIcon,
			},
		}
		_, _ = s.ChannelMessageSendEmbed(i.ChannelID, embed)
	}
}
