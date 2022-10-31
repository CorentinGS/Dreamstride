package commands

import (
	"github.com/bwmarrin/discordgo"
)

func InfoCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "/info -> Returns bot commands\n" +
					"/ping -> Returns the latency of the bot\n" +
					"/get-version -> Returns the version of the bot\n" +
					"/addrole arg user -> Adds a role to a user\n" +
					"/rmerole arg user -> Removes a role from a user\n" +
					"/ban user -> Bans a user\n" +
					"/purge arg -> Deletes a number of messages\n" +
					"/mute user time -> Mutes a user for a certain amount of time\n",
			},
		})
	}
}
