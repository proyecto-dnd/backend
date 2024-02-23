package handler

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/user_campaign"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type UserCampaignHandler struct {
	service user_campaign.UserCampaignService
}

func NewUserCampaignHandler(service *user_campaign.UserCampaignService) *UserCampaignHandler {
	return &UserCampaignHandler{service: *service}
}

func (h *UserCampaignHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempUserCampaign dto.CreateUserCampaignDto
		if err := ctx.BindJSON(&tempUserCampaign); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdUserCampaign, err := h.service.CreateUserCampaign(tempUserCampaign)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdUserCampaign)
	}
}

func (h *UserCampaignHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userCampaignList, err := h.service.GetAllUserCampaigns()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, userCampaignList)
	}
}

func (h *UserCampaignHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempUserCampaign, err := h.service.GetUserCampaignByID(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, tempUserCampaign)
	}
}

func (h *UserCampaignHandler) HandlerGetByCampaignId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempUserCampaign, err := h.service.GetUserCampaignByCampaignID(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, tempUserCampaign)
	}
}

func (h *UserCampaignHandler) HandlerGetByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		tempUserCampaign, err := h.service.GetUserCampaignByUserId(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, tempUserCampaign)
	}
}

func (h *UserCampaignHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		err = h.service.DeleteUserCampaign(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "User Campaign deleted")
	}
}