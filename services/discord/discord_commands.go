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
	sessions, err := repositories.GetAllConfirmedSessions()
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
			Description: "Finaliser la première inscription avec la clé unique reçue dans l'email de confirmation.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "key",
					Description: "Clé unique reçue par mail",
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "sessions",
			Description: "Liste des sessions disponibles.",
		},
		{
			Name:        "mes-sessions",
			Description: "Liste toutes les sessions auxquelles tu es inscrit.e.",
		},
		{
			Name:        "subscribe",
			Description: "Inscription rapide à une nouvelle session sans repasser par le formulaire en ligne.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "nom",
					Description: "Nom de l'intervenant de la session à laquelle tu veux t'inscrire.",
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
			"Could not create the Discord application commands. Err: "+err.Error(),
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
		case "mes-sessions":
			ListMySessions(session, interaction)
		case "subscribe":
			SubscribeToSession(session, interaction)
		}
	}
}
