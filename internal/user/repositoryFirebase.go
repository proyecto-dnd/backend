package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/proyecto-dnd/backend/internal/domain"
	"google.golang.org/api/iterator"
)

var (
	ctx      = &gin.Context{}
	ErrEmpty = errors.New("empty list")
)

type repositoryFirebase struct {
	app        *firebase.App
	authClient *auth.Client
	db         *sql.DB
}

func NewUserFirebaseRepository(app *firebase.App, db *sql.DB) RepositoryUsers {
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Printf("Error initializing Firebase Auth client: %v", err)
	}
	return &repositoryFirebase{app: app, authClient: authClient, db: db}
}

func (r *repositoryFirebase) Create(user domain.User) (domain.UserResponse, error) {
	//sql backup
	statement, err := r.db.Prepare(QueryInsertUser)
	if err != nil {
		return domain.UserResponse{}, err
	}
	defer statement.Close()

	fmt.Println(user)
	fmt.Println(user.DisplayName)
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password).
		DisplayName(user.Username).
		Disabled(false).
		PhotoURL("https://proyecto-dnd.vercel.app/user.png")

	//firebase create
	newUser, err := r.authClient.CreateUser(ctx, params)
	if err != nil {
		log.Printf("Error creating user: %v", err)
	}

	claims := map[string]interface{}{"displayName": user.DisplayName, "subExpiration": time.Unix(0, 0).String()}

	err = r.authClient.SetCustomUserClaims(ctx, newUser.UID, claims)
	if err != nil {
		fmt.Println("Error setting custom user claims.")
		return domain.UserResponse{}, err
	}

	//sql create
	_, err = statement.Exec(newUser.UID, user.Username, user.Email, user.Password, user.DisplayName, "https://proyecto-dnd.vercel.app/user.png")
	if err != nil {
		fmt.Println("Error setting custom user claims.")
		return domain.UserResponse{}, err
	}

	var userTemp domain.UserResponse
	userTemp.Username = newUser.DisplayName
	userTemp.Email = newUser.Email
	userTemp.Id = newUser.UID
	userTemp.DisplayName = user.DisplayName
	// newUser.PasswordHash = user.Password

	return userTemp, nil
}
func (r *repositoryFirebase) GetAll() ([]domain.UserResponse, error) {

	// var user domain.UserResponse
	var users []domain.UserResponse
	// pager := iterator.NewPager(r.authClient.Users(ctx, ""), 100, "")

	// for {
	// 	var authUsers []*auth.ExportedUserRecord
	// 	nextPageToken, err := pager.NextPage(&authUsers)
	// 	if err != nil {
	// 		log.Printf("paging error %v\n", err)
	// 	}
	// 	for _, u := range authUsers {
	// 		user.Username = u.DisplayName
	// 		user.Email = u.Email
	// 		// user.Password = u.PasswordHash
	// 		user.Id = u.UID
	// 		users = append(users, user)
	// 	}
	// 	if nextPageToken == "" {
	// 		break
	// 	}
	// }

	rows, err := r.db.Query(QueryGetAllUsers)
	if err != nil {
		return []domain.UserResponse{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var user domain.UserResponse
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.DisplayName, &user.Image); err != nil {
			return []domain.UserResponse{}, err
		}

		users = append(users, user)
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
}

func (r *repositoryFirebase) GetById(id string) (domain.UserResponse, error) {

	_, err := r.authClient.GetUser(ctx, id)
	if err != nil {
		//TODO RETURN ERROR
		log.Printf("error getting user %s: %v\n", id, err)
	}

	row, err := r.db.Query(QueryGetUserById, id)
	if err != nil {
		return domain.UserResponse{}, err
	}
	defer row.Close()

	var user domain.UserResponse
	for row.Next() {
		if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.DisplayName, &user.Image); err != nil {
			return domain.UserResponse{}, err
		}
	}
	// fmt.Println(user)
	// var user domain.User
	// user.Username = u.DisplayName
	// user.Email = u.Email
	// user.Id = u.UID

	return user, nil
}
func (r *repositoryFirebase) Update(user domain.UserUpdate, id string) (domain.UserUpdate, error) {
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
	result, err := r.db.Exec(QueryUpdateUser,
		user.Username,
		user.Email,
		user.Password,
		user.Image,
		user.DisplayName,
		id,
	)
	if err != nil {
		return domain.UserUpdate{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.UserUpdate{}, err
	}
	user.Id = id

	return user, nil
}
func (r *repositoryFirebase) Delete(id string) error {
	// userList, err := r.GetAll()
	// if err != nil {
	// 	return err
	// }
	// //extract uid from userList

	// var idList []string
	// for _, u := range userList {
	// 	id = u.Id
	// 	idList = append(idList, id)
	// }
	// r.authClient.DeleteUsers(ctx, idList)

	// fmt.Println("https://media1.tenor.com/m/LBWyQg647MoAAAAC/execute-order66-order66.gif")

	err := r.authClient.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("error deleting user: %v\n", err)
	}

	result, err := r.db.Exec(QueryDeleteUser, id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("Successfully deleted user: %s\n", id)

	return nil
}

func (r *repositoryFirebase) Patch(user domain.UserUpdate, id string) (domain.UserResponse, error) {
	var fieldsToUpdate []string
	var args []interface{}

	if user.Username != "" {
		fieldsToUpdate = append(fieldsToUpdate, "display_name = ?")
		args = append(args, user.Username)
		r.authClient.UpdateUser(ctx, id, (&auth.UserToUpdate{}).DisplayName(user.Username))
	}
	if user.Email != "" {
		fieldsToUpdate = append(fieldsToUpdate, "email = ?")
		args = append(args, user.Email)
		r.authClient.UpdateUser(ctx, id, (&auth.UserToUpdate{}).Email(user.Email))
	}
	if user.Password != "" {
		fieldsToUpdate = append(fieldsToUpdate, "password = ?")
		args = append(args, user.Password)
		r.authClient.UpdateUser(ctx, id, (&auth.UserToUpdate{}).Password(user.Password))
	}
	if user.Image != nil && *user.Image != "" {
		fieldsToUpdate = append(fieldsToUpdate, "image = ?")
		args = append(args, user.Image)
	}
	if user.DisplayName != "" {
		fieldsToUpdate = append(fieldsToUpdate, "display_name = ?")
		args = append(args, user.DisplayName)
	}

	if len(fieldsToUpdate) == 0 {
		return domain.UserResponse{}, ErrEmpty
	}

	queryString := "UPDATE user SET " + strings.Join(fieldsToUpdate, ", ") + " WHERE user_id = ?"
	args = append(args, id)

	result, err := r.db.Exec(queryString, args...)
	if err != nil {
		return domain.UserResponse{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.UserResponse{}, err
	}

	return domain.UserResponse{
		Id:          id,
		Username:    user.Username,
		Email:       user.Email,
		Image:       user.Image,
		DisplayName: user.DisplayName,
	}, nil
}

func (r *repositoryFirebase) Login(userInfo domain.UserLoginInfo) (string, error) {

	expiresIn := time.Hour * 24 * 2

	cookie, err := r.authClient.SessionCookie(ctx, userInfo.IdToken, expiresIn)
	if err != nil {
		fmt.Printf("error creating session cookie: %v\n", err)
		return "error creating session cookie", err
	}

	return cookie, nil
}

func (r *repositoryFirebase) GetJwtInfo(cookieToken string) (domain.UserTokenClaims, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(cookieToken, jwt.MapClaims{})
	if err != nil {
		return domain.UserTokenClaims{}, err
	}

	var tokenClaims domain.UserTokenClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		uid := claims["user_id"].(string)
		username := claims["name"].(string)
		email := claims["email"].(string)
		displayName := claims["displayName"].(string)
		subExpirationDate := claims["subExpiration"].(string)

		tokenClaims.Id = uid
		tokenClaims.Username = username
		tokenClaims.Email = email
		tokenClaims.DisplayName = displayName
		tokenClaims.SubExpirationDate = subExpirationDate
	}

	return tokenClaims, nil
}

func (r *repositoryFirebase) TransferDataToSql(users []domain.User) (string, error) {

	insertString, err := r.BulkInsertString(users)

	if err != nil {
		return "", err
	}
	// fmt.Println(insertString)

	result, err := r.db.Exec(insertString)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowsAffected < 1 {
		return "", errors.New("no rows affected")
	}

	// fmt.Println(rowsAffected)

	return insertString, nil
}

func (r *repositoryFirebase) BulkInsertString(users []domain.User) (string, error) {
	var values strings.Builder

	// values.WriteString("(")

	for i, user := range users {

		values.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')", user.Id, user.Username, user.Email, user.Password, user.DisplayName))

		if i < len(users)-1 {
			values.WriteString(", ")
		}
	}

	// values.WriteString(")")

	insertSQL := fmt.Sprintf("INSERT INTO user (uid, name, email, password, display_name) VALUES %s;", values.String())
	return insertSQL, nil
}

func (r *repositoryFirebase) SubscribeToPremium(id string, date string) error {
	claims := map[string]interface{}{"subExpiration": date}

	err := r.authClient.SetCustomUserClaims(ctx, id, claims)
	if err != nil {
		fmt.Println("Error setting custom user claims: " + err.Error())
		return err
	}
	return nil
}
