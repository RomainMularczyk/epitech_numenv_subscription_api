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
      Description: "Registrate a new subscriber.",
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
        sess, err := services.RegisterDiscordId(
          interaction.Member.User.ID,
          opt.StringValue(),
        )
        if err != nil {
          err := session.InteractionRespond(
            interaction.Interaction,
            &discordgo.InteractionResponse {
              Type: discordgo.InteractionResponseChannelMessageWithSource,
              Data: &discordgo.InteractionResponseData {
                Content: "Merci d'entrer votre clé communiquée par mail.",
              },
            },
          )
          if err != nil {
            logs.Output(
              logs.ERROR,
              "Could not initiate bot response.",
            )
          }
        }
        fmt.Println(sess)
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
        if err != nil {
          logs.Output(
            logs.ERROR,
            "Could not initiate bot response.",
          )
        }
      }
    }
  }
}
