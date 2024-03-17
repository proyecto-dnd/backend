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
	// mutuals is the user id and the friend's name

	// usersList is a list of all users obtained from firebase
	// the users format is: {id, username, email}
	usersList, err := r.userRepository.GetAll()
	if err != nil {
		return []domain.UserResponse{}, err
	}

	// userListByName is a list of users that start with the friend's name
	var userListByName []domain.UserResponse
	for _, user := range usersList {
		if strings.HasPrefix(user.Username, mutuals.User2Name) {
			userListByName = append(userListByName, user)
		}
	}

	user1Friends, err := r.GetFriends(mutuals.User1Id)
	if err != nil {
		return []domain.UserResponse{}, err
	}
	// I need to filter the user1's friends by the user2's id

	var tempFriendList []domain.UserResponse
	for _, friend := range user1Friends {
		for _, user := range userListByName {
			if friend.User2Id == user.Id {
				tempFriendList = append(tempFriendList, user)
			}
		}
	}

	// have to bring the friends from sql to cmpare the id of user2 for all users1
	// result is a struct with a list of users and a list of not found users from firebase
	var getUsersList auth.GetUsersResult
	for _, user := range userListByName {
		getUserResult, err := r.authClient.GetUsers(ctx, []auth.UserIdentifier{auth.UIDIdentifier{UID: user.Id}})
		if err != nil {
			log.Fatalf("error retriving user: %v\n", err)
		}
		// if getUserResult.Users[0].UID ==

		getUsersList.Users = append(getUsersList.Users, getUserResult.Users...)
		getUsersList.NotFound = append(getUsersList.NotFound, getUserResult.NotFound...)
	}
	// I need to filter the user1's friends by the user2's id

	// fmt.Println(&getUsersList.Users)
	for _, u := range getUsersList.Users {
		log.Printf("%v", u.DisplayName)
		log.Printf("%v", u.UID)
	}

	return tempFriendList, nil
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
