package commands

import "github.com/bwmarrin/discordgo"

func PurgeCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		numMessage := i.ApplicationCommandData().Options[0].IntValue()
		// Verify that an optional user was passed
		var optionalUser *discordgo.User
		if len(i.ApplicationCommandData().Options) > 1 {
			optionalUser = i.ApplicationCommandData().Options[1].UserValue(s)
		}

		messages, _ := s.ChannelMessages(i.ChannelID, int(numMessage), "", "", "")

		for _, message := range messages {
			if optionalUser != nil {
				if message.Author.ID == optionalUser.ID {
					err := s.ChannelMessageDelete(i.ChannelID, message.ID)
					if err != nil {
						_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
							Type: discordgo.InteractionResponseChannelMessageWithSource,
							Data: &discordgo.InteractionResponseData{
								Content: "Error deleting message: " + err.Error(),
							},
						})
						return
					}
				}
			} else {
				err := s.ChannelMessageDelete(i.ChannelID, message.ID)
				if err != nil {
					_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Error deleting message: " + err.Error(),
						},
					})
					return
				}
			}
		}
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Messages have been deleted",
			},
		})
	}
}
