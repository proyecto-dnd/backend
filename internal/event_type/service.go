package event_type

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repo              EventTypeRepository
}

func NewEventTypeService(repo EventTypeRepository) EventTypeService {
	return &service{repo: repo}
}

func (s *service) CreateEventType(eventTypeDto dto.CreateEventTypeDto) (domain.EventType, error) {
	eventTypeDomain := domain.EventType{
		Name: eventTypeDto.Name,
	}

	createdEventType, err := s.repo.Create(eventTypeDomain)
	if err != nil {
		return domain.EventType{}, err
	}

	return createdEventType, nil
}

func (s *service) GetAllEventTypes() ([]domain.EventType, error) {
	eventTypes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return eventTypes, nil
}

func (s *service) GetEventTypeById(id int) (domain.EventType, error) {
	eventType, err := s.repo.GetById(id)
	if err != nil {
		return domain.EventType{}, err
	}

	return eventType, nil
}

func (s *service) GetEventTypeByName(name string) (domain.EventType, error) {
	eventType, err := s.repo.GetByName(name)
	if err != nil {
		return domain.EventType{}, err
	}

	return eventType, nil
}

func (s *service) UpdateEventType(eventType dto.CreateEventTypeDto, id int) (domain.EventType, error) {

	eventTypeDomain := domain.EventType{
		Name: eventType.Name,
	}

	updatedEventType, err := s.repo.Update(eventTypeDomain, id)
	if err != nil {
		return domain.EventType{}, err
	}

	return updatedEventType, nil
}

func (s *service) DeleteEventType(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}