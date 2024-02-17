package event

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

type eventMySqlRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventMySqlRepository{db: db}
}

func (r *eventMySqlRepository) Create(event domain.Event) (domain.Event, error) {
	statement, err := r.db.Prepare(QueryCreateEvent)
	if err != nil {
		fmt.Println(err)
		return domain.Event{}, ErrPrepareStatement
	}
	
	defer statement.Close()
	result, err := statement.Exec(
		event.Type,
		event.EventDescription,
		event.Environment,
		event.Session_id,
		event.Character_involved,
		event.Dice_roll,
		event.Difficulty_Class,
	)
	if err != nil {
		return domain.Event{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Event{}, ErrGettingLastInsertId
	}
	event.EventId = int(lastId)
	
	return event, nil
}


func (r *eventMySqlRepository) GetAll() ([]domain.Event, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		if err := rows.Scan(&event.EventId, &event.Type, &event.EventDescription, &event.Environment, &event.Session_id, &event.Character_involved, &event.Dice_roll, &event.Difficulty_Class); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *eventMySqlRepository) GetById(id int) (domain.Event, error) {
	var event domain.Event
	err := r.db.QueryRow(QueryGetById, id).Scan(&event.EventId, &event.Type, &event.EventDescription, &event.Environment, &event.Session_id, &event.Character_involved, &event.Dice_roll, &event.Difficulty_Class)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Event{}, errors.New("event not found")
		}
		return domain.Event{}, err
	}
	return event, nil
}

func (r *eventMySqlRepository) GetBySessionId(session_id int) ([]domain.Event, error) {
	rows, err := r.db.Query(QueryGetBySessionId, session_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		if err := rows.Scan(&event.EventId, &event.Type, &event.EventDescription, &event.Environment, &event.Session_id, &event.Character_involved, &event.Dice_roll, &event.Difficulty_Class); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *eventMySqlRepository) GetByCharacterId(character_id int) ([]domain.Event, error) {
	rows, err := r.db.Query(QueryGetByCharacterId, character_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		if err := rows.Scan(&event.EventId, &event.Type, &event.EventDescription, &event.Environment, &event.Session_id, &event.Character_involved, &event.Dice_roll, &event.Difficulty_Class); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *eventMySqlRepository) Update(event domain.Event, id int) (domain.Event, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.Event{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(event.Type, event.EventDescription, event.Environment, event.Session_id, event.Character_involved, event.Dice_roll, event.Difficulty_Class, id)
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}

func (r *eventMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
