package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/proficiency"
)

type ProficiencyHandler struct {
	service proficiency.ProficiencyService
}

func NewProficiencyHandler(service *proficiency.ProficiencyService) *ProficiencyHandler {
	return &ProficiencyHandler{service: *service}
}

func (h *ProficiencyHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempProficiency dto.ProficiencyDto
		if err := ctx.BindJSON(&tempProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdProficiency, err := h.service.Create(tempProficiency)
		if err != nil {
			ctx.JSON(500, err)
		}
		ctx.JSON(201, createdProficiency)
	}
}

func (h *ProficiencyHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		proficiencyList, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, proficiencyList)
	}
}

func (h *ProficiencyHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		proficiency, err := h.service.GetById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, proficiency)
	}
}

func (h *ProficiencyHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		var tempProficiency dto.ProficiencyDto
		if err := ctx.BindJSON(&tempProficiency); err != nil {
			ctx.JSON(500, err)
			return
		}
		updatedProficiency, err := h.service.Update(tempProficiency, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedProficiency)
	}
}

func (h *ProficiencyHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		if err = h.service.Delete(intId); err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Deleted Proficiency with id: "+id)
	}
}
