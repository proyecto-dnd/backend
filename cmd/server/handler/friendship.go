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
func (h *FriendshipHandler) HandlerCreate() gin.HandlerFunc {
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
func (h *FriendshipHandler) HandlerDelete() gin.HandlerFunc {
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

func (h *FriendshipHandler) HandlerSearchFollowers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		var tempFriendship domain.Mutuals
		tempFriendship.User2Name = name
		if err := ctx.BindJSON(&tempFriendship); err != nil {
			ctx.JSON(500, err)
			return
		}
		followers, err := h.service.SearchFollowers(tempFriendship)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, followers)
	}
}

func (h *FriendshipHandler) HandlerGetAllFriends() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userId string

		if err := ctx.BindJSON(&userId); err != nil {
			ctx.JSON(500, err)
			return
		}

		friends, err := h.service.GetAllFriends(userId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, friends)
	}
}

func (h *FriendshipHandler) HandlerGetBySimilarName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		users, err := h.service.GetBySimilarName(name)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, users)
	}
}
