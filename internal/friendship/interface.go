package friendship

import "github.com/proyecto-dnd/backend/internal/domain"

type FriendshipRepository interface {
	Create(friendship domain.Friendship) (domain.Friendship, error)
	Delete(friendship domain.Friendship) error
	SearchFollowers(friendship domain.Mutuals) ([]domain.UserResponse, error)
	GetAllFriends(userId string) ([]domain.Friendship, error)
}

type FriendshipService interface {
	Create(friendship domain.Friendship) (domain.Friendship, error)
	Delete(friendship domain.Friendship) error
	SearchFollowers(friendship domain.Mutuals) ([]domain.UserResponse, error)
	GetAllFriends(userId string) ([]domain.Friendship, error)
}
