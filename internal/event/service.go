package event

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type service struct {
	repo EventRepository
}

func NewService(repo EventRepository) EventService {
	return &service{repo: repo}
}

func (s *service) CreateEvent(eventDto dto.EventDto) (dto.EventDto, error) {
	eventDomain := domain.Event{
		Idevent:            eventDto.Idevent,
		Type:               eventDto.Type,
		EventDescription:   eventDto.EventDescription,
		Enviroment:         eventDto.Enviroment,
		Session_id:         eventDto.Session_id,
		Character_involved: eventDto.Character_involved,
		Dice_roll:          eventDto.Dice_roll,
		Difficulty_Class:   eventDto.Difficulty_Class,
	}

	createdEvent, err := s.repo.Create(eventDomain)
	if err != nil {
		return dto.EventDto{}, err
	}

	return dto.EventDto{
		Idevent:            createdEvent.Idevent,
		Type:               createdEvent.Type,
		EventDescription:   createdEvent.EventDescription,
		Enviroment:         createdEvent.Enviroment,
		Session_id:         createdEvent.Session_id,
		Character_involved: createdEvent.Character_involved,
		Dice_roll:          createdEvent.Dice_roll,
		Difficulty_Class:   createdEvent.Difficulty_Class,
	}, nil
}

func (s *service) GetAllEvents() ([]dto.EventDto, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var eventsDto []dto.EventDto
	for _, event := range events {
		eventsDto = append(eventsDto, dto.EventDto{
			Idevent:            event.Idevent,
			Type:               event.Type,
			EventDescription:   event.EventDescription,
			Enviroment:         event.Enviroment,
			Session_id:         event.Session_id,
			Character_involved: event.Character_involved,
			Dice_roll:          event.Dice_roll,
			Difficulty_Class:   event.Difficulty_Class,
		})
	}

	return eventsDto, nil
}

func (s *service) GetEventByID(id int) (dto.EventDto, error) {
	event, err := s.repo.GetById(id)
	if err != nil {
		return dto.EventDto{}, err
	}

	return dto.EventDto{
		Idevent:            event.Idevent,
		Type:               event.Type,
		EventDescription:   event.EventDescription,
		Enviroment:         event.Enviroment,
		Session_id:         event.Session_id,
		Character_involved: event.Character_involved,
		Dice_roll:          event.Dice_roll,
		Difficulty_Class:   event.Difficulty_Class,
	}, nil
}

func (s *service) UpdateEvent(eventDto dto.EventDto) (dto.EventDto, error) {
	eventDomain := domain.Event{
		Idevent:            eventDto.Idevent,
		Type:               eventDto.Type,
		EventDescription:   eventDto.EventDescription,
		Enviroment:         eventDto.Enviroment,
		Session_id:         eventDto.Session_id,
		Character_involved: eventDto.Character_involved,
		Dice_roll:          eventDto.Dice_roll,
		Difficulty_Class:   eventDto.Difficulty_Class,
	}

	updatedEvent, err := s.repo.Update(eventDomain)
	if err != nil {
		return dto.EventDto{}, err
	}

	updatedEventDto := dto.EventDto{
		Idevent:            int(updatedEvent.Idevent),
		Type:               updatedEvent.Type,
		EventDescription:   updatedEvent.EventDescription,
		Enviroment:         updatedEvent.Enviroment,
		Session_id:         updatedEvent.Session_id,
		Character_involved: updatedEvent.Character_involved,
		Dice_roll:          updatedEvent.Dice_roll,
		Difficulty_Class:   updatedEvent.Difficulty_Class,
	}

	return updatedEventDto, nil
}

// DeleteEvent deletes an event.
func (s *service) DeleteEvent(id int) error {
	return s.repo.Delete(id)
}
