package user

import (
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"google.golang.org/api/iterator"
)

var (
	ctx = &gin.Context{}
)

type repositoryFirebase struct {
	app        *firebase.App
	authClient *auth.Client
}

func NewUserFirebaseRepository(app *firebase.App) RepositoryUsers {
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v", err)
	}
	return &repositoryFirebase{app: app, authClient: authClient}
}

func (r *repositoryFirebase) Create(user domain.User) (domain.User, error) {

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		Password(user.Password).
		DisplayName(user.Username).
		Disabled(false)

	newUser, err := r.authClient.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	var userTemp domain.User
	userTemp.Username = newUser.DisplayName
	userTemp.Email = newUser.Email
	userTemp.Id = newUser.UID

	return userTemp, nil
}
func (r *repositoryFirebase) GetAll() ([]domain.User, error) {

	var user domain.User
	var users []domain.User
	pager := iterator.NewPager(r.authClient.Users(ctx, ""), 100, "")
	for {
		var authUsers []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&authUsers)
		if err != nil {
			log.Fatalf("paging error %v\n", err)
		}
		for _, u := range authUsers {
			user.Username = u.DisplayName
			user.Email = u.Email
			user.Password = u.PasswordHash
			user.Id = u.UID
			users = append(users, user)
		}
		if nextPageToken == "" {
			break
		}
	}

	return users, nil
}
func (r *repositoryFirebase) GetById(id string) (domain.User, error) {

	u, err := r.authClient.GetUser(ctx, id)
	if err != nil {
		//TODO RETURN ERROR
		log.Fatalf("error getting user %s: %v\n", id, err)
	}

	log.Printf("Successfully fetched user data: %v\n", u)
	var user domain.User
	user.Username = u.DisplayName
	user.Email = u.Email
	user.Id = u.UID

	return user, nil
}
func (r *repositoryFirebase) Update(user domain.User, id string) (domain.User, error) {
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		Password(user.Password).
		DisplayName(user.Username)
	u, err := r.authClient.UpdateUser(ctx, id, params)
	if err != nil {
		//TODO RETURN ERROR
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)

	return user, nil
}
func (r *repositoryFirebase) Delete(id string) error {

	err := r.authClient.DeleteUser(ctx, id)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}

	log.Printf("Successfully deleted user: %s\n", id)

	return nil
}

func (r *repositoryFirebase) Patch(user domain.User, id string) (domain.User, error) {
	return user, nil
}
