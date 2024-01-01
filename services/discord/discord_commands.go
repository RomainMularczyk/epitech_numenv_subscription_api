package discord

import (
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/repositories"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DiscordUserRegistrationCommand(
	discordClient *discordgo.Session,
) {
	sessions, err := repositories.GetAllSessions()
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Could not get sessions from the db to create the Discord bot commands.",
		)
		return
	}

	var autoCompleteChoices []*discordgo.ApplicationCommandOptionChoice
	for _, session := range sessions {
		autoCompleteChoices = append(autoCompleteChoices, &discordgo.ApplicationCommandOptionChoice{
			Value: session.Speaker,
			Name:  session.Speaker,
		})
	}

	appCommands := []*discordgo.ApplicationCommand{
		{
			Name:        "register",
			Description: "Register a new subscriber.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "key",
					Description: "The key provided to register to a session.",
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "sessions",
			Description: "List all the sessions available.",
		},
		{
			Name:        "my-sessions",
			Description: "List all the sessions that I'm subscribed to.",
		},
		{
			Name:        "subscribe",
			Description: "Subscribe to session.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "name",
					Description: "The name of the session to subscribe to.",
					Type:        discordgo.ApplicationCommandOptionString,
					Choices:     autoCompleteChoices,
				},
			},
		},
	}

	_, err = discordClient.ApplicationCommandBulkOverwrite(
		os.Getenv("DISCORD_APP_ID"),
		os.Getenv("DISCORD_GUILD_ID"),
		appCommands,
	)
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Could not create the Discord application commands.",
		)
	}

	discordClient.AddHandler(discordInteractionCallback)
}

// Callback function provided to the `.AddHandler` method
// to trigger when a user uses the commands
func discordInteractionCallback(
	session *discordgo.Session,
	interaction *discordgo.InteractionCreate,
) {
	if interaction.Type == discordgo.InteractionApplicationCommand {
		data := interaction.ApplicationCommandData()
		switch data.Name {
		case "register":
			RegisterSubscriber(session, interaction)
		case "sessions":
			ListSessions(session, interaction)
		case "my-sessions":
			ListMySessions(session, interaction)
		case "subscribe":
			SubscribeToSession(session, interaction)
		}
	}
}
