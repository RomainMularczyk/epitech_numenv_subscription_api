package discord

import (
	"fmt"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/repositories"
	"time"

	"github.com/bwmarrin/discordgo"
)

// List all sessions a subcriber is subscribed to
func ListMySessions(
  s *discordgo.Session,
  i *discordgo.InteractionCreate,
) {
  // Get subscriber Id
  subscriber, err := repositories.GetSubscriberByDiscordId(i.Member.User.ID)
  if err != nil {
    sessErr := s.InteractionRespond(
      i.Interaction,
      &discordgo.InteractionResponse {
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData {
          Content: "Une erreur est survenue en tentant de récupérer la liste des sessions auxquelles vous êtes inscrit.e.",
        },
      },
    )
    if sessErr != nil {
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
    sessErr := s.InteractionRespond(
      i.Interaction,
      &discordgo.InteractionResponse {
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData {
          Content: "Une erreur est survenue en tentant de récupérer la liste des sessions auxquelles vous êtes inscrit.e.",
        },
      },
    )
    if sessErr != nil {
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
    date, err := time.Parse(time.RFC3339, session.Date)
    if err != nil {
      logs.Output(
        logs.ERROR,
        "Could not parse the date.",
      )
    }

    listSessions += fmt.Sprintf(
      "- **%s** - *%s* (Date : %s)\n",
      session.Speaker,
      session.Name,
      date.Format("DD-MM-YYYY"),
    )
  }

  sessErr := s.InteractionRespond(
    i.Interaction,
    &discordgo.InteractionResponse {
      Type: discordgo.InteractionResponseChannelMessageWithSource,
      Data: &discordgo.InteractionResponseData {
        Content: listSessions,
      },
    },
  ) 
  if sessErr != nil {
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
  sessions, err := repositories.GetAllSessions()

  listSessions := ""
  for _, session := range sessions {
    listSessions += fmt.Sprintf(
      "- **%s** - *%s* (Date : %s)\n",
      session.Speaker,
      session.Name,
      session.Date,
    )
  }

  if err != nil {
    sessErr := s.InteractionRespond(
      i.Interaction,
      &discordgo.InteractionResponse {
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData {
          Content: "Une erreur est survenue en tentant de récupérer la liste des sessions.",
        },
      },
    )
    if sessErr != nil {
      logs.Output(
        logs.ERROR,
        "Could not initiate Discord bot session response.",
      )
      return
    }
    return
  }

  sessErr := s.InteractionRespond(
    i.Interaction,
    &discordgo.InteractionResponse {
      Type: discordgo.InteractionResponseChannelMessageWithSource,
      Data: &discordgo.InteractionResponseData {
        Content: listSessions,
      },
    },
  )
  if sessErr != nil {
    logs.Output(
      logs.ERROR,
      "Could not initiate Discord bot session response.",
    )
  }
}
