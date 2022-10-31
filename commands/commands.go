package commands

import "github.com/bwmarrin/discordgo"

var (
	defaultMemberPermissions int64 = discordgo.PermissionAdministrator
	dmPermissions                  = false
	commands                       = []*discordgo.ApplicationCommand{
		{
			Name:        "get-version",
			Description: "Returns the version of the bot",
		},
		{
			Name:        "ping",
			Description: "Returns the latency of the bot",
		},
		{
			Name:        "info",
			Description: "Return bot commands",
		},
		{
			Name:        "addrole",
			Description: "Adds a role to a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The role to add",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to add the role to",
					Required:    true,
				},
			},
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermissions,
		},
		{
			Name:        "rmerole",
			Description: "Removes a role from a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The role to remove",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user who's role will be removed",
					Required:    true,
				},
			},
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermissions,
		},
		{
			Name:        "ban",
			Description: "Bans a user",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to ban",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    false,
				},
			},
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermissions,
		},
		{
			Name:        "purge",
			Description: "Deletes a number of messages",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "amount",
					Description: "The amount of messages to delete",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user whose messages will be deleted",
					Required:    false,
				},
			},
		},
		{
			Name:        "mute",
			Description: "Mutes a user for a certain amount of time",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to mute",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "time",
					Description: "The amount of time to mute the user for,in minute",
					Required:    true,
				},
			},
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermissions,
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"get-version": VersionCommand(),
		"ping":        PingCommand(),
		"info":        InfoCommand(),
		"addrole":     AddRoleCommand(),
		"rmerole":     RmRoleCommand(),
		"ban":         BanCommand(),
		"purge":       PurgeCommand(),
		"mute":        MuteCommand(),
	}
)

func GetCommands() []*discordgo.ApplicationCommand {
	return commands
}

func GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return commandHandlers
}
