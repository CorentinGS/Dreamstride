package utils

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Log(caller string, id string, messages ...string) {
	pc, _, _, _ := runtime.Caller(1)
	callingFunc := runtime.FuncForPC(pc).Name()
	log.Println(callingFunc, "called by", caller, "with ID", id, messages)
}

func CheckIfTicketExists(s *discordgo.Session, id string) bool {
	channels := GetChannels(s)
	id = strings.ToLower(id)
	for _, channel := range channels {
		if channel.Name == id {
			return true
		}
	}
	return false
}

func GetChannels(s *discordgo.Session) []*discordgo.Channel {
	channels, err := s.GuildChannels(ServerID)
	if err != nil {
		log.Println("Error while getting channels ", err)
	}
	return channels
}

func GetTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
