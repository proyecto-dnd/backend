package user

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type RepositoryUsers interface {
	Create(user domain.User) (domain.User, error)
	GetAll() ([]domain.UserResponse, error)
	GetById(id string) (domain.User, error)
	Update(user domain.UserUpdate, id string) (domain.UserUpdate, error)
	Delete(id string) error
	Patch(user domain.User, id string) (domain.User, error)
	Login(userInfo domain.UserLoginInfo) (string, error)
	GetJwtInfo(cookieToken string) (domain.UserTokenClaims, error)
	TransferDataToSql(users []domain.User) (string, error)
}

type ServiceUsers interface {
	Create(user domain.User) (domain.UserResponse, error)
	GetAll() ([]domain.UserResponse, error)
	GetById(id string) (domain.UserResponse, error)
	Update(user domain.UserUpdate, id string) (domain.UserResponse, error)
	Patch(user domain.User, id string) (domain.User, error)
	Delete(id string) error
	Login(userInfo domain.UserLoginInfo) (string, error)
	GetJwtInfo(cookieToken string) (domain.UserTokenClaims, error)
	TransferDataToSql(users []domain.User) (string, error)
}
