package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
)

func InfoCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					{
						Title: "Info",
						Description: "/info - Shows this message\n" +
							"/ping - Shows the latency of the bot\n" +
							"/version - Shows the current version of the bot\n" +
							"/ban - Bans a user\n" +
							"/mute - Mutes a user for x minutes\n" +
							"/purge - Deletes a number of messages\n" +
							"/addrole - Adds a role to a user\n" +
							"/rmerole - Removes a role from a user\n",
						Author: &discordgo.MessageEmbedAuthor{
							Name:    s.State.User.Username,
							IconURL: s.State.User.AvatarURL(""),
						},
					},
				},
			},
		})
		utils.Log(i.Member.User.Username, i.Member.User.ID, "Info command called")
	}
}
