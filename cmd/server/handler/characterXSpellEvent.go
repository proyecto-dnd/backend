package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	characterXSpellEvent "github.com/proyecto-dnd/backend/internal/characterXSpellEvent"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CharacterXSpellEventHandler struct {
	service characterXSpellEvent.CharacterXSpellEventService
}

func NewCharacterXSpellEventHandler(service characterXSpellEvent.CharacterXSpellEventService) *CharacterXSpellEventHandler {
	return &CharacterXSpellEventHandler{service: service}
}

func (h *CharacterXSpellEventHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterXSpellEvents, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEvents)
	}
}

func (h *CharacterXSpellEventHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		characterXSpellEventId, err := h.service.GetById(idInt)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEventId)
	}
}

func (h *CharacterXSpellEventHandler) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterId := ctx.Param("characterId")

		characterIdInt, err := strconv.Atoi(characterId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		characterXSpellEvents, err := h.service.GetByCharacterId(characterIdInt)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEvents)
	}
}

func (h *CharacterXSpellEventHandler) HandlerGetBySpellEventId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		spellEventId := ctx.Param("spellEventId")

		spellEventIdInt, err := strconv.Atoi(spellEventId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		characterXSpellEvents, err := h.service.GetBySpellEventId(spellEventIdInt)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, characterXSpellEvents)
	}
}

func (h *CharacterXSpellEventHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempCharacterXSpellEvent dto.CharacterXSpellEventDto
		if err := ctx.BindJSON(&tempCharacterXSpellEvent); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdCharacterXSpellEvent, err := h.service.Create(tempCharacterXSpellEvent)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdCharacterXSpellEvent)
	}
}

func (h *CharacterXSpellEventHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		err = h.service.Delete(int(intId))
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Deleted characterXspellevent with id "+id)
	}
}