package commands

import "github.com/bwmarrin/discordgo"

func AddRoleCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Get the role ID from the interaction data
		roleID := i.ApplicationCommandData().Options[0].RoleValue(s, "").ID

		// Get the user ID from the interaction data
		userID := i.ApplicationCommandData().Options[1].UserValue(s).ID

		// Add the role to the user
		err := s.GuildMemberRoleAdd(i.GuildID, userID, roleID)
		if err != nil {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error adding role: " + err.Error(),
				},
			})
			return
		}

		// Respond to the interaction
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Role added",
			},
		})
	}
}

func RmRoleCommand() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Get the role ID from the interaction data
		roleID := i.ApplicationCommandData().Options[0].RoleValue(s, "").ID

		// Get the user ID from the interaction data
		userID := i.ApplicationCommandData().Options[1].UserValue(s).ID

		// Remove the role from the user
		err := s.GuildMemberRoleRemove(i.GuildID, userID, roleID)
		if err != nil {
			_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error removing role: " + err.Error(),
				},
			})
			return
		}

		// Respond to the interaction
		_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Role removed",
			},
		})
	}
}
