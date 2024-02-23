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
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type EventService interface {
	CreateEvent(dto.CreateEventDto) (domain.Event, error)
	GetAllEvents() ([]domain.Event, error)
	GetEventById(id int) (domain.Event, error)
	GetEventByTypeId(typeId int) ([]domain.Event, error)
	GetEventsBySessionId(sessionid int) ([]domain.Event, error)
	GetEventsByProtagonistId(protagonistId int) ([]domain.Event, error)
	UpdateEvent(event dto.CreateEventDto, id int) (domain.Event, error)
	DeleteEvent(id int) error
}

type EventRepository interface {
	Create(event domain.Event) (domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetById(id int) (domain.Event, error)
	GetByTypeId(typeid int) ([]domain.Event, error)
	GetBySessionId(sessionid int) ([]domain.Event, error)
	GetByProtagonistId(protagonistid int) ([]domain.Event, error)
	Update(event domain.Event, id int) (domain.Event, error)
	Delete(id int) error
}
