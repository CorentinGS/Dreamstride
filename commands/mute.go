package commands

import (
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func timeoutUser(u string, guildId string, until int64) bool {
	//prepare http request
	request := "https://discord.com/api/guilds/" + guildId + "/members/" + u
	req, err := http.NewRequest("PUT", request, nil)
	if err != nil {
		log.Fatal(err)
	}

	//set headers
	token := os.Getenv("DISCORD_TOKEN")
	req.Header.Set("Authorization", "Bot "+token)
	//convert until to iso format
	timeout := time.Unix(int64(until), 0).Format(time.RFC3339)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("communication_disabled_until", timeout)
	//send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 || resp.StatusCode < 299 {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Panic(err)
			}
		}(resp.Body)
		return true
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}(resp.Body)
	return false

}
func MuteCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		//get the user to mute
		user := i.ApplicationCommandData().Options[0].UserValue(s).ID
		//get the guild id
		guildId := i.GuildID
		//get the time to mute for
		times, _ := strconv.Atoi(i.ApplicationCommandData().Options[1].StringValue())

		//mute the user
		if timeoutUser(user, guildId, int64(times)) {
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
