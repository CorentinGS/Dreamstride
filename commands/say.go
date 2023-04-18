package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
)

func SayCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Member.User.ID == "813286644652179467" || i.Member.User.ID == "219472739109568518" {
			channel := i.ApplicationCommandData().Options[0].ChannelValue(s).ID
			/*Prepare an embed to send*/
			embed := &discordgo.MessageEmbed{
				Description: i.ApplicationCommandData().Options[1].StringValue(),
				Color:       utils.PINK, // pink
			}
			/*Send the embed*/
			_, _ = s.ChannelMessageSendEmbed(channel, embed)
		}
		_ = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
			},
		}
	}
}
