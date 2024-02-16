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
	HandlerGetBySessionId() gin.HandlerFunc
	HandlerGetByCharacterId() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type EventService interface {
	CreateEvent(dto.EventDto) (dto.EventDto, error)
	GetAllEvents() ([]dto.EventDto, error)
	GetEventByID(id int) (dto.EventDto, error)
	GetEventsBySessionID(sessionid int) ([]dto.EventDto, error)
	GetEventsByCharacterID(characterid int) ([]dto.EventDto, error)
	UpdateEvent(dto.EventDto) (dto.EventDto, error)
	DeleteEvent(id int) error
}

type EventRepository interface {
	Create(event domain.Event) (domain.Event, error)
	GetAll()([]domain.Event, error)
	GetById(id int)(domain.Event, error)
	GetBySessionId(sessionid int)([]domain.Event, error)
	GetByCharacterId(characterid int)([]domain.Event, error)
	Update(event domain.Event) (domain.Event, error)
	Delete(id int)error
}