package friendship

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

type repositoryFriendship struct {
	db *sql.DB
}

var (
	ErrPrepareStatement = errors.New("error preparing statement")
)

func NewFriendshipRepository(db *sql.DB) FriendshipRepository {
	return &repositoryFriendship{db: db}
}

func (r *repositoryFriendship) Create(friendship domain.Friendship) (domain.Friendship, error) {
	statement, err := r.db.Prepare(QueryCreate)
	if err != nil {
		return domain.Friendship{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		friendship.User1Id,
		friendship.User2Id,
	)
	if err != nil {
		return domain.Friendship{}, err
	}

	return friendship, nil
}

func (r *repositoryFriendship) Delete(friendship domain.Friendship) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		friendship.User1Id,
		friendship.User2Id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositoryFriendship) IsFriends(userId1 string, userId2 string) (bool, error) {
	statement, err := r.db.Prepare(QueryCheckFriendship)
	if err != nil {
		return false, ErrPrepareStatement
	}
	defer statement.Close()

	var count int
	err = statement.QueryRow(userId1, userId2, userId2, userId1).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *repositoryFriendship) GetFriends(userId string) ([]domain.Friendship, error) {	
	statement, err := r.db.Prepare(QueryGetFriends)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friendships []domain.Friendship
	for rows.Next() {
		var friend domain.Friendship
		err = rows.Scan(&friend.User1Id, &friend.User2Id)
		if err != nil {
			return nil, err
		}
		if friend.User1Id == userId {
			friendships = append(friendships, friend)
		} else {
			friend.User1Id, friend.User2Id = friend.User2Id, friend.User1Id
			friendships = append(friendships, friend)
		}
	}

	return friendships, nil
}
