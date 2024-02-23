package skillxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositorySkillXCharacterData interface {
	Create(skillXCharacterData domain.SkillXCharacterData) (domain.SkillXCharacterData, error)
	Delete(skillXCharacterData domain.SkillXCharacterData) error
}