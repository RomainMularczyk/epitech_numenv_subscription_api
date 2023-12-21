package models

type ModelInterface interface {
	GetID() string
	SetID(id string)
}

type Model struct {
	Id string `json:"id"`
}

func (m *Model) GetID() string {
	return m.Id
}

func (m *Model) SetID(id string) {
	m.Id = id
}
