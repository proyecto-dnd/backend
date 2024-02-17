package campaign

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repo CampaignRepository
}

func NewCampaignService(repo CampaignRepository) CampaignService {
	return &service{repo: repo}
}

func (s *service) CreateCampaign(campaignDto dto.CreateCampaignDto) (domain.Campaign, error) {
	campaignDomain := domain.Campaign{
		DungeonMaster: campaignDto.DungeonMaster,
		Name:          campaignDto.Name,
		Description:   campaignDto.Description,
		Image:         campaignDto.Image,
	}
	
	createdCampaign, err := s.repo.Create(campaignDomain)
	if err != nil {
		return domain.Campaign{}, err
	}

	return createdCampaign, nil
}

func (s *service) GetAllCampaigns() ([]domain.Campaign, error) {
	campaigns, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignByID(id int) (domain.Campaign, error) {
	campaign, err := s.repo.GetById(id)
	if err != nil {
		return domain.Campaign{}, err
	}

	return campaign, nil
}

func (s *service) UpdateCampaign(campaignDto dto.CreateCampaignDto, id int) (domain.Campaign, error) {
	campaignDomain := domain.Campaign{
		DungeonMaster: campaignDto.DungeonMaster,
		Name:          campaignDto.Name,
		Description:   campaignDto.Description,
		Image:         campaignDto.Image,
	}

	updatedCampaign, err := s.repo.Update(campaignDomain, id)
	if err != nil {
		return domain.Campaign{}, err
	}

	return updatedCampaign, nil
}

func (s *service) DeleteCampaign(id int) error {
	return s.repo.Delete(id)
}
