package event_type

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type eventTypeMySqlRepository struct {
	db *sql.DB
}

func NewEventTypeRepository(db *sql.DB) EventTypeRepository {
	return &eventTypeMySqlRepository{db: db}
}

func (r *eventTypeMySqlRepository) Create(eventType domain.EventType) (domain.EventType, error) {
	statement, err := r.db.Prepare(QueryCreateEventType)
	if err != nil {
		return domain.EventType{}, ErrPrepareStatement
	}

	defer statement.Close()
	result, err := statement.Exec(
		eventType.Name,
	)
	if err != nil {
		return domain.EventType{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.EventType{}, ErrGettingLastInsertId
	}
	eventType.EventTypeId = int(lastId)

	return eventType, nil
}

func (r *eventTypeMySqlRepository) GetAll() ([]domain.EventType, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var eventTypes []domain.EventType
	for rows.Next() {
		var eventType domain.EventType
		err := rows.Scan(
			&eventType.EventTypeId,
			&eventType.Name,
		)
		if err != nil {
			return nil, err
		}
		eventTypes = append(eventTypes, eventType)
	}

	return eventTypes, nil
}

func (r *eventTypeMySqlRepository) GetById(id int) (domain.EventType, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var eventType domain.EventType
	err := row.Scan(
		&eventType.EventTypeId,
		&eventType.Name,
	)
	if err != nil {
		return domain.EventType{}, err
	}

	return eventType, nil
}

func (r *eventTypeMySqlRepository) GetByName(name string) (domain.EventType, error) {
	row := r.db.QueryRow(QueryGetByName, name)
	var eventType domain.EventType
	err := row.Scan(
		&eventType.EventTypeId,
		&eventType.Name,
	)
	if err != nil {
		return domain.EventType{}, err
	}

	return eventType, nil
}

func (r *eventTypeMySqlRepository) Update(eventType domain.EventType, id int) (domain.EventType, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.EventType{}, ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(
		eventType.Name,
		id,
	)
	if err != nil {
		return domain.EventType{}, err
	}

	eventType.EventTypeId = id

	return eventType, nil
}

func (r *eventTypeMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
