package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/user"
)

type UserHandler struct {
	service user.ServiceUsers
}

func NewUserHandler(service *user.ServiceUsers) *UserHandler {
	return &UserHandler{service: *service}
}

func (h *UserHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userTemp domain.User
		if err := ctx.BindJSON(&userTemp); err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}

		createdUser, err := h.service.Create(userTemp)
		if err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}

		// TEMP SUCCESS RESPONSE
		ctx.JSON(201, createdUser)
	}
}

func (h *UserHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userList, err := h.service.GetAll()
		if err != nil {
			// TEMP ERROR RESPONSE
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, userList)
	}
}

func (h *UserHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		userTemp, err := h.service.GetById(id)
		if err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, userTemp)
	}
}

func (h *UserHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var userTemp domain.User
		if err := ctx.BindJSON(&userTemp); err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}
		updatedUser, err := h.service.Update(userTemp, id)
		if err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, updatedUser)
	}
}

func (h *UserHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var userTemp domain.User
		if err := ctx.BindJSON(&userTemp); err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}
		patchedUser, err := h.service.Patch(userTemp, id)
		if err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, patchedUser)
	}
}

func (h *UserHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		err := h.service.Delete(id)
		if err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err)
			return
		}

		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, "Deleted User with id "+id)
	}
}

func (h *UserHandler) HandlerLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempUserInfo domain.UserLoginInfo

		err := ctx.BindJSON(&tempUserInfo)
		if err != nil {
			// TEMP ERROR RESPONSE
			log.Println("BINDING JSON: " + err.Error())
			ctx.JSON(500, err)
			return
		}
		// log.Println(tempUserInfo)
		cookie, err := h.service.Login(tempUserInfo)
		if err != nil {
			// TEMP ERROR RESPONSE
			log.Println("LOGIN SERVICE: " + err.Error())
			ctx.JSON(500, err)
			return
		}
		ctx.SetCookie("Session", cookie, 3600, "/", "localhost", false, false)
		// log.Println(cookie)
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, "Setted Cookie")
		return
	}
}
