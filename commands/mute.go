package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

func timeoutUser(s *discordgo.Session, userID string, until time.Duration) bool {
	// Initialise une session HTTP

	url := "https://discord.com/api/v9/guilds/" + utils.SERVER_ID + "/members/" + userID
	timeout := time.Now().Add(until)

	json := map[string]string{
		"communication_disabled_until": timeout.Format(time.RFC3339),
	}

	_, err := s.RequestWithBucketID("PATCH", url, json,
		discordgo.EndpointGuildMember(utils.SERVER_ID, userID))
	if err != nil {
		return false
	}
	return true
}

func MuteCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		user := i.ApplicationCommandData().Options[0].UserValue(s).ID
		times, _ := strconv.Atoi(i.ApplicationCommandData().Options[1].StringValue())

		//mute the user
		if timeoutUser(s, user, time.Duration(times)*time.Minute) {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "User has been muted for " + string(rune(times)) + " minutes",
				},
			})
		} else {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error muting user",
				},
			})
		}
	}
}
