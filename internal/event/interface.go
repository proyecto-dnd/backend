package event

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type EventHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerGetEventByTypeId() gin.HandlerFunc
	HandlerGetBySessionId() gin.HandlerFunc
	HandlerGetByProtagonistId() gin.HandlerFunc
	HandlerGetCharactersAffectedByEventId() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type EventService interface {
	CreateEvent(dto.CreateEventDto) (domain.Event, error)
	GetAllEvents() ([]dto.ResponseEventDto, error)
	GetEventById(id int) (dto.ResponseEventDto, error)
	GetEventsByTypeId(typeId int) ([]dto.ResponseEventDto, error)
	GetEventsBySessionId(sessionid int) ([]dto.ResponseEventDto, error)
	GetEventsByProtagonistId(protagonistId int) ([]dto.ResponseEventDto, error)
	GetCharactersAffectedByEventId(eventid int) ([]dto.ResponseEventDto, error)
	UpdateEvent(event dto.CreateEventDto, id int) (domain.Event, error)
	DeleteEvent(id int) error
}

type EventRepository interface {
	Create(event domain.Event) (domain.Event, error)
	GetAll() ([]dto.EventRepositoryResponseDto, error)
	GetById(id int) (dto.EventRepositoryResponseDto, error)
	GetByTypeId(typeid int) ([]dto.EventRepositoryResponseDto, error)
	GetBySessionId(sessionid int) ([]dto.EventRepositoryResponseDto, error)
	GetByProtagonistId(protagonistid int) ([]dto.EventRepositoryResponseDto, error)
	Update(event domain.Event, id int) (domain.Event, error)
	Delete(id int) error
}
