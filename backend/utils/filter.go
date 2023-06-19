package utils

import (
	"strings"

	"github.com/viniciusataide/velozient-challenge-go/domain"
)

func Filter(list []domain.PasswordCard, filter string) []domain.PasswordCard {
	var cards []domain.PasswordCard
	for _, card := range list {
		if strings.Contains(strings.ToLower(card.Name), strings.ToLower(filter)) {
			cards = append(cards, card)
		}
	}
	return cards
}
