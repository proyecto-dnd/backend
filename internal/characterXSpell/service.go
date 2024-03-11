package characterXspell

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	characterXSpellRepository CharacterXSpellRepository
}

func NewCharacterXSpellService(characterXSpellRepository CharacterXSpellRepository) CharacterXSpellService {
	return &service{characterXSpellRepository: characterXSpellRepository}
}

func (s *service) Create(characterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error) {
	createdCharacterXSpell, err := s.characterXSpellRepository.Create(characterXSpell)
	if err != nil {
		return domain.CharacterXSpell{}, err
	}
	return createdCharacterXSpell, nil
}

func (s *service) Delete(id int) error {
	return s.characterXSpellRepository.Delete(id)
}
