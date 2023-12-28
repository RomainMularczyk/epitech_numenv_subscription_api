package models

type Session struct {
  Id             string `json:"id" binding:"required"`
	Name           string `json:"name" binding:"required"`
  Speaker        string `json:"speaker" binding:"required"`
  Date           string `json:"date" binding:"required"`
  Type           string `json:"type" binding:"required"`
  NumSubscribers string `json:"num_subscribers" binding:"required"`
}

func (m *Session) GetID() string {
	return m.Id
}

func (m *Session) SetID(id string) {
	m.Id = id
}
