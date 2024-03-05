package user_campaign

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	userCampaignRepository UserCampaignRepository
}

func NewUserCampaignService(userCampaignRepository UserCampaignRepository) UserCampaignService {
	return &service{userCampaignRepository: userCampaignRepository}
}

func (s *service) CreateUserCampaign(userCampaignDto dto.CreateUserCampaignDto) (domain.UserCampaign, error) {
	userCampaignDomain := domain.UserCampaign{
		CampaignId: userCampaignDto.CampaignId,
		UserId:     userCampaignDto.UserId,
		IsDm:       userCampaignDto.IsDm,
		IsOwner:    userCampaignDto.IsOwner,
	}

	createdUserCampaign, err := s.userCampaignRepository.Create(userCampaignDomain)
	if err != nil {
		return domain.UserCampaign{}, err
	}

	return createdUserCampaign, nil
}

func (s *service) GetAllUserCampaigns() ([]domain.UserCampaign, error) {
	userCampaigns, err := s.userCampaignRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return userCampaigns, nil
}

func (s *service) GetUserCampaignById(id int) (domain.UserCampaign, error) {
	userCampaign, err := s.userCampaignRepository.GetById(id)
	if err != nil {
		return domain.UserCampaign{}, err
	}

	return userCampaign, nil
}

func (s *service) GetUserCampaignByCampaignId(id int) ([]domain.UserCampaign, error) {
	userCampaigns, err := s.userCampaignRepository.GetByCampaignId(id)
	if err != nil {
		return nil, err
	}

	return userCampaigns, nil
}

func (s *service) GetUserCampaignByUserId(id string) ([]domain.UserCampaign, error) {
	userCampaigns, err := s.userCampaignRepository.GetByUserId(id)
	if err != nil {
		return nil, err
	}

	return userCampaigns, nil
}

func (s *service) DeleteUserCampaign(id int) error {
	err := s.userCampaignRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
