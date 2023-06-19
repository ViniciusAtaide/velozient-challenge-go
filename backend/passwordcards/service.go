package passwordcards

import (
	"github.com/viniciusataide/velozient-challenge-go/domain"
	"github.com/viniciusataide/velozient-challenge-go/utils"
)

type Service struct{ Repository *Repository }

func ProvideService(repository *Repository) *Service {
	return &Service{
		repository,
	}
}

func (s *Service) List(pagination *utils.Pagination, filter string) ([]domain.PasswordCard, error) {
	return s.Repository.List(pagination, filter), nil
}

func (s *Service) Create(dto domain.PasswordCardCreate) (*domain.PasswordCard, error) {
	card, err := s.Repository.Create(dto)

	return card, err
}

func (s *Service) Delete(id string) error {
	err := s.Repository.Delete(id)

	return err
}

func (s *Service) Update(id string, pc *domain.PasswordCardUpdate) error {
	err := s.Repository.Update(id, pc)

	return err
}

func (s *Service) Get(id string) *domain.PasswordCard {
	pc := s.Repository.Get(id)

	return pc
}
