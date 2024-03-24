package friendship

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/user"
)

type repositoryFriendship struct {
	db             *sql.DB
	userRepository user.RepositoryUsers
	firebaseApp    *firebase.App
	authClient     *auth.Client
}

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrDuplicateName    = errors.New("name already exists")
	ctx                 = &gin.Context{}
)

func NewFriendshipRepository(db *sql.DB, userRepository user.RepositoryUsers, app *firebase.App) FriendshipRepository {
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Printf("Error initializing Firebase Auth client: %v", err)
	}
	return &repositoryFriendship{db: db, userRepository: userRepository, authClient: authClient}
}

func (r *repositoryFriendship) SearchFollowers(mutuals domain.Mutuals) ([]domain.UserResponse, error) {

	usersListChan := make(chan []domain.UserResponse)
	user1FriendsChan := make(chan []domain.Friendship)
	tempFriendListChan := make(chan domain.UserResponse)

	go func() {
		usersList, err := r.userRepository.GetAll()
		if err != nil {
			usersListChan <- []domain.UserResponse{}
			return
		}
		usersListChan <- usersList
	}()

	go func() {
		user1Friends, err := r.GetAllFriends(mutuals.User1Id)
		if err != nil {
			user1FriendsChan <- []domain.Friendship{}
			return
		}
		user1FriendsChan <- user1Friends
	}()

	go func() {
		usersList := <-usersListChan
		userListByName := make([]domain.UserResponse, 0)
		for _, user := range usersList {
			if strings.HasPrefix(strings.ToLower(user.Username), strings.ToLower(mutuals.User2Name)) {
				userListByName = append(userListByName, user)
			}
		}
		user1Friends := <-user1FriendsChan
		for _, friend := range user1Friends {
			for _, user := range userListByName {
				if friend.User2Id == user.Id {
					tempFriendListChan <- user
				}
			}
		}
		close(tempFriendListChan)
	}()

	tempFriendList := []domain.UserResponse{}
	for user := range tempFriendListChan {
		tempFriendList = append(tempFriendList, user)
	}

	return tempFriendList, nil

	// var tempFriendList []domain.UserResponse
	// for _, friend := range user1Friends {
	// 	for _, user := range userListByName {
	// 		if friend.User2Id == user.Id {
	// 			tempFriendList = append(tempFriendList, user)
	// 		}
	// 	}
	// }

	// return tempFriendList, nil
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

func (r *repositoryFriendship) GetAllFriends(userId string) ([]domain.Friendship, error) {
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

func (r *repositoryFriendship) GetBySimilarName(input string) ([]domain.UserResponse, error) {
	// usersList, err := r.userRepository.GetAll()
	// if err != nil {
	// 	return []domain.UserResponse{}, err
	// }

	usersListByName := make([]domain.UserResponse, 0)
	// for _, user := range usersList {
	// 	if strings.HasPrefix(strings.ToLower(user.Username), strings.ToLower(input)) {
	// 		usersListByName = append(usersListByName, user)
	// 	}
	// }

	return usersListByName, nil
}
