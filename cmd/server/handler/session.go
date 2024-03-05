package handler

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/session"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type SessionHandler struct {
	service session.SessionService
}

func NewSessionHandler(service *session.SessionService) *SessionHandler {
	return &SessionHandler{service: *service}
}

func (h *SessionHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempSession dto.CreateSessionDto
		if err := ctx.BindJSON(&tempSession); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdSession, err := h.service.CreateSession(tempSession)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdSession)
	}
}

func (h *SessionHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionList, err := h.service.GetAllSessions()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, sessionList)
	}
}

func (h *SessionHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempSession, err := h.service.GetSessionById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempSession)
	}
}

func (h *SessionHandler) HandlerGetByCampaignId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempSession, err := h.service.GetSessionsByCampaignId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempSession)
	}
}

func (h *SessionHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		var tempSession dto.CreateSessionDto
		if err := ctx.BindJSON(&tempSession); err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}

		updatedSession, err := h.service.UpdateSession(tempSession, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedSession)
	}
}

func (h *SessionHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		

		serviceErr := h.service.DeleteSession(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr)
			return
		}

		ctx.JSON(200, "Deleted Session with id "+id)
	}
}
