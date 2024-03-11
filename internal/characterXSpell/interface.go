package characterXspell

import "github.com/proyecto-dnd/backend/internal/domain"

type CharacterXSpellRepository interface {
	Create(CharacterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error)
	Delete(id int) error
}

type CharacterXSpellService interface {
	Create(CharacterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error)
	Delete(id int) error
}
