package weaponxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryWeaponXCharacterData interface {
	Create(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	GetAll() ([]domain.WeaponXCharacterData, error)
	GetById(id int64) (domain.WeaponXCharacterData, error)
	GetByCharacterDataId(id int64) ([]domain.WeaponXCharacterData, error)
	Update(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	Delete(id int64) error
	DeleteByCharacterDataId(id int64) error
}

type ServiceWeaponXCharacterData interface {
	Create(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	GetAll() ([]domain.WeaponXCharacterData, error)
	GetById(id int64) (domain.WeaponXCharacterData, error)
	GetByCharacterDataId(id int64) ([]domain.WeaponXCharacterData, error)
	Update(WeaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error)
	Delete(id int64) error
	DeleteByCharacterDataId(id int64) error
}