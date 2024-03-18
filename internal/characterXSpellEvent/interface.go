package characterxspellevent

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CharacterXSpellEventRepository interface {
	GetAll() ([]domain.CharacterXSpellEvent, error)
	GetById(id int) (domain.CharacterXSpellEvent, error)
	GetByCharacterId(characterId int) ([]domain.CharacterXSpellEvent, error)
	GetBySpellEventId(spellEventId int) ([]domain.CharacterXSpellEvent, error)
	Create(CharacterXSpellEvent domain.CharacterXSpellEvent) (domain.CharacterXSpellEvent, error)
	Delete(id int) error
}

type CharacterXSpellEventService interface {
	GetAll() ([]domain.CharacterXSpellEvent, error)
	GetById(id int) (domain.CharacterXSpellEvent, error)
	GetByCharacterId(characterId int) ([]domain.CharacterXSpellEvent, error)
	GetBySpellEventId(spellEventId int) ([]domain.CharacterXSpellEvent, error)
	Create(CharacterXSpellEvent dto.CharacterXSpellEventDto) (domain.CharacterXSpellEvent, error)
	Delete(id int) error
}
