package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
)

type WeaponXCharacterDataHandler struct {
	service weaponxcharacterdata.ServiceWeaponXCharacterData
}

func NewWeaponXCharacterDataHandler(service *weaponxcharacterdata.ServiceWeaponXCharacterData) *WeaponXCharacterDataHandler {
    return &WeaponXCharacterDataHandler{service: *service}
}

func (h * WeaponXCharacterDataHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeaponXCharacterData domain.WeaponXCharacterData
		if err := ctx.BindJSON(&tempWeaponXCharacterData); err != nil {
			ctx.AbortWithError(400, err)
            return
		}
		createdWeaponXCharacterData, err := h.service.Create(tempWeaponXCharacterData)
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.JSON(201, createdWeaponXCharacterData)
	}
}

func (h *WeaponXCharacterDataHandler) HandlerDelete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil{
			ctx.AbortWithError(400, err)
            return
		}
		err = h.service.Delete(int64(id))
		if err != nil{
			ctx.AbortWithError(404, err)
		}
		ctx.JSON(204, nil)
	}
}

func (h *WeaponXCharacterDataHandler) HandlerDeleteByCharacterDataId() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil{
			ctx.AbortWithError(400, err)
            return
		}
		err = h.service.Delete(int64(id))
		if err != nil{
			ctx.AbortWithError(404, err)
		}
		ctx.JSON(204, nil)
	}
}

func (h *WeaponXCharacterDataHandler) HandlerGetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		weaponXCharacterDataList, err := h.service.GetAll()
		if err != nil {
			ctx.AbortWithError(500, err)
		}
		ctx.JSON(200, weaponXCharacterDataList)
	}
}

func (h *WeaponXCharacterDataHandler) HandlerGetByCharacterDataId() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err!= nil{
            ctx.AbortWithError(400, err)
            return
        }
		weaponXCharacterDataList, err := h.service.GetByCharacterDataId(int64(id))
		if err != nil {
			ctx.AbortWithError(500, err)
		}
		ctx.JSON(200, weaponXCharacterDataList)
	}
}

func (h *WeaponXCharacterDataHandler) HandlerGetById() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err!= nil{
            ctx.AbortWithError(400, err)
            return
        }
		weaponXCharacterData, err := h.service.GetByCharacterDataId(int64(id))
		if err != nil {
			ctx.AbortWithError(500, err)
		}
		ctx.JSON(200, weaponXCharacterData)
	}
}

func (h * WeaponXCharacterDataHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeaponXCharacterData domain.WeaponXCharacterData
		if err := ctx.BindJSON(&tempWeaponXCharacterData); err != nil {
			ctx.AbortWithError(400, err)
            return
		}
		updatedWeaponXCharacterData, err := h.service.Create(tempWeaponXCharacterData)
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.JSON(201, updatedWeaponXCharacterData)
	}
}