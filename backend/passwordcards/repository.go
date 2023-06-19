package passwordcards

import (
	"errors"

	"github.com/rs/xid"
	"github.com/viniciusataide/velozient-challenge-go/domain"
	"github.com/viniciusataide/velozient-challenge-go/utils"
)

type Repository struct {
	PasswordCards *[]domain.PasswordCard
	crypto        *utils.Crypto
}

func ProvideRepository(PasswordCards *[]domain.PasswordCard, crypto *utils.Crypto) *Repository {
	return &Repository{PasswordCards, crypto}
}

func (r *Repository) List(pagination *utils.Pagination, filter string) []domain.PasswordCard {
	var cards []domain.PasswordCard

	if filter != "" {
		cards = utils.Filter(*r.PasswordCards, filter)
	} else {
		cards = *r.PasswordCards
	}

	if pagination.Size > len(cards) {
		pagination.Size = len(cards)
	}
	if pagination.Offset > len(cards) {
		pagination.Offset = len(cards)
	}
	if pagination.Size == 0 {
		pagination.Size = len(cards)
	}

	cards = cards[pagination.Offset : pagination.Offset+pagination.Size]

	decryptedCards := make([]domain.PasswordCard, len(cards))

	for i := range cards {
		decryptedCards[i] = cards[i]
		decryptedCards[i].Password = r.crypto.DecryptAES(cards[i].Password)
	}

	return decryptedCards
}

func (r *Repository) Create(card domain.PasswordCardCreate) (*domain.PasswordCard, error) {
	existing := r.List(&utils.Pagination{Size: 10, Offset: 0}, card.Name)

	if len(existing) > 0 {
		return nil, errors.New("Constraint: PasswordCard name already exists")
	}
	xid := xid.New()

	// SQL operation
	newCard := domain.PasswordCard{
		Id:       xid,
		URL:      card.URL,
		Name:     card.Name,
		Username: card.Username,
		Password: r.crypto.EncryptAES(card.Password),
	}

	*r.PasswordCards = append(*r.PasswordCards, newCard)
	return &newCard, nil
}

func (r *Repository) Delete(id string) error {
	existing := r.Get(id)

	if existing == nil {
		return errors.New("NotFound: PasswordCard")
	}

	*r.PasswordCards = utils.Remove(*r.PasswordCards, *existing)

	return nil
}

func (r *Repository) Get(id string) *domain.PasswordCard {
	for i := range *r.PasswordCards {
		card := (*r.PasswordCards)[i]
		if card.Id.String() == id {
			card.Password = r.crypto.DecryptAES(card.Password)
			return &card
		}
	}
	return nil
}

func (r *Repository) Update(id string, dto *domain.PasswordCardUpdate) error {
	existing := r.Get(id)

	if existing == nil {
		return errors.New("NotFound: PasswordCard")
	}

	utils.Update(r.PasswordCards, id, dto, r.crypto)

	return nil
}
