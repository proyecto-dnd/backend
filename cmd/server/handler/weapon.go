package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/weapon"
	"strconv"
)

type WeaponHandler struct {
	service weapon.WeaponService
}

func NewWeaponHandler(service *weapon.WeaponService) *WeaponHandler {
	return &WeaponHandler{service: *service}
}

func (h *WeaponHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempWeapon dto.CreateWeaponDto
		if err := ctx.BindJSON(&tempWeapon); err != nil {
			ctx.JSON(500, err)
			return
		}

		createdWeapon, err := h.service.CreateWeapon(tempWeapon)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(201, createdWeapon)
	}
}

func (h *WeaponHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		weaponList, err := h.service.GetAllWeapons()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, weaponList)
	}
}

func (h *WeaponHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		tempWeapon, err := h.service.GetWeaponById(intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, tempWeapon)
	}
}

func (h *WeaponHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		var tempWeapon dto.CreateWeaponDto
		if err := ctx.BindJSON(&tempWeapon); err != nil {
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}

		updatedWeapon, err := h.service.UpdateWeapon(tempWeapon, intId)
		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, updatedWeapon)
	}
}

func (h *WeaponHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		serviceErr := h.service.DeleteWeapon(intId)
		if serviceErr != nil {
			ctx.JSON(500, serviceErr)
			return
		}

		ctx.JSON(200, "Deleted Weapon with id "+id)
	}
}
