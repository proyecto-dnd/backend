package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/event_type"
)

type EventTypeHandler struct {
	service event_type.EventTypeService
}

func NewEventTypeHandler(service *event_type.EventTypeService) *EventTypeHandler {
	return &EventTypeHandler{service: *service}
}

// eventType godoc
// @Summary Create eventType
// @Tags eventType
// @Accept json
// @Produce json
// @Param body body dto.CreateEventTypeDto true "CreateEventTypeDto"
// @Success 201 {object} domain.EventType
// @Failure 500 {object} error
// @Router /event_type [post]
func (h *EventTypeHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempEventType dto.CreateEventTypeDto
		if err := ctx.BindJSON(&tempEventType); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdEventType, err := h.service.CreateEventType(tempEventType)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdEventType)
	}
}

// eventType godoc
// @Summary Get all eventTypes
// @Tags eventType
// @Produce json
// @Success 200 {array} domain.EventType
// @Failure 500 {object} error
// @Router /event_type [get]
func (h *EventTypeHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventTypeList, err := h.service.GetAllEventTypes()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventTypeList)
	}
}

// eventType godoc
// @Summary Get eventType by id
// @Tags eventType
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.EventType
// @Failure 500 {object} error
// @Router /event_type/{id} [get]
func (h *EventTypeHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempEventType, err := h.service.GetEventTypeById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempEventType)
	}
}

// eventType godoc
// @Summary Get eventType by name
// @Tags eventType
// @Produce json
// @Param name path string true "name"
// @Success 200 {object} domain.EventType
// @Failure 500 {object} error
// @Router /event_type/{name} [get]
func (h *EventTypeHandler) HandlerGetByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		tempEventType, err := h.service.GetEventTypeByName(name)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempEventType)
	}
}

// eventType godoc
// @Summary Update eventType
// @Tags eventType
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.CreateEventTypeDto true "CreateEventTypeDto"
// @Success 200 {object} domain.EventType
// @Failure 500 {object} error
// @Router /event_type/{id} [put]
func (h *EventTypeHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempEventType dto.CreateEventTypeDto
		if err := ctx.BindJSON(&tempEventType); err != nil {
			ctx.JSON(500, err)
			return
		}

		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		updatedEventType, err := h.service.UpdateEventType(tempEventType, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, updatedEventType)
	}
}

// eventType godoc
// @Summary Delete eventType
// @Tags eventType
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /event_type/{id} [delete]
func (h *EventTypeHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		err = h.service.DeleteEventType(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Deleted")
	}
}