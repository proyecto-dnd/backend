package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/feature"
)

type FeatureHandler struct {
	service feature.FeatureService
}

func NewFeatureHandler(service *feature.FeatureService) *FeatureHandler {
	return &FeatureHandler{service: *service}
}

func (h *FeatureHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempFeature dto.CreateFeatureDto
		if err := ctx.BindJSON(&tempFeature); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdFeature, err := h.service.CreateFeature(tempFeature)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdFeature)
	}
}

func (h *FeatureHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		featureList, err := h.service.GetAllFeatures()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, featureList)
	}
}

func (h *FeatureHandler) HandlerGetAllFeaturesByCharacterId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		characterId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		featureList, err := h.service.GetAllFeaturesByCharacterId(characterId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, featureList)
	}
}

func (h *FeatureHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempFeature, err := h.service.GetFeatureById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempFeature)
	}
}

func (h *FeatureHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		var tempFeature dto.CreateFeatureDto
		if err := ctx.BindJSON(&tempFeature); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdFeature, err := h.service.UpdateFeature(tempFeature, id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdFeature)
	}
}

func (h *FeatureHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		err = h.service.DeleteFeature(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "Feature deleted")
	}
}