package weapon

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type WeaponHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type WeaponService interface {
	CreateWeapon(dto.CreateWeaponDto) (domain.Weapon, error)
	GetAllWeapons() ([]domain.Weapon, error)
	GetWeaponById(id int) (domain.Weapon, error)
	UpdateWeapon(Weapon dto.CreateWeaponDto, id int) (domain.Weapon, error)
	DeleteWeapon(id int) error
}

type WeaponRepository interface {
	Create(Weapon domain.Weapon) (domain.Weapon, error)
	GetAllWeapons() ([]domain.Weapon, error)
	GetWeaponById(id int) (domain.Weapon, error)
	UpdateWeapon(Weapon domain.Weapon, id int) (domain.Weapon, error)
	DeleteWeapon(id int) error
}
