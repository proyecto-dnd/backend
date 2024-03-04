package friendship

import "github.com/proyecto-dnd/backend/internal/domain"

type FriendshipRepository interface {
	Create(friendship domain.Friendship) (domain.Friendship, error)
	Delete(friendship domain.Friendship) error
}

type FriendshipService interface {
	Create(friendship domain.Friendship) (domain.Friendship, error)
	Delete(friendship domain.Friendship) error
}
