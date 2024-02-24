package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/item"
)

type ItemHandler struct {
	service item.ServiceItem
}

func NewItemHandler(service *item.ServiceItem) *ItemHandler {
    return &ItemHandler{service: *service}
}

func (h *ItemHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempItem domain.Item
        if err := ctx.BindJSON(&tempItem); err!= nil {
            ctx.JSON(400, err)
            return
        }
		createdItem, err := h.service.Create(tempItem)
		if err!= nil {
            ctx.JSON(500, err)
            return
        }
		ctx.JSON(201, createdItem)
    }
}

func (h *ItemHandler) HandlerDelete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
        id, err := strconv.Atoi(ctx.Param("id"))
        if err!= nil {
            ctx.JSON(400, err)
            return
        }
        err = h.service.Delete(int64(id))
        if err!= nil {
            ctx.JSON(404, err)
            return
        }
        ctx.JSON(204, nil)
    }
}

func (h *ItemHandler) HandlerGetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		items, err := h.service.GetAll()
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        ctx.JSON(200, items)
	}
}

func (h *ItemHandler) HandlerGetByCampaignId() gin.HandlerFunc{
	return func(ctx *gin.Context) {
        id := ctx.Param("id")

        intId, err := strconv.Atoi(id)
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        items, err := h.service.GetByCampaignId(int64(intId))
        if err!= nil {
            ctx.JSON(500, err)
            return
        }
        ctx.JSON(200, items)
    }
}