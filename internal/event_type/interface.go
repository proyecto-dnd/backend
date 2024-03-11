package event_type

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type EventTypeHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerGetByName() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type EventTypeService interface {
	CreateEventType(eventType dto.CreateEventTypeDto) (domain.EventType, error)
	GetAllEventTypes() ([]domain.EventType, error)
	GetEventTypeById(id int) (domain.EventType, error)
	GetEventTypeByName(name string) (domain.EventType, error)
	UpdateEventType(eventType dto.CreateEventTypeDto, id int) (domain.EventType, error)
	DeleteEventType(id int) error
}

type EventTypeRepository interface {
	Create(eventType domain.EventType) (domain.EventType, error)
	GetAll()([]domain.EventType, error)
	GetById(id int)(domain.EventType, error)
	GetByName(name string)(domain.EventType, error)
	Update(eventType domain.EventType, id int)(domain.EventType, error)
	Delete(id int) error
}