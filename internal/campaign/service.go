package campaign

import (
	"github.com/proyecto-dnd/backend/internal/session"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	campaignRepository CampaignRepository
	sessionService  session.SessionService
}

func NewCampaignService(campaignRepository CampaignRepository, sessionService session.SessionService) CampaignService {
	return &service{campaignRepository: campaignRepository, sessionService: sessionService}
}

func (s *service) CreateCampaign(campaignDto dto.CreateCampaignDto) (domain.Campaign, error) {
	campaignDomain := domain.Campaign{
		DungeonMaster: campaignDto.DungeonMaster,
		Name:          campaignDto.Name,
		Description:   campaignDto.Description,
		Image:         campaignDto.Image,
	}
	
	createdCampaign, err := s.campaignRepository.Create(campaignDomain)
	if err != nil {
		return domain.Campaign{}, err
	}

	return createdCampaign, nil
}

func (s *service) GetAllCampaigns() ([]domain.Campaign, error) {
	// campaigns, err := s.campaignRepository.GetAll()
	// if err != nil {
	// 	return nil, err
	// }

	// var responseCampaigns []dto.ResponseCampaignDto
	// for _, campaign := range campaigns {
	// 	sessions, err := s.sessionService.GetSessionsByCampaignId(campaign.CampaignId)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	responseCampaign := dto.ResponseCampaignDto{
	// 		DungeonMaster: campaign.DungeonMaster,
	// 		Name:          campaign.Name,
	// 		Description:   campaign.Description,
	// 		Image:         campaign.Image,
	// 		Sessions:      sessions,
	// 	}

	// 	responseCampaigns = append(responseCampaigns, responseCampaign)
	
	// }

	// return responseCampaigns, nil
	return s.campaignRepository.GetAll()
}

func (s *service) GetCampaignByID(id int) (dto.ResponseCampaignDto, error) {
	campaign, err := s.campaignRepository.GetById(id)
	if err != nil {
		return dto.ResponseCampaignDto{}, err
	}

	sessions, err := s.sessionService.GetSessionsByCampaignId(campaign.CampaignId)
	if err != nil {
		return dto.ResponseCampaignDto{}, err
	}

	responseCampaign := dto.ResponseCampaignDto{
		DungeonMaster: campaign.DungeonMaster,
		Name:          campaign.Name,
		Description:   campaign.Description,
		Image:         campaign.Image,
		Sessions:      sessions,
	}

	return responseCampaign, nil
}

func (s *service) GetCampaignsByUserId(id string) ([]dto.ResponseCampaignDto, error) {
	campaigns, err := s.campaignRepository.GetCampaignsByUserId(id)
	if err != nil {
		return nil, err
	}

	var responseCampaigns []dto.ResponseCampaignDto
	for _, campaign := range campaigns {
		sessions, err := s.sessionService.GetSessionsByCampaignId(campaign.CampaignId)
		if err != nil {
			return nil, err
		}

		responseCampaign := dto.ResponseCampaignDto{
			DungeonMaster: campaign.DungeonMaster,
			Name:          campaign.Name,
			Description:   campaign.Description,
			Image:         campaign.Image,
			Sessions:      sessions,
		}

		responseCampaigns = append(responseCampaigns, responseCampaign)
	}

	return responseCampaigns, nil
}
		

func (s *service) UpdateCampaign(campaignDto dto.CreateCampaignDto, id int) (dto.ResponseCampaignDto, error) {
	campaignDomain := domain.Campaign{
		DungeonMaster: campaignDto.DungeonMaster,
		Name:          campaignDto.Name,
		Description:   campaignDto.Description,
		Image:         campaignDto.Image,
	}

	updatedCampaign, err := s.campaignRepository.Update(campaignDomain, id)
	if err != nil {
		return dto.ResponseCampaignDto{}, err
	}

	sessions, err := s.sessionService.GetSessionsByCampaignId(updatedCampaign.CampaignId)
	if err != nil {
		return dto.ResponseCampaignDto{}, err
	}

	responseCampaign := dto.ResponseCampaignDto{
		DungeonMaster: updatedCampaign.DungeonMaster,
		Name:          updatedCampaign.Name,
		Description:   updatedCampaign.Description,
		Image:         updatedCampaign.Image,
		Sessions:      sessions,
	}

	return responseCampaign, nil
}

func (s *service) DeleteCampaign(id int) error {
	return s.campaignRepository.Delete(id)
}
