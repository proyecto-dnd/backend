package handler

import (
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

// event godoc
// @Summary Create event
// @Tags event
// @Accept json
// @Produce json
// @Param body body dto.CreateEventDto true "CreateEventDto"
// @Success 201 {object} domain.Event
// @Failure 500 {object} error
// @Router /event [post]
func (h *EventHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempEvent dto.CreateEventDto
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

// event godoc
// @Summary Get all events
// @Tags event
// @Produce json
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event [get]
func (h *EventHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		eventList, err := h.service.GetAllEvents()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Get event by id
// @Tags event
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/{id} [get]
func (h *EventHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempEvent, err := h.service.GetEventById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempEvent)
	}
}

// event godoc
// @Summary Get event by id
// @Tags event
// @Produce json
// @Param id path int true "type_id"
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/type/{id} [get]
func (h *EventHandler) HandlerGetByTypeId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		eventList, err := h.service.GetEventsByTypeId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Get event by session id
// @Tags event
// @Produce json
// @Param id path int true "session_id"
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/session/{id} [get]
func (h *EventHandler) HandlerGetBySessionId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		eventList, err := h.service.GetEventsBySessionId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventList)
	}
}


// event godoc
// @Summary Get event by protagonist id
// @Tags event
// @Produce json
// @Param id path int true "protagonist_id"
// @Success 200 {array} dto.ResponseEventDto
// @Failure 500 {object} error
// @Router /event/protagonist/{id} [get]
func (h *EventHandler) HandlerGetByProtagonistId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		eventList, err := h.service.GetEventsByProtagonistId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, eventList)
	}
}

// event godoc
// @Summary Update event
// @Tags event
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body dto.CreateEventDto true "CreateEventDto"
// @Success 200 {object} domain.Event
// @Failure 500 {object} error
// @Router /event/{id} [put]
func (h *EventHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		var tempEvent dto.CreateEventDto
		if err := ctx.BindJSON(&tempEvent); err != nil {
			ctx.JSON(500, err)
			return
		}

		updatedEvent, err := h.service.UpdateEvent(tempEvent, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedEvent)
	}
}

func (h *EventHandler) HandlerGetCharactersAffectedByEventId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempEvent, err := h.service.GetCharactersAffectedByEventId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempEvent)
	}
}

// event godoc
// @Summary Delete event
// @Tags event
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /event/{id} [delete]
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

		ctx.JSON(200, "Deleted Event with id "+id)
	}
}
