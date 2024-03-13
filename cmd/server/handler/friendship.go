package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/friendship"
)

type FriendshipHandler struct {
	service friendship.FriendshipService
}

func NewFriendshipHandler(service *friendship.FriendshipService) *FriendshipHandler {
	return &FriendshipHandler{service: *service}
}

// friendship godoc
// @Summary Create friendship
// @Tags friendship
// @Accept json
// @Produce json
// @Param body body domain.Friendship true "Friendship"
// @Success 201 {object} domain.Friendship
// @Failure 500 {object} error
// @Router /friendship [post]
func (h *FriendshipHandler) CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempFriendship domain.Friendship
		if err := ctx.BindJSON(&tempFriendship); err != nil {
			ctx.JSON(500, err)
			return
		}
		createdFriendship, err := h.service.Create(tempFriendship)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdFriendship)
	}
}

// friendship godoc
// @Summary Delete friendship
// @Tags friendship
// @Accept json
// @Produce json
// @Param body body domain.Friendship true "Friendship"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /friendship [delete]
func (h *FriendshipHandler) DeleteHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempFriendship domain.Friendship
		if err := ctx.BindJSON(&tempFriendship); err != nil {
			ctx.JSON(500, err)
			return
		}
		_, err := h.service.Create(tempFriendship)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, "Deleted succesfully friendship with user1_id: "+tempFriendship.User1Id+" and user2_id: "+tempFriendship.User2Id)
	}
}
