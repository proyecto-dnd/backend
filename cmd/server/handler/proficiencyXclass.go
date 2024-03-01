package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/proficiencyXclass.go"
)

type ProficiencyXClassHandler struct {
	service proficiencyXclass.ProficiencyXClassService
}

func NewProficiencyXClassHandler(service proficiencyXclass.ProficiencyXClassService) *ProficiencyXClassHandler {
	return &ProficiencyXClassHandler{service: service}
}

func (h *ProficiencyXClassHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempProficiencyXClass domain.ProficiencyXClass
		if err := ctx.BindJSON(&tempProficiencyXClass); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdProficiencyXClass, err := h.service.Create(tempProficiencyXClass)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(201, createdProficiencyXClass)
	}
}

func (h *ProficiencyXClassHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempProficiencyXClass domain.ProficiencyXClass
		if err := ctx.BindJSON(&tempProficiencyXClass); err != nil {
			ctx.JSON(500, err)
			return
		}
		if err := h.service.Delete(tempProficiencyXClass); err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, "Deleted row with class_id: "+strconv.Itoa(tempProficiencyXClass.ClassId)+" and proficiency_id: "+strconv.Itoa(tempProficiencyXClass.ProficiencyId))
	}
}
