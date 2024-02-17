package event

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repo EventRepository
}

func NewEventService(repo EventRepository) EventService {
	return &service{repo: repo}
}

func (s *service) CreateEvent(eventDto dto.EventDto) (domain.Event, error) {
	eventDomain := domain.Event{
		Type:               eventDto.Type,
		EventDescription:   eventDto.EventDescription,
		Environment:         eventDto.Environment,
		Session_id:         eventDto.Session_id,
		Character_involved: eventDto.Character_involved,
		Dice_roll:          eventDto.Dice_roll,
		Difficulty_Class:   eventDto.Difficulty_Class,
	}
	
	createdEvent, err := s.repo.Create(eventDomain)
	if err != nil {
		return domain.Event{}, err
	}

	return createdEvent, nil
}

func (s *service) GetAllEvents() ([]domain.Event, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var eventsArray []domain.Event
	for _, event := range events {
		eventsArray = append(eventsArray, domain.Event{
			Idevent:            event.Idevent,
			Type:               event.Type,
			EventDescription:   event.EventDescription,
			Environment:         event.Environment,
			Session_id:         event.Session_id,
			Character_involved: event.Character_involved,
			Dice_roll:          event.Dice_roll,
			Difficulty_Class:   event.Difficulty_Class,
		})
	}

	return eventsArray, nil
}

func (s *service) GetEventByID(id int) (domain.Event, error) {
	event, err := s.repo.GetById(id)
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}

func (s *service) GetEventsBySessionID(sessionid int) ([]domain.Event, error) {
	events, err := s.repo.GetBySessionId(sessionid)
	if err != nil {
		return nil, err
	}

	var eventsDto []domain.Event
	for _, event := range events {
		eventsDto = append(eventsDto, domain.Event{
			Idevent:            event.Idevent,
			Type:               event.Type,
			EventDescription:   event.EventDescription,
			Environment:         event.Environment,
			Session_id:         event.Session_id,
			Character_involved: event.Character_involved,
			Dice_roll:          event.Dice_roll,
			Difficulty_Class:   event.Difficulty_Class,
		})
	}

	return eventsDto, nil
}

func (s *service) GetEventsByCharacterID(characterid int) ([]domain.Event, error) {
	events, err := s.repo.GetByCharacterId(characterid)
	if err != nil {
		return nil, err
	}

	var eventsDto []domain.Event
	for _, event := range events {
		eventsDto = append(eventsDto, domain.Event{
			Idevent:            event.Idevent,
			Type:               event.Type,
			EventDescription:   event.EventDescription,
			Environment:         event.Environment,
			Session_id:         event.Session_id,
			Character_involved: event.Character_involved,
			Dice_roll:          event.Dice_roll,
			Difficulty_Class:   event.Difficulty_Class,
		})
	}

	return eventsDto, nil
}

func (s *service) UpdateEvent(eventDto dto.EventDto, id int) (domain.Event, error) {
	eventDomain := domain.Event{
		Type:               eventDto.Type,
		EventDescription:   eventDto.EventDescription,
		Environment:         eventDto.Environment,
		Session_id:         eventDto.Session_id,
		Character_involved: eventDto.Character_involved,
		Dice_roll:          eventDto.Dice_roll,
		Difficulty_Class:   eventDto.Difficulty_Class,
	}

	updatedEvent, err := s.repo.Update(eventDomain, id)
	if err != nil {
		return domain.Event{}, err
	}

	updatedEventDto := domain.Event{
		Idevent:            int(updatedEvent.Idevent),
		Type:               updatedEvent.Type,
		EventDescription:   updatedEvent.EventDescription,
		Environment:         updatedEvent.Environment,
		Session_id:         updatedEvent.Session_id,
		Character_involved: updatedEvent.Character_involved,
		Dice_roll:          updatedEvent.Dice_roll,
		Difficulty_Class:   updatedEvent.Difficulty_Class,
	}

	return updatedEventDto, nil
}

func (s *service) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}
