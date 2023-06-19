package utils

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/rs/xid"
	"github.com/viniciusataide/velozient-challenge-go/domain"
)

const CARDS_LENGTH = 3

func ProvideTestCards(crypto *Crypto) *[]domain.PasswordCard {
	cards := make([]domain.PasswordCard, CARDS_LENGTH)

	for i := 0; i < CARDS_LENGTH; i++ {
		card := &domain.PasswordCard{
			Username: faker.Username(),
			Id:       xid.New(),
			URL:      faker.URL(),
			Name:     faker.Name(),
			Password: crypto.EncryptAES(fmt.Sprintf("password%d", i)),
		}
		cards[i] = *card
	}

	return &cards
}

func Remove(cards []domain.PasswordCard, card domain.PasswordCard) []domain.PasswordCard {
	var index int
	for i, c := range cards {
		if card.Id.String() == c.Id.String() {
			index = i
		}
	}
	return append(cards[:index], cards[index+1:]...)
}

func Update(lst *[]domain.PasswordCard, id string, newPc *domain.PasswordCardUpdate, crypto *Crypto) {
	list := *lst
	for i := range list {
		if list[i].Id.String() == id {
			if list[i].Name != newPc.Name && newPc.Name != "" {
				list[i].Name = newPc.Name
			}
			if list[i].Password != newPc.Password && newPc.Password != "" {
				list[i].Password = crypto.EncryptAES(newPc.Password)
			}
			if list[i].URL != newPc.URL && newPc.URL != "" {
				list[i].URL = newPc.URL
			}
			if list[i].Username != newPc.Username && newPc.Username != "" {
				list[i].Username = newPc.Username
			}
		}
	}
}
