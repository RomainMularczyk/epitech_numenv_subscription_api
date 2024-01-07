package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/repositories"
	"numenv_subscription_api/utils"
)

// List all sessions a subcriber is subscribed to
func ListMySessions(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	// Initiate Discord bot session response
	sessErr := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		},
	)
	if sessErr != nil {
		logs.Output(
			logs.ERROR,
			"Could not initiate Discord bot session response.",
		)
		return
	}

	// Get subscriber Id
	subscriber, err := repositories.GetSubscriberByDiscordId(i.Member.User.ID)
	if err != nil {
		_, err = s.FollowupMessageCreate(
			i.Interaction,
			false,
			&discordgo.WebhookParams{
				Content: "Une erreur est survenue en tentant de récupérer la liste des sessions auxquelles vous êtes inscrit.e.",
			},
		)
		if err != nil {
			logs.Output(
				logs.ERROR,
				"Could not initiate Discord bot session response.",
			)
			return
		}
		return
	}

	// Get all sessions a subscriber is registered to
	sessions, err := repositories.GetAllSessionsBySubscriberId(subscriber.Id)
	if err != nil {
		_, err = s.FollowupMessageCreate(
			i.Interaction,
			false,
			&discordgo.WebhookParams{
				Content: "Une erreur est survenue en tentant de récupérer la liste des sessions auxquelles vous êtes inscrit.e.",
			},
		)
		if err != nil {
			logs.Output(
				logs.ERROR,
				"Could not initiate Discord bot session response.",
			)
			return
		}
		return
	}

	listSessions := "Vous êtes inscrit.e aux sessions suivantes :\n"
	for _, session := range sessions {
		listSessions += fmt.Sprintf(
			"- **%s** - *%s* (Date : %s)\n",
			session.Speaker,
			session.Name,
			utils.FormatDate(session.Date),
		)
	}

	_, err = s.FollowupMessageCreate(
		i.Interaction,
		false,
		&discordgo.WebhookParams{
			Content: listSessions,
		},
	)
	if err != nil {
		logs.Output(
			logs.ERROR,
			"Could not initiate Discord bot session response.",
		)
	}
}

// List all available sessions
func ListSessions(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	// Initiate Discord bot session response
	sessErr := s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		},
	)
	if sessErr != nil {
		logs.Output(
			logs.ERROR,
			"Could not initiate Discord bot session response.",
		)
		return
	}

	//Get all sessions
	sessions, err := repositories.GetAllConfirmedSessions()

	listSessions := ""
	for _, session := range sessions {
		listSessions += fmt.Sprintf(
			"- **%s** - *%s* (Date : %s)\n",
			session.Speaker,
			session.Name,
			utils.FormatDate(session.Date),
		)
	}

	if err != nil {
		_, err = s.FollowupMessageCreate(
			i.Interaction,
			false,
			&discordgo.WebhookParams{
				Content: "Une erreur est survenue en tentant de récupérer la liste des sessions.",
			},
		)
		if err != nil {
			logs.Output(
				logs.ERROR,
				"Could not initiate Discord bot session response.",
			)
			return
		}
		return
	}

	_, err = s.FollowupMessageCreate(
		i.Interaction,
		false,
		&discordgo.WebhookParams{
			Content: listSessions,
		},
	)

	if err != nil {
		logs.Output(
			logs.ERROR,
			"Could not initiate Discord bot session response.",
		)
	}
}
