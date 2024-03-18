package characterxspellevent

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	characterXSpellEventRepository CharacterXSpellEventRepository
}

func NewCharacterXSpellEventService(characterXSpellEventRepository CharacterXSpellEventRepository) CharacterXSpellEventService {
	return &service{characterXSpellEventRepository: characterXSpellEventRepository}
}

func (s *service) GetAll() ([]domain.CharacterXSpellEvent, error) {
	characterXSpellEvents, err := s.characterXSpellEventRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return characterXSpellEvents, nil
}

func (s *service) GetById(id int) (domain.CharacterXSpellEvent, error) {
	characterXSpellEvent, err := s.characterXSpellEventRepository.GetById(id)
	if err != nil {
		return domain.CharacterXSpellEvent{}, err
	}

	return characterXSpellEvent, nil
}

func (s *service) GetByCharacterId(characterId int) ([]domain.CharacterXSpellEvent, error) {
	characterXSpellEvents, err := s.characterXSpellEventRepository.GetByCharacterId(characterId)
	if err != nil {
		return nil, err
	}

	return characterXSpellEvents, nil
}

func (s *service) GetBySpellEventId(spellEventId int) ([]domain.CharacterXSpellEvent, error) {
	characterXSpellEvents, err := s.characterXSpellEventRepository.GetBySpellEventId(spellEventId)
	if err != nil {
		return nil, err
	}

	return characterXSpellEvents, nil
}

func (s *service) Create(characterXSpellEvent dto.CharacterXSpellEventDto) (domain.CharacterXSpellEvent, error) {
	newCharacterXSpellEvent := domain.CharacterXSpellEvent{
		CharacterId: characterXSpellEvent.CharacterId,
		SpellEventId: characterXSpellEvent.SpellEventId,
	}

	createdCharacterXSpellEvent, err := s.characterXSpellEventRepository.Create(newCharacterXSpellEvent)
	if err != nil {
		return domain.CharacterXSpellEvent{}, err
	}

	return createdCharacterXSpellEvent, nil
}

func (s *service) Delete(id int) error {
	err := s.characterXSpellEventRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
