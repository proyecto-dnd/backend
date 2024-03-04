package armorXCharacterData

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	armorXcharacterRepo RepositoryArmorXCharacterData
}

// CreateArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) CreateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error) {
	newArmorRelationship, err := s.armorXcharacterRepo.CreateArmorXCharacterData(data)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}
	return newArmorRelationship, nil
}

// DeleteArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) DeleteArmorXCharacterData(id int64) error {
	err := s.armorXcharacterRepo.DeleteArmorXCharacterData(id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByCharacterDataIdArmor implements ServiceArmorXCharacterData.
func (s *service) DeleteByCharacterDataIdArmor(id int64) error {
	err := s.armorXcharacterRepo.DeleteByCharacterDataIdArmor(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) GetAllArmorXCharacterData() ([]domain.ArmorXCharacterData, error) {
	armorRelationships, err := s.armorXcharacterRepo.GetAllArmorXCharacterData()
	if err != nil {
		return []domain.ArmorXCharacterData{}, nil
	}
	return armorRelationships, nil
}

// GetByIdArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) GetByIdArmorXCharacterData(id int64) (domain.ArmorXCharacterData, error) {
	armorRelationship, err := s.armorXcharacterRepo.GetByIdArmorXCharacterData(id)
	if err != nil {
		return domain.ArmorXCharacterData{}, nil
	}
	return armorRelationship, nil
}

// GetByCharacterDataIdArmor implements ServiceArmorXCharacterData.
func (s *service) GetByCharacterDataIdArmor(id int64) ([]domain.ArmorXCharacterData, error) {
	armorRelationships, err := s.armorXcharacterRepo.GetByCharacterDataIdArmor(id)
	if err != nil {
		return []domain.ArmorXCharacterData{}, nil
	}
	return armorRelationships, nil
}

// UpdateArmorXCharacterData implements ServiceArmorXCharacterData.
func (s *service) UpdateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error) {
	newArmorRelationship, err := s.armorXcharacterRepo.UpdateArmorXCharacterData(data)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}
	return newArmorRelationship, nil
}

func NewServiceArmorXCharacterData(armorXcharacterRepo RepositoryArmorXCharacterData) ServiceArmorXCharacterData {
	return &service{armorXcharacterRepo: armorXcharacterRepo}
}
