package commands

import (
	"Dreamstride/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"sync"
	"time"
)

const rateLimit = 1 // rate limit of 2 requests per second

type rateLimiter struct {
	tokens int
	tick   *time.Ticker
	mutex  sync.Mutex
}

func newRateLimiter(rate int) *rateLimiter {
	return &rateLimiter{
		tokens: rate,
		tick:   time.NewTicker(time.Second / time.Duration(rate)),
	}
}

func (r *rateLimiter) limit() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.tokens > 0 {
		r.tokens--
		return
	}

	r.mutex.Unlock()
	<-r.tick.C
	r.mutex.Lock()
	r.tokens--
}
func PurgeCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		numMessage := i.ApplicationCommandData().Options[0].IntValue()

		var wg sync.WaitGroup
		wg.Add(int(numMessage))

		var optionalUser *discordgo.User
		if len(i.ApplicationCommandData().Options) > 1 {
			optionalUser = i.ApplicationCommandData().Options[1].UserValue(s)
			utils.Log(i.Member.User.Username, i.Member.User.ID, "purged messages from", optionalUser.Username, optionalUser.ID)
		} else {
			utils.Log(i.Member.User.Username, i.Member.User.ID, "purged messages")
		}
		messages, _ := s.ChannelMessages(i.ChannelID, int(numMessage), "", "", "")

		rateLimiter := newRateLimiter(rateLimit)
		for _, message := range messages {
			message := message
			go func(messageID string) {
				defer wg.Done()
				rateLimiter.limit()
				if optionalUser != nil {
					if message.Author.ID == optionalUser.ID {
						_ = s.ChannelMessageDelete(i.ChannelID, messageID)
					}
				} else {
					_ = s.ChannelMessageDelete(i.ChannelID, messageID)
				}
				time.Sleep(time.Millisecond * 200)
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
