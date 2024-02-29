package weaponxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	weaponXCharacterDataRepo RepositoryWeaponXCharacterData
}

// Create implements ServiceItemXTableCharacter.
func (s *service) Create(weaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error) {
	newItemRelationship, err := s.weaponXCharacterDataRepo.Create(weaponXCharacterData)
	if err!= nil {
        return domain.WeaponXCharacterData{}, err
    }
	return newItemRelationship, nil
}

// Delete implements ServiceItemXTableCharacter.
func (s *service) Delete(id int64) error {
	err := s.weaponXCharacterDataRepo.Delete(id)
    if err!= nil {
        return err
    }
    return nil
}

// DeleteByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) DeleteByCharacterDataId(id int64) error {
	err := s.weaponXCharacterDataRepo.DeleteByCharacterDataId(id)
    if err!= nil {
        return err
    }
    return nil
}

// GetAll implements ServiceItemXTableCharacter.
func (s *service) GetAll() ([]domain.WeaponXCharacterData, error) {
	itemRelationships, err := s.weaponXCharacterDataRepo.GetAll()
	if err != nil {
		return []domain.WeaponXCharacterData{}, nil
	}
	return itemRelationships, nil
}

// GetByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) GetByCharacterDataId(id int64) ([]domain.WeaponXCharacterData, error) {
	itemRelationships, err := s.weaponXCharacterDataRepo.GetByCharacterDataId(id)
    if err!= nil {
        return []domain.WeaponXCharacterData{}, nil
    }
    return itemRelationships, nil
}

// GetById implements ServiceItemXTableCharacter.
func (s *service) GetById(id int64) (domain.WeaponXCharacterData, error) {
	itemRelationship, err := s.weaponXCharacterDataRepo.GetById(id)
    if err!= nil {
        return domain.WeaponXCharacterData{}, nil
    }
    return itemRelationship, nil
}

// Update implements ServiceItemXTableCharacter.
func (s *service) Update(weaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error) {
	newItemRelationship, err := s.weaponXCharacterDataRepo.Update(weaponXCharacterData)
    if err!= nil {
        return domain.WeaponXCharacterData{}, err
    }
    return newItemRelationship, nil
}

func NewWeaponXCharacterDataService(weaponXCharacterDataRepo RepositoryWeaponXCharacterData) ServiceWeaponXCharacterData {
	return &service{weaponXCharacterDataRepo: weaponXCharacterDataRepo}
}
