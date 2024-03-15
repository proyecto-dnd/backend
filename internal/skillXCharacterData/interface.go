package skillxcharacterdata

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositorySkillXCharacter interface {
	Create(skillXCharacterData domain.SkillXCharacterData) (domain.SkillXCharacterData, error)
	Delete(skillXCharacterData domain.SkillXCharacterData) error
}

type ServiceSkillXCharacter interface {
	Create(skillXCharacterData domain.SkillXCharacterData) (domain.SkillXCharacterData, error)
	Delete(skillXCharacterData domain.SkillXCharacterData) error
}
