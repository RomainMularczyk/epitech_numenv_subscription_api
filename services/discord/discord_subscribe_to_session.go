package discord

import (
	"context"
	"fmt"
	dbError "numenv_subscription_api/errors/db"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models"
	"numenv_subscription_api/repositories"
	"numenv_subscription_api/services"
	"numenv_subscription_api/utils"
	"os"
	"reflect"

	"github.com/bwmarrin/discordgo"
)

func SubscribeToSession(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	options := i.ApplicationCommandData().Options
	for _, opt := range options {
		switch opt.Type {
		case discordgo.ApplicationCommandOptionString:
			speaker := opt.StringValue()

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

			sess, err := RegisterSubscriberToNewSession(i.Member.User.ID, speaker)
			if err != nil {
				msg := "Une erreur est survenue en tentant de vous inscrire à la session."
				if reflect.TypeOf(err) == reflect.TypeOf(dbError.AlreadyRegisteredError{}) {
					msg = "Vous êtes déjà inscrit.e à cette session."
				}
				_, err := s.FollowupMessageCreate(
					i.Interaction,
					false,
					&discordgo.WebhookParams{
						Content: msg,
					},
				)
				if err != nil {
					return
				}

				return
			}

			err = s.GuildMemberRoleAdd(
				os.Getenv("DISCORD_GUILD_ID"),
				i.Member.User.ID,
				sess.DiscordRoleId,
			)
			if err != nil {
				fmt.Println(err)
				logs.Output(
					logs.ERROR,
					"Could set role to subscriber.",
				)
			}

			_, err = s.FollowupMessageCreate(
				i.Interaction,
				false,
				&discordgo.WebhookParams{
					Content: fmt.Sprintf(
						`Vous êtes inscrit à la session : **%s** - *%s*. Elle aura lieu le %s.`,
						sess.Speaker, sess.Name, utils.FormatDate(sess.Date),
					),
				},
			)

			if err != nil {
				logs.Output(
					logs.ERROR,
					"Could not initiate bot response.",
				)
			}
		}
	}
}

func RegisterSubscriberToNewSession(discordId string, speaker string) (*models.Session, error) {
	ctx := context.Background()

	// Get subscriber
	subscriber, err := repositories.GetSubscriberByDiscordId(discordId)
	if err != nil {
		return nil, err
	}

	// Add subscriber to session
	sess, _, err := services.SubscribeToSession(ctx, subscriber, speaker)
	if err != nil {
		return nil, err
	}

	return sess, nil
}