package characterdata

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type RepositoryCharacterData interface {
	Create(character domain.CharacterData) (domain.CharacterData, error)
	GetAll()([]domain.CharacterData, error)
	GetById(id int)(domain.CharacterData, error)
	GetByUserId(userid string)([]domain.CharacterData, error)
	GetByUserIdAndCampaignId(userid string, campaignid int)([]domain.CharacterData, error)
	GetByCampaignId(campaignid int)([]domain.CharacterData, error)
	Update(character domain.CharacterData) (domain.CharacterData, error)
	Delete(id int)error
}

type ServiceCharacterData interface {
	Create(character domain.CharacterData) (dto.FullCharacterData, error)
	GetAll()([]dto.FullCharacterData, error)
	GetById(id int)(dto.FullCharacterData, error)
	GetByUserId(userid string)([]dto.FullCharacterData, error)
	GetByUserIdAndCampaignId(userid string, campaignid int)([]dto.FullCharacterData, error)
	GetByCampaignId(campaignid int)([]dto.FullCharacterData, error)
	Update(character domain.CharacterData) (dto.FullCharacterData, error)
	Delete(id int)error
}