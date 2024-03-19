package dice_event

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrLasInsertId      = errors.New("error getting last insert id")
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) DiceEventRepository {
	return &repository{db: db}
}

func (r *repository) Create(diceEvent domain.DiceEvent) (domain.DiceEvent, error) {
	statement, err := r.db.Prepare(QueryCreate)
	if err != nil {
		return domain.DiceEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		diceEvent.DiceEventId,
		diceEvent.Stat,
		diceEvent.Difficulty,
		diceEvent.DiceRolled,
		diceEvent.DiceResult,
		diceEvent.EventProtagonist,
		diceEvent.Description,
		diceEvent.SessionId,
		diceEvent.TimeStamp,
	)
	if err != nil {
		return domain.DiceEvent{}, err
	}
	return domain.DiceEvent{}, nil
}
