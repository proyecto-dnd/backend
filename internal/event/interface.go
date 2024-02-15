package event

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type RepositoryEvent interface {
	Create(event domain.Event) (domain.Event, error)
	GetAll()([]domain.Event, error)
	GetById(id int)(domain.Event, error)
	GetBySessionId(sessionid int)([]domain.Event, error)
	GetByCharacterId(characterid int)([]domain.Event, error)
	Update(event domain.Event) 
	Delete(id int)error
}