package models

type Subscriber struct {
	Id            string  `json:"id"`
	Firstname     string  `json:"firstname" binding:"required" validate:"required"`
	Lastname      string  `json:"lastname" binding:"required" validate:"required"`
	Email         string  `json:"email" binding:"required" validate:"required,email"`
	Institution   string  `json:"institution" binding:"required" validate:"required"`
	EpitechDegree string  `json:"epitech_degree"`
	DiscordId     *string `json:"discord_id"`
}

type SubscriberWithChallenge struct {
	Subscriber
	Challenge
	Altcha *string `json:"altcha" binding:"required" validate:"required"`
}

func FilterOutAltcha(m *SubscriberWithChallenge) Subscriber {
	return Subscriber{
		Id:            m.Id,
		Firstname:     m.Firstname,
		Lastname:      m.Lastname,
		Email:         m.Email,
		Institution:   m.Institution,
		EpitechDegree: m.EpitechDegree,
		DiscordId:     m.DiscordId,
	}
}

func (m *Subscriber) GetID() string {
	return m.Id
}

func (m *Subscriber) SetID(id string) {
	m.Id = id
}
