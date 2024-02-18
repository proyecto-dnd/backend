package item

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryItem interface {
	Create(item domain.Item) (domain.Item, error)
	GetAll() ([]domain.Item, error)
	GetByCampaignId(campaignId int64) ([]domain.Item, error)
	GetById(id int64) (domain.Item, error)
	Update(item domain.Item) (domain.Item, error)
	Delete(id int64) error
}

type ServiceItem interface {
	Create(item domain.Item) (domain.Item, error)
	GetAll() ([]domain.Item, error)
	GetByCampaignId(campaignId int64) ([]domain.Item, error)
	GetById(id int64) (domain.Item, error)
	Update(item domain.Item) (domain.Item, error)
	Delete(id int64) error
}