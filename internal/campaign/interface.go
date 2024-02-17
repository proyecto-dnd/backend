package campaign

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CampaignHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type CampaignService interface {
	CreateCampaign(dto.CreateCampaignDto) (domain.Campaign, error)
	GetAllCampaigns() ([]domain.Campaign, error)
	GetCampaignByID(id int) (domain.Campaign, error)
	UpdateCampaign(Campaign dto.CreateCampaignDto, id int) (domain.Campaign, error)
	DeleteCampaign(id int) error
}

type CampaignRepository interface {
	Create(Campaign domain.Campaign) (domain.Campaign, error)
	GetAll()([]domain.Campaign, error)
	GetById(id int)(domain.Campaign, error)
	Update(Campaign domain.Campaign, id int) (domain.Campaign, error)
	Delete(id int)error
}