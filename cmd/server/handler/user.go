package handler

import (
	"fmt"
	"log"
	"strconv"
	"time"

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

// user godoc
// @Summary Register user
// @Tags user
// @Accept json
// @Produce json
// @Param body body domain.User true "User"
// @Success 201 {object} domain.UserResponse
// @Failure 500 {object} error
// @Router /user/register [post]
func (h *UserHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userTemp domain.User
		if err := ctx.BindJSON(&userTemp); err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(400, err)
			return
		}

		createdUser, err := h.service.Create(userTemp)
		if err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(400, err)
			return
		}

		// TEMP SUCCESS RESPONSE
		ctx.JSON(201, createdUser)
	}
}

// user godoc
// @Summary Get all users
// @Tags user
// @Produce json
// @Success 200 {array} domain.UserResponse
// @Failure 500 {object} error
// @Router /user [get]
func (h *UserHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userList, err := h.service.GetAll()
		if err != nil {
			// TEMP ERROR RESPONSE
			fmt.Println(err)
			ctx.JSON(500, err)
			return
		}

		// _, err = h.service.TransferDataToSql(userList)
		// if err != nil {
		// 	// TEMP ERROR RESPONSE
		// 	fmt.Println(err)
		// 	ctx.JSON(500, err)
		// 	return
		// }

		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, userList)
	}
}

// user godoc
// @Summary Get user by id
// @Tags user
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.UserResponse
// @Failure 500 {object} error
// @Router /user/{id} [get]
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

// user godoc
// @Summary Update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body domain.User true "User"
// @Success 200 {object} domain.User
// @Failure 500 {object} error
// @Router /user/{id} [put]
func (h *UserHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var userTemp domain.UserUpdate
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

// user godoc
// @Summary Partially update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body domain.User true "User"
// @Success 200 {object} domain.User
// @Failure 500 {object} error
// @Router /user/{id} [patch]
func (h *UserHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var userTemp domain.UserUpdate
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

// user godoc
// @Summary Delete user
// @Tags user
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /user/{id} [delete]
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

// user godoc
// @Summary Login user
// @Tags user
// @Accept json
// @Produce json
// @Param body body domain.UserLoginInfo true "UserLoginInfo"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /user/login [post]
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
	}
}

func (h *UserHandler) HandlerGetJwtInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		// log.Println(cookie.Value)
		jwtClaimsInfo, err := h.service.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err)
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, jwtClaimsInfo)
	}
}

func (h *UserHandler) HandlerSubPremium() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		monthsParam := ctx.Param("months")

		months, err := strconv.Atoi(monthsParam)
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		err = h.service.SubscribeToPremium(cookie.Value, time.Now().AddDate(0, months, 0).String())
		if err != nil {
			ctx.JSON(400, err)
			return
		}

		ctx.JSON(200, "Subscribed to Premium")
	}
}
