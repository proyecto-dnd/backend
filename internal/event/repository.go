package event

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
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
		return domain.Event{}, ErrPrepareStatement
	}
	
	defer statement.Close()
	result, err := statement.Exec(
		event.Type,
		event.Environment,
		event.Session_id,
		event.EventProtagonistId,
		event.Dice_rolled,
		event.Difficulty_Class,
		event.EventTarget,
		event.EventResolution,
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


func (r *eventMySqlRepository) GetAll() ([]dto.EventRepositoryResponseDto, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.EventRepositoryResponseDto
	for rows.Next() {
		var event dto.EventRepositoryResponseDto
		if err := rows.Scan(&event.EventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.Dice_rolled, &event.Difficulty_Class, &event.EventTarget, &event.EventResolution, &event.TypeName); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *eventMySqlRepository) GetById(id int) (dto.EventRepositoryResponseDto, error) {
	var event dto.EventRepositoryResponseDto
	err := r.db.QueryRow(QueryGetById, id).Scan(&event.EventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.Dice_rolled, &event.Difficulty_Class, &event.EventTarget, &event.EventResolution, &event.TypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.EventRepositoryResponseDto{}, errors.New("event not found")
		}
		return dto.EventRepositoryResponseDto{}, err
	}
	return event, nil
}

func (r *eventMySqlRepository) GetByTypeId(type_id int) ([]dto.EventRepositoryResponseDto, error) {
	rows, err := r.db.Query(QueryGetByTypeId, type_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.EventRepositoryResponseDto
	for rows.Next() {
		var event dto.EventRepositoryResponseDto
		if err := rows.Scan(&event.EventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.Dice_rolled, &event.Difficulty_Class, &event.EventTarget, &event.EventResolution, &event.TypeName); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *eventMySqlRepository) GetBySessionId(session_id int) ([]dto.EventRepositoryResponseDto, error) {
	rows, err := r.db.Query(QueryGetBySessionId, session_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.EventRepositoryResponseDto
	for rows.Next() {
		var event dto.EventRepositoryResponseDto
		if err := rows.Scan(&event.EventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.Dice_rolled, &event.Difficulty_Class, &event.EventTarget, &event.EventResolution, &event.TypeName); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (r *eventMySqlRepository) GetByProtagonistId(protagonist_id int) ([]dto.EventRepositoryResponseDto, error) {
	rows, err := r.db.Query(QueryGetByProtagonistId, protagonist_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []dto.EventRepositoryResponseDto
	for rows.Next() {
		var event dto.EventRepositoryResponseDto
		if err := rows.Scan(&event.EventId, &event.Type, &event.Environment, &event.Session_id, &event.EventProtagonistId, &event.Dice_rolled, &event.Difficulty_Class, &event.EventTarget, &event.EventResolution, &event.TypeName); err != nil {
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

	_, err = statement.Exec(event.Type, event.Environment, event.Session_id, event.EventProtagonistId, event.Dice_rolled, event.Difficulty_Class, event.EventTarget, event.EventResolution, id)
	if err != nil {
		return domain.Event{}, err
	}

	event.EventId = id

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
