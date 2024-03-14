package user

import (
	"fmt"
	"log"
	"time"

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
		log.Printf("Error initializing Firebase Auth client: %v", err)
	}
	return &repositoryFirebase{app: app, authClient: authClient}
}

func (r *repositoryFirebase) Create(user domain.User) (domain.User, error) {

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(true).
		Password(user.Password).
		DisplayName(user.Username).
		Disabled(false)

	newUser, err := r.authClient.CreateUser(ctx, params)
	if err != nil {
		log.Printf("Error creating user: %v", err)
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
			log.Printf("paging error %v\n", err)
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
func (r *repositoryFirebase) GetByName(name string) ([]domain.User, error) {
	var user domain.User
	var users []domain.User
	pager := iterator.NewPager(r.authClient.Users(ctx, ""), 50, "")
	for {
		var authUsers []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&authUsers)
		if err != nil {
			log.Printf("paging error %v\n", err)
		}
		for _, u := range authUsers {
			if u.DisplayName == name {
				user.Username = u.DisplayName
				user.Id = u.UID
				users = append(users, user)
			}
		}
		if nextPageToken == "" {
			break
		}
	}
	return users, nil

	/*
		function debounce(func, delay) {
		let timeoutId;
		return function() {
			const context = this;
			const args = arguments;
			clearTimeout(timeoutId);
			timeoutId = setTimeout(() => {
			func.apply(context, args);
			}, delay);
		};
		}

		// Usage example
		const debouncedFunction = debounce(function() {
		console.log("Function debounced!");
		}, 300); // Debounce delay of 300 milliseconds

		// Attach this debounced function to your input event listener
		input.addEventListener("input", debouncedFunction);
	*/
}
func (r *repositoryFirebase) GetById(id string) (domain.User, error) {

	u, err := r.authClient.GetUser(ctx, id)
	if err != nil {
		//TODO RETURN ERROR
		log.Printf("error getting user %s: %v\n", id, err)
	}

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
	_, err := r.authClient.UpdateUser(ctx, id, params)
	if err != nil {
		//TODO RETURN ERROR
		log.Printf("error updating user: %v\n", err)
	}
	// log.Printf("Successfully updated user: %v\n", u)

	return user, nil
}
func (r *repositoryFirebase) Delete(id string) error {

	err := r.authClient.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("error deleting user: %v\n", err)
	}

	log.Printf("Successfully deleted user: %s\n", id)

	return nil
}

func (r *repositoryFirebase) Patch(user domain.User, id string) (domain.User, error) {
	return user, nil
}

func (r *repositoryFirebase) Login(userInfo domain.UserLoginInfo) (string, error) {

	expiresIn := time.Hour * 24 * 2

	cookie, err := r.authClient.SessionCookie(ctx, userInfo.IdToken, expiresIn)
	if err != nil {
		fmt.Printf("error creating session cookie: %v\n", err)
		return "", err
	}

	// ctx.SetCookie("Session", cookie, 3600, "/", "localhost", false, false)
	// log.Println("LISTO LA PETICION")
	return cookie, nil
}
