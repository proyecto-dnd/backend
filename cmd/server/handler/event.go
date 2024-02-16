package handler

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/event"
)

type EventHandler struct {
	service event.EventService
}

func NewEventHandler(service *event.EventService) *EventHandler {
	return &EventHandler{service: *service}
}

func (h *EventHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempEvent dto.EventDto
		if err := ctx.BindJSON(&tempEvent); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdEvent, err := h.service.CreateEvent(tempEvent)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdEvent)
	}
}

func (h *EventHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventList, err := h.service.GetAllEvents()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventList)
	}
}

func (h *EventHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempEvent, err := h.service.GetEventByID(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempEvent)
	}
}

func (h *EventHandler) HandlerGetBySessionId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		eventList, err := h.service.GetEventsBySessionID(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventList)
	}
}

func (h *EventHandler) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		eventList, err := h.service.GetEventsByCharacterID(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventList)
	}
}

func (h *EventHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempEvent dto.EventDto
		if err := ctx.BindJSON(&tempEvent); err != nil {
			ctx.JSON(500, err)
			return
		}

		updatedUser, err := h.service.UpdateEvent(tempEvent)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedUser)
	}
}

func (h *EventHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		

		serviceErr := h.service.DeleteEvent(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr)
			return
		}

		ctx.JSON(200, "Deleted User with id "+id)
	}
}
