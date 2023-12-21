package models

type Subscriber struct {
	Id            string `json:"id"`
  Firstname     string `json:"first_name" binding:"required"`
  Lastname      string `json:"last_name" binding:"required"`
  Email         string `json:"email" binding:"required"`
  Institution   string `json:"institution" binding:"required"`
	EpitechDegree string `json:"epitech_degree"`
	DiscordId     string `json:"discord_id"`
	UniqueStr     string `json:"unique_str"`
}

func (m *Subscriber) GetUniqueStr() string {
  return m.Id
}

func (m *Subscriber) SetUniqueStr(uniqueStr string) {
  m.UniqueStr = uniqueStr
}

func (m *Subscriber) GetID() string {
	return m.Id
}

func (m *Subscriber) SetID(id string) {
	m.Id = id
}

