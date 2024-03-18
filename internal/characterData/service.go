package characterdata

import (
	"fmt"
	"sync"
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
	newCharacterDto, err := s.GetById(newCharacter.Character_Id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return newCharacterDto, nil
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
			fmt.Println("algo")
			return []dto.FullCharacterData{}, err
		}
		fullCharacters = append(fullCharacters, fullCharacter)

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
		fullCharacters = append(fullCharacters, fullCharacter)

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
	return fullCharacter, nil
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
		fullCharacters = append(fullCharacters, fullCharacter)

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
		fullCharacters = append(fullCharacters, fullCharacter)

	}

	return fullCharacters, nil
}

// Update implements ServiceCharacterData.
func (s *service) Update(character domain.CharacterData) (dto.FullCharacterData, error) {
	updatedCharacter, err := s.characterRepo.Update(character)
	if err != nil {
		return dto.FullCharacterData{}, err
	}

	updatedFullCharacter, err := s.GetById(updatedCharacter.Character_Id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return updatedFullCharacter, nil
}

func characterDataToFullCharacterData(character domain.CharacterData, items []domain.ItemXCharacterData, weapons []domain.WeaponXCharacterData, armor []domain.ArmorXCharacterData, skills []domain.Skill, features []domain.Feature, spells []domain.Spell, proficiencies []domain.Proficiency) dto.FullCharacterData {
	return dto.FullCharacterData{
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

// TO DO: Finish Armor and skill implementation
func (s *service) fetchAndConvertToFullCharacterData(character *domain.CharacterData) (dto.FullCharacterData, error) {
	errChan := make(chan error, 5)
	itemChan := make(chan []domain.ItemXCharacterData, 1)
	weaponChan := make(chan []domain.WeaponXCharacterData, 1)
	featureChan := make(chan []domain.Feature, 1)
	spellChan := make(chan []domain.Spell, 1)
	proficiencyChan := make(chan []domain.Proficiency, 1)
	var wg sync.WaitGroup
	wg.Add(5) //TO DO change to 7 when armor and skills are working

	go func() {
		defer func() {
			close(itemChan)
			wg.Done()
		}()
		items, err := s.itemService.GetByCharacterDataId(character.Character_Id)

		errChan <- err
		itemChan <- items
	}()

	go func() {
		defer func() {
			close(weaponChan)
			wg.Done()
		}()
		weapons, err := s.weaponService.GetByCharacterDataId(character.Character_Id)
		errChan <- err
		weaponChan <- weapons
	}()

	go func() {
		defer func() {
			close(featureChan)
			wg.Done()
		}()
		featuresDto, err := s.featureService.GetAllFeaturesByCharacterId(character.Character_Id)
		errChan <- err
		featureChan <- featuresDto.Features
	}()

	go func() {
		defer func() {
			close(spellChan)
			wg.Done()
		}()
		spells, err := s.spellService.GetByCharacterDataId(character.Character_Id)
		errChan <- err
		spellChan <- spells
	}()

	go func() {
		defer func() {
			close(proficiencyChan)
			wg.Done()
		}()
		proficiencies, err := s.proficiencyService.GetByCharacterDataId(character.Character_Id)

		errChan <- err
		proficiencyChan <- proficiencies
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan{
		if err != nil {
			fmt.Println(err)
		return dto.FullCharacterData{}, err
		}
	}

	// armor, err := s.armorService.GetByCharacterDataIdArmor(character.Character_Id)
	// if err != nil {
	// 	fmt.Println("murio armor"+err.Error())
	// 	return dto.FullCharacterData{}, err
	// }
	armor := []domain.ArmorXCharacterData{}
	// skills, err := s.skillService.GetByCharacterId(character.Character_Id)
	// if err != nil {
	// 	fmt.Println("murio skils"+err.Error())
	// 	return dto.FullCharacterData{}, err
	// }
	skills := []domain.Skill{}

	return characterDataToFullCharacterData(*character, <-itemChan, <-weaponChan, armor, skills, <-featureChan, <-spellChan, <-proficiencyChan), nil
}
