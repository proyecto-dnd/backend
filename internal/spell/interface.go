package spell

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type RepositorySpell interface {
	Create(spell dto.SpellDto) (domain.Spell, error)
	GetAll() ([]domain.Spell, error)
	GetById(id int) (domain.Spell, error)
	Update(item dto.SpellDto, id int) (domain.Spell, error)
	Delete(id int) error
}


type ServiceSpell interface {
	Create(spell domain.Spell) (domain.Spell, error)
	GetAll() ([]domain.Spell, error)
	GetById(id int) (domain.Spell, error)
	Update(item domain.Spell, id int) (domain.Spell, error)
	Delete(id int) error
}
