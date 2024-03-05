package friendship

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type service struct {
	repository FriendshipRepository
}

func NewFriendshipService(repository FriendshipRepository) FriendshipService {
	return &service{repository: repository}
}

func (s *service) Create(friendship domain.Friendship) (domain.Friendship, error) {
	createdFriendship, err := s.repository.Create(friendship)
	if err != nil {
		return domain.Friendship{}, err
	}

	return createdFriendship, nil
}

func (s *service) Delete(friendship domain.Friendship) error {
	return s.repository.Delete(friendship)
}
