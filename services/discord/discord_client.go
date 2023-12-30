package discord

import (
	"os"
	"os/signal"
	"numenv_subscription_api/errors/logs"

	"github.com/bwmarrin/discordgo"
)

func DiscordClient() error {
  // Provide Discord bot connection token
  session, err := discordgo.New(
    "Bot " + os.Getenv("DISCORD_BOT_TOKEN"),
  )
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not connect to Discord server.",
    )
    return err
  }

  // Providing Discord bot permissions
  session.Identify.Intents = discordgo.IntentsAll

  // Connection Discord bot
  err = session.Open()
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not connect Discord bot.",
    )
    return err
  }
  logs.Output(logs.INFO, "Discord bot connected.")

  // Adding Discord bot commands
  DiscordUserRegistrationCommand(session)

  // Catching Ctrl + C signal to shutdown Discord bot
  logs.Output(logs.INFO, "Press Ctrl+C to stop Discord bot.")
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt)
  <- stop
  

  return nil
}
