package session

import (
	"fmt"

	"github.com/proyecto-dnd/backend/internal/campaign"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	sessionRepo  SessionRepository
	campaignRepo campaign.CampaignRepository
}

func NewSessionService(sessionRepo SessionRepository, campaignRepo campaign.CampaignRepository) SessionService {
	return &service{sessionRepo: sessionRepo, campaignRepo: campaignRepo}
}

func (s *service) CreateSession(sessionDto dto.CreateSessionDto) (domain.Session, error) {

	_, err := s.campaignRepo.GetById(sessionDto.CampaignId)
	if err != nil {
		return domain.Session{}, err
	}

	sessionDomain := domain.Session{
		Start:       sessionDto.Start,
		End:         sessionDto.End,
		Description: sessionDto.Description,
		CampaignId:  sessionDto.CampaignId,
	}

	createdSession, err := s.sessionRepo.Create(sessionDomain)
	if err != nil {
		return domain.Session{}, err
	}

	return createdSession, nil
}

func (s *service) GetAllSessions() ([]domain.Session, error) {
	sessions, err := s.sessionRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *service) GetSessionByID(id int) (domain.Session, error) {
	session, err := s.sessionRepo.GetById(id)
	if err != nil {
		return domain.Session{}, err
	}

	return session, nil
}

func (s *service) GetSessionByCampaignID(id int) ([]domain.Session, error) {
	sessions, err := s.sessionRepo.GetByCampaignId(id)
	if err != nil {
		return nil, err
	}

	return sessions, nil

}

func (s *service) UpdateSession(sessionDto dto.CreateSessionDto, id int) (domain.Session, error) {
	_, err := s.campaignRepo.GetById(sessionDto.CampaignId)
	if err != nil {
		return domain.Session{}, err
	}

	sessionDomain := domain.Session{
		Start:       sessionDto.Start,
		End:         sessionDto.End,
		Description: sessionDto.Description,
		CampaignId:  sessionDto.CampaignId,
	}

	updatedSession, err := s.sessionRepo.Update(sessionDomain, id)
	if err != nil {
		fmt.Println(err)
		return domain.Session{}, err
	}

	return updatedSession, nil
}

func (s *service) DeleteSession(id int) error {
	return s.sessionRepo.Delete(id)
}
