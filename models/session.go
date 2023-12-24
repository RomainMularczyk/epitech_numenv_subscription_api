package models

type Session struct {
	Id             string `json:"id"`
	Name           string `json:"name" binding:"required"`
	NumSubscribers string `json:"num_subscribers"`
}

func (m *Session) GetID() string {
	return m.Id
}

func (m *Session) SetID(id string) {
	m.Id = id
}
