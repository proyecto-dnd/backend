package campaign

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type campaignMySqlRepository struct {
	db *sql.DB
}

func NewCampaignRepository(db *sql.DB) CampaignRepository {
	return &campaignMySqlRepository{db: db}
}

func (r *campaignMySqlRepository) Create(campaign domain.Campaign) (domain.Campaign, error) {
	statement, err := r.db.Prepare(QueryCreateCampaign)
	if err != nil {
		fmt.Println(err)
		return domain.Campaign{}, ErrPrepareStatement
	}
	
	defer statement.Close()
	result, err := statement.Exec(
		campaign.DungeonMaster,
		campaign.Name,
		campaign.Description,
		campaign.Image,
	)
	if err != nil {
		return domain.Campaign{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Campaign{}, ErrGettingLastInsertId
	}
	campaign.CampaignId = int(lastId)
	
	return campaign, nil
}


func (r *campaignMySqlRepository) GetAll() ([]domain.Campaign, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []domain.Campaign
	for rows.Next() {
		var campaign domain.Campaign
		if err := rows.Scan(&campaign.CampaignId, &campaign.DungeonMaster, &campaign.Name, &campaign.Description, &campaign.Image); err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}

func (r *campaignMySqlRepository) GetById(id int) (domain.Campaign, error) {
	var campaign domain.Campaign
	err := r.db.QueryRow(QueryGetById, id).Scan(&campaign.CampaignId, &campaign.DungeonMaster, &campaign.Name, &campaign.Description, &campaign.Image)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Campaign{}, errors.New("campaign not found")
		}
		return domain.Campaign{}, err
	}
	return campaign, nil
}

func (r *campaignMySqlRepository) GetCampaignsByUserId(id string) ([]domain.Campaign, error) {
	rows, err := r.db.Query(QueryGetByUserId, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var campaigns []domain.Campaign
	for rows.Next() {
		var campaign domain.Campaign
		if err := rows.Scan(&campaign.CampaignId, &campaign.DungeonMaster, &campaign.Name, &campaign.Description, &campaign.Image); err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}

func (r *campaignMySqlRepository) Update(campaign domain.Campaign, id int) (domain.Campaign, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.Campaign{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(campaign.DungeonMaster, campaign.Name, campaign.Description, campaign.Image, id)
	if err != nil {
		return domain.Campaign{}, err
	}

	campaign.CampaignId = id

	return campaign, nil
}

func (r *campaignMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
