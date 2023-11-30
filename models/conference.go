package models

type Conference struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

func (m *Conference) GetID() string {
	return m.Id
}

func (m *Conference) SetID(id string) {
	m.Id = id
}
