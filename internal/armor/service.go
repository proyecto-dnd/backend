package armor

import (
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type armorService struct {
	armorRepo ArmorRepository
}

func NewArmorService(armorRepo ArmorRepository) ArmorService {
	return &armorService{armorRepo: armorRepo}
}

func (s *armorService) CreateArmor(armorDto dto.CreateArmorDto) (domain.Armor, error) {
	armorDomain := domain.Armor{
		Material:       armorDto.Material,
		Name:           armorDto.Name,
		Weight:         armorDto.Weight,
		Price:          armorDto.Price,
		Category:       armorDto.Category,
		ProtectionType: armorDto.ProtectionType,
		Description:    armorDto.Description,
		Penalty:        armorDto.Penalty,
		Strength:       armorDto.Strength,
		ArmorClass:     armorDto.ArmorClass,
		DexBonus:       armorDto.DexBonus,
		Basic:          armorDto.Basic,
	}

	createdArmor, err := s.armorRepo.Create(armorDomain)
	if err != nil {
		return domain.Armor{}, err
	}

	return createdArmor, nil
}

func (s *armorService) GetAllArmor() ([]domain.Armor, error) {
	armors, err := s.armorRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return armors, nil
}

func (s *armorService) GetArmorByID(id int) (domain.Armor, error) {
	armor, err := s.armorRepo.GetById(id)
	if err != nil {
		return domain.Armor{}, err
	}

	return armor, nil
}

func (s *armorService) UpdateArmor(armorDto dto.CreateArmorDto, id int) (domain.Armor, error) {
	armorDomain := domain.Armor{
		Material:       armorDto.Material,
		Name:           armorDto.Name,
		Weight:         armorDto.Weight,
		Price:          armorDto.Price,
		Category:       armorDto.Category,
		ProtectionType: armorDto.ProtectionType,
		Description:    armorDto.Description,
		Penalty:        armorDto.Penalty,
		Strength:       armorDto.Strength,
		ArmorClass:     armorDto.ArmorClass,
		DexBonus:       armorDto.DexBonus,
		Basic:          armorDto.Basic,
	}

	updatedArmor, err := s.armorRepo.Update(armorDomain, id)
	if err != nil {
		fmt.Println(err)
		return domain.Armor{}, err
	}

	return updatedArmor, nil
}

func (s *armorService) DeleteArmor(id int) error {
	return s.armorRepo.Delete(id)
}
