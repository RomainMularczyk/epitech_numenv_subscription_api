package models

type Subscriber struct {
	Id             string `json:"id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Institution    string `json:"institution"`
	Formation      string `json:"formation"`
	OtherFormation string `json:"other_formation"`
}

func (m *Subscriber) GetID() string {
	return m.Id
}

func (m *Subscriber) SetID(id string) {
	m.Id = id
}
