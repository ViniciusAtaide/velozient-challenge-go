package domain

import (
	"github.com/rs/xid"
)

type PasswordCard struct {
	Id       xid.ID `json:"id" swaggertype:"string"`
	URL      string `json:"url"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ProvideCards() *[]PasswordCard {
	cards := make([]PasswordCard, 0)
	return &cards
}
