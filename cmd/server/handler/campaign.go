package handler

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/campaign"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CampaignHandler struct {
	service campaign.CampaignService
}

func NewCampaignHandler(service *campaign.CampaignService) *CampaignHandler {
	return &CampaignHandler{service: *service}
}

func (h *CampaignHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempCampaign dto.CreateCampaignDto
		if err := ctx.BindJSON(&tempCampaign); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdCampaign, err := h.service.CreateCampaign(tempCampaign)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdCampaign)
	}
}

func (h *CampaignHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		campaignList, err := h.service.GetAllCampaigns()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, campaignList)
	}
}

func (h *CampaignHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempCampaign, err := h.service.GetCampaignByID(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempCampaign)
	}
}

func (h *CampaignHandler) HandlerGetByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		tempCampaign, err := h.service.GetCampaignsByUserId(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempCampaign)
	}
}

func (h *CampaignHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		var tempCampaign dto.CreateCampaignDto
		if err := ctx.BindJSON(&tempCampaign); err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}

		updatedCampaign, err := h.service.UpdateCampaign(tempCampaign, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedCampaign)
	}
}

func (h *CampaignHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		

		serviceErr := h.service.DeleteCampaign(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr)
			return
		}

		ctx.JSON(200, "Deleted Campaign with id "+id)
	}
}
