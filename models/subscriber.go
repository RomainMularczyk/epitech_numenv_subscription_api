package models

type Subscriber struct {
	Id            string `json:"id"`
	Firstname     string `json:"first_name"`
	Lastname      string `json:"last_name"`
	Email         string `json:"email"`
	Institution   string `json:"institution"`
	EpitechDegree string `json:"epitech_degree"`
	DiscordId     string `json:"discord_id"`
	UniqueStr     string `json:"unique_str"`
}

func (m *Subscriber) GetID() string {
	return m.Id
}

func (m *Subscriber) SetID(id string) {
	m.Id = id
}
