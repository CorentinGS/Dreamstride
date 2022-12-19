package commands

import (
	"Dreamstride/utils"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

func timeoutUser(s *discordgo.Session, userID string, until time.Duration) bool {
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
		userName := i.ApplicationCommandData().Options[0].UserValue(s).Username
		times, _ := strconv.Atoi(i.ApplicationCommandData().Options[1].StringValue())
		utils.Log(i.Member.User.Username, i.Member.User.ID, "Mute command called to mute "+userName+" for "+strconv.Itoa(times)+" minutes")
		if timeoutUser(s, user, time.Duration(times)*time.Minute) {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: userName + " has been muted for " + i.ApplicationCommandData().Options[1].StringValue() + " minutes",
				},
			})
		} else {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error muting " + userName,
				},
			})
		}
	}
}
