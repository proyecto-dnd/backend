package event

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("Error preparing statement")
	ErrGettingLastInsertId = errors.New("Error getting last insert id")
)

type eventMySqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryEvent.
func (r *eventMySqlRepository) Create(event domain.Event) (domain.Event, error) {
	statement, err := r.db.Prepare(QueryCreateEvent)
	if err != nil {
		return domain.Event{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		event.Type,
		event.EventDescription,
		event.Enviroment,
		event.Session_id,
		event.Character_id,
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
	event.Idevent = lastId
	return event, nil
}

// Delete implements RepositoryTableCharacter.
func (r *eventMySqlRepository) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements RepositoryTableCharacter.
func (r *eventMySqlRepository) GetAll() ([]domain.Event, error) {
	panic("unimplemented")
}

// GetByCharacterId implements RepositoryEvent.
func (r *eventMySqlRepository) GetByCharacterId(character_id int) ([]domain.Event, error) {
	panic("unimplemented")
}

// GetById implements RepositoryEvent.
func (r *eventMySqlRepository) GetById(id int) (domain.Event, error) {
	panic("unimplemented")
}

// GetBySessionId implements RepositoryEvent.
func (r *eventMySqlRepository) GetBySessionId(session_id int) ([]domain.Event, error) {
	panic("unimplemented")
}

// Update implements RepositoryEvent.
func (r *eventMySqlRepository) Update(event domain.Event) {
	panic("unimplemented")
}

func NewEventRepository(db *sql.DB) RepositoryEvent {
	return &eventMySqlRepository{db}
}
