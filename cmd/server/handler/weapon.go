package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/weapon"
)

type WeaponHandler struct {
	service weapon.ServiceWeapon
}

func NewWeaponHandler(service *weapon.ServiceWeapon) WeaponHandler{
	return WeaponHandler{service: *service}
}

func (h *WeaponHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeapon domain.Weapon
		if err := ctx.BindJSON(&tempWeapon); err != nil {
			ctx.AbortWithError(400, err)
            return
		}

		createdWeapon, err := h.service.Create(tempWeapon)
		if err!= nil {
            ctx.AbortWithError(500, err)
            return
        }
		ctx.JSON(201, createdWeapon)
	}
}

func (h *WeaponHandler) HandlerDelete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.AbortWithError(400, err)
			return
		}
		err = h.service.Delete(int64(id))
		if err != nil {
			ctx.AbortWithError(404, err)
			return
		}
		ctx.JSON(204, nil)
	}
}

func (h *WeaponHandler) HandlerGetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		weapons, err := h.service.GetAll()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, weapons)
	}
}

func (h *WeaponHandler) HandlerGetByCampaignId() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}
		weapons, err := h.  service.GetByCampaignId(int64(id))
		if err != nil{
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, weapons)
	}
}

func (h *WeaponHandler) HandlerGetById() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			// We should change unsuccessful responses to abortwith status or abort with status json
			ctx.AbortWithError(400, err)
			return
		}
		weapon, err := h.  service.GetById(int64(id))
		if err != nil{
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(200, weapon)
	}
}


func (h *WeaponHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeapon domain.Weapon
		if err := ctx.BindJSON(&tempWeapon); err != nil {
			ctx.AbortWithError(400, err)
            return
		}

		updatedWeapon, err := h.service.Create(tempWeapon)
		if err!= nil {
            ctx.AbortWithError(500, err)
            return
        }
		ctx.JSON(201, updatedWeapon)
	}
}