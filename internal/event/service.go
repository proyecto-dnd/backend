package event

import (
	"github.com/proyecto-dnd/backend/internal/character"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

// TODO: Agregar session y character repository para verificar que los ids existan
type service struct {
	repo EventRepository
	charactersService character.New
}

func NewEventService(repo EventRepository) EventService {
	return &service{repo: repo}
}

func (s *service) CreateEvent(eventDto dto.CreateEventDto) (domain.Event, error) {
	eventDomain := domain.Event{
		Type:               eventDto.Type,
		Environment:        eventDto.Environment,
		Session_id:         eventDto.Session_id,
		EventProtagonistId: eventDto.EventProtagonistId,
		Dice_rolled:        eventDto.Dice_rolled,
		Difficulty_Class:   eventDto.Difficulty_Class,
		EventTarget:        eventDto.EventTarget,
		EventResolution:    eventDto.EventResolution,
	}

	createdEvent, err := s.repo.Create(eventDomain)
	if err != nil {
		return domain.Event{}, err
	}

	return createdEvent, nil
}

func (s *service) GetAllEvents() ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var eventsToReturn []dto.ResponseEventDto
	for _, event := range events {
		eventType := domain.EventType{EventTypeId: event.Type, Name: event.TypeName}
		eventToReturn := dto.ResponseEventDto{
			EventId:            event.EventId,
			Type:               eventType,
			Environment:        event.Environment,
			Session_id:         event.Session_id,
			EventProtagonistId: event.EventProtagonistId,
			Dice_rolled:        event.Dice_rolled,
			Difficulty_Class:   event.Difficulty_Class,
			EventTarget:        event.EventTarget,
			EventResolution:    event.EventResolution,
		}
		eventsToReturn = append(eventsToReturn, eventToReturn)
	}

	return eventsToReturn, nil
}

func (s *service) GetEventById(id int) (dto.ResponseEventDto, error) {
	event, err := s.repo.GetById(id)
	if err != nil {
		return dto.ResponseEventDto{}, err
	}

	eventType := domain.EventType{EventTypeId: event.Type, Name: event.TypeName}
	eventToReturn := dto.ResponseEventDto{
		EventId:            event.EventId,
		Type:               eventType,
		Environment:        event.Environment,
		Session_id:         event.Session_id,
		EventProtagonistId: event.EventProtagonistId,
		Dice_rolled:        event.Dice_rolled,
		Difficulty_Class:   event.Difficulty_Class,
		EventTarget:        event.EventTarget,
		EventResolution:    event.EventResolution,
	}

	return eventToReturn, nil
}

func (s *service) GetEventsByTypeId(typeid int) ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetByTypeId(typeid)
	if err != nil {
		return nil, err
	}

	var eventsToReturn []dto.ResponseEventDto
	for _, event := range events {
		eventType := domain.EventType{EventTypeId: event.Type, Name: event.TypeName}
		eventToReturn := dto.ResponseEventDto{
			EventId:            event.EventId,
			Type:               eventType,
			Environment:        event.Environment,
			Session_id:         event.Session_id,
			EventProtagonistId: event.EventProtagonistId,
			Dice_rolled:        event.Dice_rolled,
			Difficulty_Class:   event.Difficulty_Class,
			EventTarget:        event.EventTarget,
			EventResolution:    event.EventResolution,
		}
		eventsToReturn = append(eventsToReturn, eventToReturn)
	}

	return eventsToReturn, nil
}

func (s *service) GetEventsBySessionId(sessionid int) ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetBySessionId(sessionid)
	if err != nil {
		return nil, err
	}

	var eventsToReturn []dto.ResponseEventDto
	for _, event := range events {
		eventType := domain.EventType{EventTypeId: event.Type, Name: event.TypeName}
		eventToReturn := dto.ResponseEventDto{
			EventId:            event.EventId,
			Type:               eventType,
			Environment:        event.Environment,
			Session_id:         event.Session_id,
			EventProtagonistId: event.EventProtagonistId,
			Dice_rolled:        event.Dice_rolled,
			Difficulty_Class:   event.Difficulty_Class,
			EventTarget:        event.EventTarget,
			EventResolution:    event.EventResolution,
		}
		eventsToReturn = append(eventsToReturn, eventToReturn)
	}

	return eventsToReturn, nil
}

func (s *service) GetEventsByProtagonistId(protagonistid int) ([]dto.ResponseEventDto, error) {
	events, err := s.repo.GetByProtagonistId(protagonistid)
	if err != nil {
		return nil, err
	}

	var eventsToReturn []dto.ResponseEventDto
	for _, event := range events {
		eventType := domain.EventType{EventTypeId: event.Type, Name: event.TypeName}
		eventToReturn := dto.ResponseEventDto{
			EventId:            event.EventId,
			Type:               eventType,
			Environment:        event.Environment,
			Session_id:         event.Session_id,
			EventProtagonistId: event.EventProtagonistId,
			Dice_rolled:        event.Dice_rolled,
			Difficulty_Class:   event.Difficulty_Class,
			EventTarget:        event.EventTarget,
			EventResolution:    event.EventResolution,
		}
		eventsToReturn = append(eventsToReturn, eventToReturn)
	}

	return eventsToReturn, nil
}



func (s *service) UpdateEvent(eventDto dto.CreateEventDto, id int) (domain.Event, error) {
	eventDomain := domain.Event{
		Type:               eventDto.Type,
		Environment:        eventDto.Environment,
		Session_id:         eventDto.Session_id,
		EventProtagonistId: eventDto.EventProtagonistId,
		Dice_rolled:        eventDto.Dice_rolled,
		Difficulty_Class:   eventDto.Difficulty_Class,
		EventTarget:        eventDto.EventTarget,
		EventResolution:    eventDto.EventResolution,
	}

	updatedEvent, err := s.repo.Update(eventDomain, id)
	if err != nil {
		return domain.Event{}, err
	}

	return updatedEvent, nil
}

func (s *service) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}
