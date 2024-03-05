package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	itemxcharacterdata "github.com/proyecto-dnd/backend/internal/itemXCharacterData"
)

type ItemXCharacterDataHandler struct {
	service itemxcharacterdata.ServiceItemXCharacterData
}

func NewItemXCharacterDataHandler(service *itemxcharacterdata.ServiceItemXCharacterData) *ItemXCharacterDataHandler {
    return &ItemXCharacterDataHandler{service: *service}
}

func (h *ItemXCharacterDataHandler) HandlerCreate() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var tempItemXCharacterData domain.ItemXCharacterData
		if err := ctx.BindJSON(&tempItemXCharacterData); err != nil{
			ctx.AbortWithError(400, err)
			return
		}
		createdItemXCharacterData, err := h.service.Create(tempItemXCharacterData)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(201, createdItemXCharacterData)
	}
}

func (h *ItemXCharacterDataHandler) HandlerDelete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err!= nil{
			ctx.AbortWithError(400, err)
			return
		}
		err = h.service.Delete(id)
		if err!= nil{
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(204, nil)
	}
}

func (h *ItemXCharacterDataHandler) HandlerDeleteByCharacterId() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err!= nil{
			ctx.AbortWithError(400, err)
			return
		}
		err = h.service.DeleteByCharacterDataId(id)
		if err!= nil{
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(204, nil)
	}
}

func (h *ItemXCharacterDataHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemXCharacterDataList, err := h.service.GetAll()
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, itemXCharacterDataList)
	}
}

func (h *ItemXCharacterDataHandler) HandlerGetByCharacterDataId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		itemXCharacterDataList, err := h.service.GetByCharacterDataId(id)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, itemXCharacterDataList)
	}
}


func (h *ItemXCharacterDataHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		itemXCharacterData, err := h.service.GetById(id)
		if err != nil {
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(200, itemXCharacterData)
	}
}

func (h *ItemXCharacterDataHandler) HandlerUpdate() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var tempItemXCharacterData domain.ItemXCharacterData
		if err := ctx.BindJSON(&tempItemXCharacterData); err != nil{
			ctx.AbortWithError(400, err)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}

		tempItemXCharacterData.Character_Item_Id = id

		updatedItemXCharacterData, err := h.service.Update(tempItemXCharacterData)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, updatedItemXCharacterData)
	}
}