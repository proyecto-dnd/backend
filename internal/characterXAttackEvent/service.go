package characterxattackevent

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	characterXAttackEventRepository CharacterXAttackEventRepository
}

func NewCharacterXAttackEventService(characterXAttackEventRepository CharacterXAttackEventRepository) CharacterXAttackEventService {
	return &service{characterXAttackEventRepository: characterXAttackEventRepository}
}

func (s *service) GetAll() ([]domain.CharacterXAttackEvent, error) {
	characterXSpellEvents, err := s.characterXAttackEventRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return characterXSpellEvents, nil
}

func (s *service) GetById(id int) (domain.CharacterXAttackEvent, error) {
	characterXSpellEvent, err := s.characterXAttackEventRepository.GetById(id)
	if err != nil {
		return domain.CharacterXAttackEvent{}, err
	}

	return characterXSpellEvent, nil
}

func (s *service) GetByCharacterId(characterId int) ([]domain.CharacterXAttackEvent, error) {
	characterXSpellEvents, err := s.characterXAttackEventRepository.GetByCharacterId(characterId)
	if err != nil {
		return nil, err
	}

	return characterXSpellEvents, nil
}

func (s *service) GetBySpellEventId(spellEventId int) ([]domain.CharacterXAttackEvent, error) {
	characterXSpellEvents, err := s.characterXAttackEventRepository.GetBySpellEventId(spellEventId)
	if err != nil {
		return nil, err
	}

	return characterXSpellEvents, nil
}

func (s *service) Create(characterXSpellEvent dto.CharacterXAttackEventDto) (domain.CharacterXAttackEvent, error) {
	newCharacterXSpellEvent := domain.CharacterXAttackEvent{
		CharacterId: characterXSpellEvent.CharacterId,
		SpellEventId: characterXSpellEvent.SpellEventId,
	}

	createdCharacterXSpellEvent, err := s.characterXAttackEventRepository.Create(newCharacterXSpellEvent)
	if err != nil {
		return domain.CharacterXAttackEvent{}, err
	}

	return createdCharacterXSpellEvent, nil
}

func (s *service) Delete(id int) error {
	err := s.characterXAttackEventRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
