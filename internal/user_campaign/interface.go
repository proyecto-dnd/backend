package user_campaign

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type UserCampaignHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerGetByCampaignId() gin.HandlerFunc
	HandlerGetByUserId() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type UserCampaignService interface {
	CreateUserCampaign(dto.CreateUserCampaignDto) (domain.UserCampaign, error)
	GetAllUserCampaigns() ([]domain.UserCampaign, error)
	GetUserCampaignByID(id int) (domain.UserCampaign, error)
	GetUserCampaignByCampaignID(id int) ([]domain.UserCampaign, error)
	GetUserCampaignByUserId(id int) ([]domain.UserCampaign, error)
	UpdateUserCampaign(UserCampaign dto.CreateUserCampaignDto, id int) (domain.UserCampaign, error)
	DeleteUserCampaign(id int) error
}

type UserCampaignRepository interface {
	Create(UserCampaign domain.UserCampaign) (domain.UserCampaign, error)
	GetAll()([]domain.UserCampaign, error)
	GetById(id int)(domain.UserCampaign, error)
	GetByCampaignId(id int)([]domain.UserCampaign, error)
	GetByUserId(id int)([]domain.UserCampaign, error)
	Update(UserCampaign domain.UserCampaign, id int) (domain.UserCampaign, error)
	Delete(id int)error
}