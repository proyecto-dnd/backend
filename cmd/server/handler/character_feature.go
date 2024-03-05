package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/character_feature"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CharacterFeature struct {
	service character_feature.CharacterFeatureService
}

func NewCharacterFeatureHandler(service *character_feature.CharacterFeatureService) *CharacterFeature {
	return &CharacterFeature{service: *service}
}

func (h *CharacterFeature) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempCharacterFeature dto.CreateCharacterFeatureDto
		if err := ctx.BindJSON(&tempCharacterFeature); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdCharacterFeature, err := h.service.CreateCharacterFeature(tempCharacterFeature)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdCharacterFeature)
	}
}

func (h *CharacterFeature) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterFeatures, err := h.service.GetAllCharacterFeatures()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, characterFeatures)
	}
}

func (h *CharacterFeature) HandlerGetByFeatureId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempCharacterFeature, err := h.service.GetCharacterFeatureByFeatureId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, tempCharacterFeature)
	}
}

func (h *CharacterFeature) HandlerGetByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempCharacterFeature, err := h.service.GetCharacterFeatureByCharacterId(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, tempCharacterFeature)
	}
}

func (h *CharacterFeature) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idFeature, err := strconv.Atoi(ctx.Param("idFeature"))
		idCharacter, err := strconv.Atoi(ctx.Param("idCharacter"))
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		err = h.service.DeleteCharacterFeature(idFeature, idCharacter)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(204, nil)
	}
}