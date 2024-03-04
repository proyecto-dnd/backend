package itemxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryItemXCharacterData interface {
	Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	GetAll() ([]domain.ItemXCharacterData, error)
	GetById(id int64) (domain.ItemXCharacterData, error)
	GetByCharacterDataId(id int64) ([]domain.ItemXCharacterData, error)
	Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	Delete(id int64) error
	DeleteByCharacterDataId(id int64) error
}

type ServiceItemXCharacterData interface {
	Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	GetAll() ([]domain.ItemXCharacterData, error)
	GetById(id int64) (domain.ItemXCharacterData, error)
	GetByCharacterDataId(id int64) ([]domain.ItemXCharacterData, error)
	Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error)
	Delete(id int64) error
	DeleteByCharacterDataId(id int64) error
}