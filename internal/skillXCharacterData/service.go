package skillxcharacterdata

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type service struct {
	repository RepositorySkillXCharacterData
}

func NewService(repository RepositorySkillXCharacterData) *service {
	return &service{repository: repository}
}

func (s *service) Create(skillXCharacterData domain.SkillXCharacterData) (domain.SkillXCharacterData, error) {
	newSkillXCharacter, err := s.repository.Create(skillXCharacterData)
	if err != nil {
		return domain.SkillXCharacterData{}, err
	}

	return newSkillXCharacter, nil
}

func (s *service) Delete(skillXCharacterData domain.SkillXCharacterData) error {
	return s.repository.Delete(skillXCharacterData)
}
