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