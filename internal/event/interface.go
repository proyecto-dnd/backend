package event

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type EventRepository interface {
	Create(event domain.Event) (domain.Event, error)
	GetAll()([]domain.Event, error)
	GetById(id int)(domain.Event, error)
	GetBySessionId(sessionid int)([]domain.Event, error)
	GetByCharacterId(characterid int)([]domain.Event, error)
	Update(event domain.Event) (domain.Event, error)
	Delete(id int)error
}

type EventService interface {
	CreateEvent(dto.EventDto) (dto.EventDto, error)
	GetAllEvents() ([]dto.EventDto, error)
	GetEventByID(id int) (dto.EventDto, error)
	UpdateEvent(dto.EventDto) (dto.EventDto, error)
	DeleteEvent(id int) error
}