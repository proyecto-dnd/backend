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

func (s *service) SearchFollowers(friendship domain.Mutuals) ([]domain.UserResponse, error) {
	return s.repository.SearchFollowers(friendship)
}

func (s *service) GetAllFriends(userId string) ([]domain.Friendship, error) {
	return s.repository.GetAllFriends(userId)
}

func (s *service) GetBySimilarName(input string) ([]domain.UserResponse, error) {
	users, err := s.repository.GetBySimilarName(input)
	if err != nil {
		return []domain.UserResponse{}, err
	}

	if len(users) == 0 {
		return []domain.UserResponse{}, nil
	}

	maxLength := 5
	if len(users) < maxLength {
		maxLength = len(users)
	}
	fiveResultsforUsers := users[:maxLength]

	return fiveResultsforUsers, nil
}
