package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"sync"
)

func PurgeCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		numMessage := i.ApplicationCommandData().Options[0].IntValue()

		var wg sync.WaitGroup
		wg.Add(int(numMessage))

		var optionalUser *discordgo.User
		if len(i.ApplicationCommandData().Options) > 1 {
			optionalUser = i.ApplicationCommandData().Options[1].UserValue(s)
		}

		messages, _ := s.ChannelMessages(i.ChannelID, int(numMessage), "", "", "")

		for _, message := range messages {
			message := message
			go func(messageID string) {
				defer wg.Done()
				if optionalUser != nil {
					if message.Author.ID == optionalUser.ID {
						_ = s.ChannelMessageDelete(i.ChannelID, messageID)
					}
				} else {
					_ = s.ChannelMessageDelete(i.ChannelID, messageID)
				}
			}(message.ID)
		}

		wg.Wait()
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Deleted %d messages", numMessage),
			},
		})
	}
}
