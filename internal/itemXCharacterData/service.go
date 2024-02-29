package itemxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	itemXCharacterDataRepo RepositoryItemXCharacterData
}

// Create implements ServiceItemXTableCharacter.
func (s *service) Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	newItemRelationship, err := s.itemXCharacterDataRepo.Create(itemXCharacterData)
	if err!= nil {
        return domain.ItemXCharacterData{}, err
    }
	return newItemRelationship, nil
}

// Delete implements ServiceItemXTableCharacter.
func (s *service) Delete(id int64) error {
	err := s.itemXCharacterDataRepo.Delete(id)
    if err!= nil {
        return err
    }
    return nil
}

// DeleteByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) DeleteByCharacterDataId(id int64) error {
	err := s.itemXCharacterDataRepo.DeleteByCharacterDataId(id)
    if err!= nil {
        return err
    }
    return nil
}

// GetAll implements ServiceItemXTableCharacter.
func (s *service) GetAll() ([]domain.ItemXCharacterData, error) {
	itemRelationships, err := s.itemXCharacterDataRepo.GetAll()
	if err != nil {
		return []domain.ItemXCharacterData{}, nil
	}
	return itemRelationships, nil
}

// GetByCharacterDataId implements ServiceItemXTableCharacter.
func (s *service) GetByCharacterDataId(id int64) ([]domain.ItemXCharacterData, error) {
	itemRelationships, err := s.itemXCharacterDataRepo.GetByCharacterDataId(id)
    if err!= nil {
        return []domain.ItemXCharacterData{}, nil
    }
    return itemRelationships, nil
}

// GetById implements ServiceItemXTableCharacter.
func (s *service) GetById(id int64) (domain.ItemXCharacterData, error) {
	itemRelationship, err := s.itemXCharacterDataRepo.GetById(id)
    if err!= nil {
        return domain.ItemXCharacterData{}, nil
    }
    return itemRelationship, nil
}

// Update implements ServiceItemXTableCharacter.
func (s *service) Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	newItemRelationship, err := s.itemXCharacterDataRepo.Update(itemXCharacterData)
    if err!= nil {
        return domain.ItemXCharacterData{}, err
    }
    return newItemRelationship, nil
}

func NewItemXCharacterDataService(itemXCharacterDataRepo RepositoryItemXCharacterData) ServiceItemXCharacterData {
	return &service{itemXCharacterDataRepo: itemXCharacterDataRepo}
}
