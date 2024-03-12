package characterdata

import (
	"github.com/proyecto-dnd/backend/internal/armorXCharacterData"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/feature"
	"github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	"github.com/proyecto-dnd/backend/internal/proficiency"
	"github.com/proyecto-dnd/backend/internal/skill"
	"github.com/proyecto-dnd/backend/internal/spell"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
)

// To do: Optimize query quantity, Add goroutines to db calls

type service struct {
	characterRepo      RepositoryCharacterData
	itemService        itemxcharacterdata.ServiceItemXCharacterData
	weaponService      weaponxcharacterdata.ServiceWeaponXCharacterData
	armorService       armorXCharacterData.ServiceArmorXCharacterData
	skillService       skill.ServiceSkill
	featureService     feature.FeatureService
	spellService       spell.ServiceSpell
	proficiencyService proficiency.ProficiencyService
}

func NewServiceCharacterData(characterRepo RepositoryCharacterData, itemService itemxcharacterdata.ServiceItemXCharacterData, weaponService weaponxcharacterdata.ServiceWeaponXCharacterData, armorService armorXCharacterData.ServiceArmorXCharacterData, skillService skill.ServiceSkill, featureService feature.FeatureService, spellService spell.ServiceSpell, proficiencyService proficiency.ProficiencyService) ServiceCharacterData {
	return &service{characterRepo: characterRepo, itemService: itemService, weaponService: weaponService, armorService: armorService, skillService: skillService, featureService: featureService, spellService: spellService, proficiencyService: proficiencyService}
}

// Create implements ServiceCharacterData.
func (s *service) Create(character domain.CharacterData) (dto.FullCharacterData, error) {
	newCharacter, err := s.characterRepo.Create(character)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return *characterDataToFullCharacterData(newCharacter, []domain.ItemXCharacterData{}, []domain.WeaponXCharacterData{}, []domain.ArmorXCharacterData{}, []domain.Skill{}, []domain.Feature{}, []domain.Spell{}, []domain.Proficiency{}), nil
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
		fullCharacter, err := s.fetchAndConvertToFullCharacterData(&v)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, *fullCharacter)

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
		fullCharacter, err := s.fetchAndConvertToFullCharacterData(&v)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, *fullCharacter)

	}
	return fullCharacters, nil
}

// GetById implements ServiceCharacterData.
func (s *service) GetById(id int) (dto.FullCharacterData, error) {
	character, err := s.characterRepo.GetById(id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}

	fullCharacter, err := s.fetchAndConvertToFullCharacterData(&character)
		if err != nil {
			return dto.FullCharacterData{}, err
		}
	return *fullCharacter, nil
}

// GetByUserId implements ServiceCharacterData.
func (s *service) GetByUserId(userid string) ([]dto.FullCharacterData, error) {
	characters, err := s.characterRepo.GetByUserId(userid)
	if err != nil {
		return []dto.FullCharacterData{}, err
	}

	var fullCharacters []dto.FullCharacterData
	for _, v := range characters {
		fullCharacter, err := s.fetchAndConvertToFullCharacterData(&v)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, *fullCharacter)

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
		fullCharacter, err := s.fetchAndConvertToFullCharacterData(&v)
		if err != nil {
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, *fullCharacter)

	}

	return fullCharacters, nil
}

// Update implements ServiceCharacterData.
func (s *service) Update(character domain.CharacterData) (dto.FullCharacterData, error) {
	updatedCharacter, err := s.characterRepo.Update(character)
	if err != nil {
		return dto.FullCharacterData{}, err
	}

	updatedFullCharacter, err := s.fetchAndConvertToFullCharacterData(&updatedCharacter)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return *updatedFullCharacter, nil
}

func characterDataToFullCharacterData(character domain.CharacterData, items []domain.ItemXCharacterData, weapons []domain.WeaponXCharacterData, armor []domain.ArmorXCharacterData, skills []domain.Skill, features []domain.Feature, spells []domain.Spell, proficiencies []domain.Proficiency) *dto.FullCharacterData {
	return &dto.FullCharacterData{
		Character_Id:  character.Character_Id,
		User_Id:       character.User_Id,
		Campaign_Id:   character.Campaign_Id,
		Race:          character.Race,
		Class:         character.Class,
		Background:    character.Background,
		Name:          character.Name,
		Story:         character.Story,
		Alignment:     character.Alignment,
		Age:           character.Age,
		Hair:          character.Hair,
		Eyes:          character.Eyes,
		Skin:          character.Skin,
		Height:        character.Height,
		Weight:        character.Weight,
		ImgUrl:        character.ImgUrl,
		Str:           character.Str,
		Dex:           character.Dex,
		Int:           character.Int,
		Con:           character.Con,
		Wiz:           character.Wiz,
		Cha:           character.Cha,
		Hitpoints:     character.Hitpoints,
		HitDice:       character.HitDice,
		Speed:         character.Speed,
		Armor_Class:   character.Armor_Class,
		Level:         character.Level,
		Exp:           character.Exp,
		Items:         items,
		Weapons:       weapons,
		Armor:         armor,
		Skills:        skills,
		Features:      features,
		Spells:        spells,
		Proficiencies: proficiencies,
	}
}

func fullCharacterDataToCharacterData(character dto.FullCharacterData) *domain.CharacterData {
	return &domain.CharacterData{
		Character_Id: character.Character_Id,
		User_Id:      character.User_Id,
		Campaign_Id:  character.Campaign_Id,
		Race:         character.Race,
		Class:        character.Class,
		Background:   character.Background,
		Name:         character.Name,
		Story:        character.Story,
		Alignment:    character.Alignment,
		Age:          character.Age,
		Hair:         character.Hair,
		Eyes:         character.Eyes,
		Skin:         character.Skin,
		Height:       character.Height,
		Weight:       character.Weight,
		ImgUrl:       character.ImgUrl,
		Str:          character.Str,
		Dex:          character.Dex,
		Int:          character.Int,
		Con:          character.Con,
		Wiz:          character.Wiz,
		Cha:          character.Cha,
		Hitpoints:    character.Hitpoints,
		HitDice:      character.HitDice,
		Speed:        character.Speed,
		Armor_Class:  character.Armor_Class,
		Level:        character.Level,
		Exp:          character.Exp,
	}
}

// TO DO: Implement this function using Goroutines
func (s *service) fetchAndConvertToFullCharacterData(character *domain.CharacterData) (*dto.FullCharacterData, error) {
	items, err := s.itemService.GetByCharacterDataId(character.Character_Id)
	if err != nil {
		return &dto.FullCharacterData{}, err
	}
	weapons, err := s.weaponService.GetByCharacterDataId(character.Character_Id)
	if err != nil {
		return &dto.FullCharacterData{}, err
	}
	armor, err := s.armorService.GetByCharacterDataIdArmor(character.Character_Id)
	if err != nil {
		return &dto.FullCharacterData{}, err
	}
	skills, err := s.skillService.GetByCharacterId(character.Character_Id)
	if err != nil {
		return &dto.FullCharacterData{}, err
	}
	features, err := s.featureService.GetByCharacterDataId(character.Character_Id)
	if err != nil {
		return &dto.FullCharacterData{}, err
	}
	spells, err := s.spellService.GetByCharacterDataId(character.Character_Id)
	if err != nil {
		return &dto.FullCharacterData{}, err
	}
	proficiencies, err := s.proficiencyService.GetByCharacterDataId(character.Character_Id)
	if err != nil {
		return &dto.FullCharacterData{}, err
	}
	return characterDataToFullCharacterData(*character, items, weapons, armor, skills, features, spells, proficiencies), nil
}
