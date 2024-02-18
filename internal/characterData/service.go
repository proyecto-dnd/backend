package characterdata

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/item"
	itemxcharacterdata "github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	"github.com/proyecto-dnd/backend/internal/skill"
)

type service struct {
	characterRepo RepositoryCharacterData
	itemService itemxcharacterdata.ServiceItemXCharacterData
	skillService skill.ServiceSkill
}

func NewServiceCharacterData(characterRepo RepositoryCharacterData, itemService itemxcharacterdata.ServiceItemXCharacterData, skillService skill.ServiceSkill) ServiceCharacterData {
	return &service{characterRepo: characterRepo, itemService: itemService, skillService: skillService}
}

// Create implements ServiceCharacterData.
func (s *service) Create(character domain.CharacterData) (dto.FullCharacterData, error) {
	newCharacter, err := s.characterRepo.Create(character)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return CharacterDataToFullCharacterData(newCharacter, []domain.ItemXCharacterData{}, []domain.Skill{}), nil
}

// Delete implements ServiceCharacterData.
func (s *service) Delete(id int) error {
	err := s.characterRepo.Delete(id)
    if err!= nil {
        return err
    }
    return nil
}

// GetAll implements ServiceCharacterData.
func (s *service) GetAll() ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetAll()
    if err!= nil {
        return []dto.FullCharacterData{}, err
    }
	var fullCharacters []dto.FullCharacterData
	for _, v := range characters {
		skills, err := s.skillService.GetByCharacterId(int(v.CharacterId))
		if err!= nil {
            return []dto.FullCharacterData{}, err
        }
		items, err := s.itemService.GetByCharacterDataId(v.CharacterId)
		if err!= nil {
            return []dto.FullCharacterData{}, err
        }
		fullCharacters = append(fullCharacters, CharacterDataToFullCharacterData(v, items,skills))
		
		
	}
    return fullCharacters, nil
}

// GetByCampaignId implements ServiceCharacterData.
func (s *service) GetByCampaignId(campaignid int) ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetByCampaignId(campaignid)
	if err!= nil {
        return []dto.FullCharacterData{}, err
    }
	return characters, nil
}

// GetById implements ServiceCharacterData.
func (s *service) GetById(id int) (dto.FullCharacterData, error) {
	character, err := s.characterRepo.GetById(id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return character, nil
}

// GetByUserId implements ServiceCharacterData.
func (s *service) GetByUserId(userid string) ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetByUserId(userid)
	if err!= nil {
        return []dto.FullCharacterData{}, err
    }
	return characters, nil
}

// GetByUserIdAndCampaignId implements ServiceCharacterData.
func (s *service) GetByUserIdAndCampaignId(userid string, campaignid int) ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetByUserIdAndCampaignId(userid, campaignid)
    if err!= nil {
        return []dto.FullCharacterData{}, err
    }
    return characters, nil
}

// Update implements ServiceCharacterData.
func (s *service) Update(character domain.CharacterData) (dto.FullCharacterData, error) {
	updatedCharacter, err := s.characterRepo.Update(character)
    if err!= nil {
        return dto.FullCharacterData{}, err
    }
    return CharacterDataToFullCharacterData(updatedCharacter), nil
}

func CharacterDataToFullCharacterData(character domain.CharacterData,items []domain.ItemXCharacterData, skills []domain.Skill) dto.FullCharacterData {
	return dto.FullCharacterData{
		CharacterId: character.CharacterId,
		UserId: character.UserId,
        Name: character.Name,
        Class: character.Class,
        Race: character.Race,
        Background: character.Background,
		Hitpoints: character.Hitpoints,
        Speed: character.Speed,
        ArmorClass: character.ArmorClass,
        Level: character.Level,
        Exp: character.Exp,
		CampaignId: character.CampaignId,
        Str: character.Str,
        Dex: character.Dex,
        Int: character.Int,
        Wiz: character.Wiz,
		Con: character.Con,
		Cha: character.Cha,
        Items: items,
        Skills: skills,
	}
}

func FullCharacterDataToCharacterData(character dto.FullCharacterData) domain.CharacterData {
	return domain.CharacterData{
		CharacterId: character.CharacterId,
		UserId: character.UserId,
        Name: character.Name,
        Class: character.Class,
        Race: character.Race,
        Background: character.Background,
		Hitpoints: character.Hitpoints,
        Speed: character.Speed,
        ArmorClass: character.ArmorClass,
        Level: character.Level,
        Exp: character.Exp,
		CampaignId: character.CampaignId,
        Str: character.Str,
        Dex: character.Dex,
        Int: character.Int,
        Wiz: character.Wiz,
		Con: character.Con,
		Cha: character.Cha,
	}
}