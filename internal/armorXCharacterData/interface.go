package armorXCharacterData

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryArmorXCharacterData interface {
	CreateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
	DeleteArmorXCharacterData(id int64) error
	DeleteByCharacterDataIdArmor(id int64) error
	GetAllArmorXCharacterData() ([]domain.ArmorXCharacterData, error)
	GetByIdArmorXCharacterData(id int64) (domain.ArmorXCharacterData, error)
	GetByCharacterDataIdArmor(id int64) ([]domain.ArmorXCharacterData, error)
	UpdateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
}

type ServiceArmorXCharacterData interface {
	CreateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
	DeleteArmorXCharacterData(id int64) error
	DeleteByCharacterDataIdArmor(id int64) error
	GetAllArmorXCharacterData() ([]domain.ArmorXCharacterData, error)
	GetByIdArmorXCharacterData(id int64) (domain.ArmorXCharacterData, error)
	GetByCharacterDataIdArmor(id int64) ([]domain.ArmorXCharacterData, error)
	UpdateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error)
}
