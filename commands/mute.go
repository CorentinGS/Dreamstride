package commands

import (
	"Dreamstride/utils"
	"fmt"
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
	resp, err := s.RequestWithBucketID("PATCH", url, json,
		discordgo.EndpointGuildMember(utils.SERVER_ID, userID))
	if err != nil {
		return false
	}
	var tmp string
	err = discordgo.Unmarshal(resp, &tmp)

	fmt.Printf("%s\n", tmp)
	return false
}

func MuteCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		//get the user to mute
		user := i.ApplicationCommandData().Options[0].UserValue(s).ID
		//get the guild id
		//get the time to mute for
		times, _ := strconv.Atoi(i.ApplicationCommandData().Options[1].StringValue())

		//mute the user
		if timeoutUser(s, user, time.Duration(times)*time.Minute) {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "User has been muted",
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
