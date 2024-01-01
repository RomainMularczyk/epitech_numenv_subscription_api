package discord

import (
	"fmt"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/services"
	"numenv_subscription_api/utils"
	"os"

	"github.com/bwmarrin/discordgo"
)

func RegisterSubscriber(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
) {
	options := i.ApplicationCommandData().Options
	for _, opt := range options {
		switch opt.Type {
		case discordgo.ApplicationCommandOptionString:

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

			// Save Discord Id in the subscribers table
			uniqueStr := opt.StringValue()
			sess, err := services.RegisterDiscordId(
				i.Member.User.ID,
				uniqueStr,
			)
			if err != nil {
				_, err = s.FollowupMessageCreate(
					i.Interaction,
					false,
					&discordgo.WebhookParams{
						Content: "Merci d'entrer votre clé communiquée par mail.",
					},
				)
				if sessErr != nil {
					logs.Output(
						logs.ERROR,
						"Could not initiate Discord bot session response.",
					)
				}
				return
			}
			// Get subscriber by unique str
			subscriber, err := services.GetSubscriberByUniqueStr(uniqueStr)
			if err != nil {
				return
			}
			// Bot response to user interaction

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

			s.GuildMemberRoleAdd(
				os.Getenv("DISCORD_GUILD_ID"),
				i.Member.User.ID,
				sess.DiscordRoleId,
			)
			s.GuildMemberNickname(
				os.Getenv("DISCORD_GUILD_ID"),
				i.Member.User.ID,
				fmt.Sprintf(
					"%s (%s %c)",
					i.Member.User.Username,
					subscriber.Firstname,
					subscriber.Lastname[0],
				),
			)
			if err != nil {
				fmt.Println(err)
				logs.Output(
					logs.ERROR,
					"Could not initiate bot response.",
				)
			}
		}
	}
}
