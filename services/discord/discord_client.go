package discord

import (
	"numenv_subscription_api/errors/logs"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DiscordClient() {
  discordClient, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
  if err != nil {
    logs.Output(
      logs.ERROR,
      "Could not connect to Discord server.",
    )
  }

  discordClient.Identify.Intents = discordgo.IntentsAll

  discordClient.GuildMemberRoleAdd("1175146908965159013", "236577940651900928", "1188650394335842344")

}
