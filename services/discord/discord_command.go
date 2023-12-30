package discord

import (
	"fmt"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/services"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DiscordUserRegistrationCommand(
  discordClient *discordgo.Session,
) {
  appCommand := []*discordgo.ApplicationCommand {
    {
      Name: "register",
      Description: "Register a new subscriber.",
      Options: []*discordgo.ApplicationCommandOption {
        {
          Name: "key",
          Description: "The key provided to register to a session.",
          Type: discordgo.ApplicationCommandOptionString,
        },
      },
    },
  }

  _, err := discordClient.ApplicationCommandBulkOverwrite(
    os.Getenv("DISCORD_APP_ID"),
    os.Getenv("DISCORD_GUILD_ID"),
    appCommand,
  )
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not create the Discord application command.",
    )
  }

  discordClient.AddHandler(registerSubscriberDiscordIdCallback)
}

// Callback function provided to the `.AddHandler` method
// to trigger when a user triggers the `/register` command
func registerSubscriberDiscordIdCallback(
  session *discordgo.Session, 
  interaction *discordgo.InteractionCreate,
) {
  if interaction.Type == discordgo.InteractionApplicationCommand {
    options := interaction.ApplicationCommandData().Options
    for _, opt := range options {
      switch opt.Type {
      case discordgo.ApplicationCommandOptionString:
        // Save Discord Id in the subscribers table
        uniqueStr := opt.StringValue()
        sess, err := services.RegisterDiscordId(
          interaction.Member.User.ID,
          uniqueStr,
        )
        if err != nil {
          sessErr := session.InteractionRespond(
            interaction.Interaction,
            &discordgo.InteractionResponse {
              Type: discordgo.InteractionResponseChannelMessageWithSource,
              Data: &discordgo.InteractionResponseData {
                Content: "Merci d'entrer votre clé communiquée par mail.",
              },
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
        err = session.InteractionRespond(
          interaction.Interaction,
          &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData {
              Content: fmt.Sprintf(
                `Vous êtes inscrit à la session : **%s** - *%s*. Elle aura lieu le %s.`,
                sess.Speaker, sess.Name, sess.Date,
              ),
            },
          },
        )
        session.GuildMemberRoleAdd(
          os.Getenv("DISCORD_GUILD_ID"),
          interaction.Member.User.ID,
          sess.DiscordRoleId,
        )
        session.GuildMemberNickname(
          os.Getenv("DISCORD_GUILD_ID"),
          interaction.Member.User.ID,
          fmt.Sprintf(
            "%s (%s%v)",
            interaction.Member.User.Username,
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
}
