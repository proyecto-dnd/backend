package spell

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositorySpell interface {
	Create(spell domain.Spell) (domain.Spell, error)
	GetAll() ([]domain.Spell, error)
	GetById(id int) (domain.Spell, error)
	Update(item domain.Spell, id int) (domain.Spell, error)
	Delete(id int) error
}


type ServiceSpell interface {
	Create(spell domain.Spell) (domain.Spell, error)
	GetAll() ([]domain.Spell, error)
	GetById(id int) (domain.Spell, error)
	Update(item domain.Spell, id int) (domain.Spell, error)
	Delete(id int) error
}
