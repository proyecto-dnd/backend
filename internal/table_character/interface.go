package tablecharacter

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type RepositoryTableCharacter interface {
	Create(character domain.TableCharacter) (domain.TableCharacter, error)
	GetAll()([]domain.TableCharacter, error)
	GetById(id int)(domain.TableCharacter, error)
	GetByUserId(userid int)([]domain.TableCharacter, error)
	GetByUserIdAndCampaignId(userid, campaignid int)([]domain.TableCharacter, error)
	GetByCampaignId(campaignid int)([]domain.TableCharacter, error)
	Update(character domain.TableCharacter) 
	Delete(id int)error
}