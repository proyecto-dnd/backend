package characterdata

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	"github.com/proyecto-dnd/backend/internal/skill"
)

// To do: Optimize query quantity, Add goroutines to db calls

type service struct {
	characterRepo RepositoryCharacterData
	itemService   itemxcharacterdata.ServiceItemXCharacterData
	skillService  skill.ServiceSkill
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
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements ServiceCharacterData.
func (s *service) GetAll() ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetAll()
	if err != nil {
		return []dto.FullCharacterData{}, err
	}
	var fullCharacters []dto.FullCharacterData
	for _, v := range characters {
		skills, err := s.skillService.GetByCharacterId(int(v.Character_Id))
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		items, err := s.itemService.GetByCharacterDataId(v.Character_Id)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, CharacterDataToFullCharacterData(v, items, skills))

	}
	return fullCharacters, nil
}

// GetByCampaignId implements ServiceCharacterData.
func (s *service) GetByCampaignId(campaignid int) ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetByCampaignId(campaignid)
	if err != nil {
		return []dto.FullCharacterData{}, err
	}
	var fullCharacters []dto.FullCharacterData
	for _, v := range characters {
		skills, err := s.skillService.GetByCharacterId(int(v.Character_Id))
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		items, err := s.itemService.GetByCharacterDataId(v.Character_Id)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, CharacterDataToFullCharacterData(v, items, skills))

	}
	return fullCharacters, nil
}

// GetById implements ServiceCharacterData.
func (s *service) GetById(id int) (dto.FullCharacterData, error) {
	character, err := s.characterRepo.GetById(id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}

	skills, err := s.skillService.GetByCharacterId(int(character.Character_Id))
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	items, err := s.itemService.GetByCharacterDataId(character.Character_Id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return CharacterDataToFullCharacterData(character, items, skills), nil
}

// GetByUserId implements ServiceCharacterData.
func (s *service) GetByUserId(userid string) ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetByUserId(userid)
	if err != nil {
		return []dto.FullCharacterData{}, err
	}

	var fullCharacters []dto.FullCharacterData
	for _, v := range characters {
		skills, err := s.skillService.GetByCharacterId(int(v.Character_Id))
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		items, err := s.itemService.GetByCharacterDataId(v.Character_Id)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, CharacterDataToFullCharacterData(v, items, skills))

	}
	return fullCharacters, nil
}

// GetByUserIdAndCampaignId implements ServiceCharacterData.
func (s *service) GetByUserIdAndCampaignId(userid string, campaignid int) ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetByUserIdAndCampaignId(userid, campaignid)
	if err != nil {
		return []dto.FullCharacterData{}, err
	}

	var fullCharacters []dto.FullCharacterData
	for _, v := range characters {
		skills, err := s.skillService.GetByCharacterId(int(v.Character_Id))
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		items, err := s.itemService.GetByCharacterDataId(v.Character_Id)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, CharacterDataToFullCharacterData(v, items, skills))

	}

	return fullCharacters, nil
}

// Update implements ServiceCharacterData.
func (s *service) Update(character domain.CharacterData) (dto.FullCharacterData, error) {
	updatedCharacter, err := s.characterRepo.Update(character)
	if err != nil {
		return dto.FullCharacterData{}, err
	}

	skills, err := s.skillService.GetByCharacterId(int(character.Character_Id))
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	items, err := s.itemService.GetByCharacterDataId(character.Character_Id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return CharacterDataToFullCharacterData(updatedCharacter, items, skills), nil
}

func CharacterDataToFullCharacterData(character domain.CharacterData, items []domain.ItemXCharacterData, skills []domain.Skill) dto.FullCharacterData {
	return dto.FullCharacterData{
		Character_Id: character.Character_Id,
		User_Id:      character.User_Id,
		Campaign_Id: character.Campaign_Id,
		Race: character.Race,
		Class: character.Race,
		Background: character.Background,
		Name: character.Name,
		Story: character.Story,
		Alignment: character.Alignment,
		Age: character.Age,
		Hair: character.Hair,
		Eyes: character.Eyes,
		Skin: character.Skin,
		Height: character.Height,
		Weight: character.Weight,
		ImgUrl: character.ImgUrl,
		Str: character.Str,
		Dex: character.Dex,
		Int: character.Int,
		Con: character.Con,
		Wiz: character.Wiz,
		Cha: character.Cha,
		Hitpoints: character.Hitpoints,
		HitDice: character.HitDice,
		Speed: character.Speed,
		Armor_Class: character.Armor_Class,
		Level: character.Level,
		Exp: character.Exp,
		Items: items,
		Skills: skills,
	}
}

func FullCharacterDataToCharacterData(character dto.FullCharacterData) domain.CharacterData {
	return domain.CharacterData{
		Character_Id: character.Character_Id,
		User_Id:      character.User_Id,
		Campaign_Id: character.Campaign_Id,
		Race: character.Race,
		Class: character.Race,
		Background: character.Background,
		Name: character.Name,
		Story: character.Story,
		Alignment: character.Alignment,
		Age: character.Age,
		Hair: character.Hair,
		Eyes: character.Eyes,
		Skin: character.Skin,
		Height: character.Height,
		Weight: character.Weight,
		ImgUrl: character.ImgUrl,
		Str: character.Str,
		Dex: character.Dex,
		Int: character.Int,
		Con: character.Con,
		Wiz: character.Wiz,
		Cha: character.Cha,
		Hitpoints: character.Hitpoints,
		HitDice: character.HitDice,
		Speed: character.Speed,
		Armor_Class: character.Armor_Class,
		Level: character.Level,
		Exp: character.Exp,
	}
}