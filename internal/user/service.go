package user

import (
	"log"

	"github.com/proyecto-dnd/backend/internal/domain"
)

type service struct {
	repositoryFirebase RepositoryUsers
}

func NewServiceUser(repositoryFirebase RepositoryUsers) ServiceUsers {
	return &service{repositoryFirebase: repositoryFirebase}
}

func userToUserResponse(user domain.User) domain.UserResponse {
	var userResponse domain.UserResponse
	userResponse.Id = user.Id
	userResponse.Username = user.Username
	userResponse.Email = user.Email

	return userResponse
}

func (s *service) Create(user domain.User) (domain.UserResponse, error) {
	newUser, err := s.repositoryFirebase.Create(user)
	if err != nil {
		return domain.UserResponse{}, err
	}

	return userToUserResponse(newUser), nil
}

func (s *service) GetAll() ([]domain.UserResponse, error) {
	users, err := s.repositoryFirebase.GetAll()
	if err != nil {
		return []domain.UserResponse{}, err
	}

	var usersResponse []domain.UserResponse

	for _, u := range users {
		usersResponse = append(usersResponse, userToUserResponse(u))
	}

	return usersResponse, nil
}

func (s *service) GetById(id string) (domain.UserResponse, error) {
	user, err := s.repositoryFirebase.GetById(id)
	if err != nil {
		return domain.UserResponse{}, err
	}

	return userToUserResponse(user), nil
}

func (s *service) Update(user domain.User, id string) (domain.User, error) {
	updatedUser, err := s.repositoryFirebase.Update(user, id)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (s *service) Patch(user domain.User, id string) (domain.User, error) {
	patchedUser, err := s.repositoryFirebase.Patch(user, id)
	if err != nil {
		return domain.User{}, err
	}

	return patchedUser, nil
}

func (s *service) Delete(id string) error {
	err := s.repositoryFirebase.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) Login(userInfo domain.UserLoginInfo) (string, error) {

	cookie, err := s.repositoryFirebase.Login(userInfo)
	if err != nil {
		log.Println("ACAAAAAAAAAA")
		log.Printf(err.Error())
		return "", err
	}

	return cookie, nil
}
