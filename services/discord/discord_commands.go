package discord

import (
	"numenv_subscription_api/errors/logs"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DiscordUserRegistrationCommand(
  discordClient *discordgo.Session,
) {
  appCommands := []*discordgo.ApplicationCommand {
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
    {
      Name: "sessions",
      Description: "List all the sessions available.",
    },
    {
      Name: "my-sessions",
      Description: "List all the sessions that I'm subscribed to.",
    },
  }

  _, err := discordClient.ApplicationCommandBulkOverwrite(
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
// to trigger when a user uses the `/register` command
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
    }
  }
}
