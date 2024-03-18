package characterxspellevent

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("characters not found")
)

type characterXSpellEventRepository struct {
	db *sql.DB
}

func NewCharacterXSpellEventRepository(db *sql.DB) CharacterXSpellEventRepository {
	return &characterXSpellEventRepository{db: db}
}

func (r *characterXSpellEventRepository) GetAll() ([]domain.CharacterXSpellEvent, error) {
	statement, err := r.db.Prepare(QueryGetAll)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	var characterXSpellEvents []domain.CharacterXSpellEvent
	for rows.Next() {
		var characterXSpellEvent domain.CharacterXSpellEvent
		err := rows.Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
		if err != nil {
			return nil, err
		}
		characterXSpellEvents = append(characterXSpellEvents, characterXSpellEvent)
	}
	return characterXSpellEvents, nil
}

func (r *characterXSpellEventRepository) GetById(id int) (domain.CharacterXSpellEvent, error) {
	statement, err := r.db.Prepare(QueryGetById)
	if err != nil {
		return domain.CharacterXSpellEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	var characterXSpellEvent domain.CharacterXSpellEvent
	err = statement.QueryRow(id).Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
	if err != nil {
		return domain.CharacterXSpellEvent{}, err
	}
	return characterXSpellEvent, nil
}

func (r *characterXSpellEventRepository) GetByCharacterId(characterId int) ([]domain.CharacterXSpellEvent, error) {
	statement, err := r.db.Prepare(QueryGetByCharacterId)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(characterId)
	if err != nil {
		return nil, err
	}

	var characterXSpellEvents []domain.CharacterXSpellEvent
	for rows.Next() {
		var characterXSpellEvent domain.CharacterXSpellEvent
		err := rows.Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
		if err != nil {
			return nil, err
		}
		characterXSpellEvents = append(characterXSpellEvents, characterXSpellEvent)
	}
	return characterXSpellEvents, nil
}

func (r *characterXSpellEventRepository) GetBySpellEventId(spellEventId int) ([]domain.CharacterXSpellEvent, error) {
	statement, err := r.db.Prepare(QueryGetBySpellEventId)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(spellEventId)
	if err != nil {
		return nil, err
	}

	var characterXSpellEvents []domain.CharacterXSpellEvent
	for rows.Next() {
		var characterXSpellEvent domain.CharacterXSpellEvent
		err := rows.Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
		if err != nil {
			return nil, err
		}
		characterXSpellEvents = append(characterXSpellEvents, characterXSpellEvent)
	}
	return characterXSpellEvents, nil
}

func (r *characterXSpellEventRepository) Create(characterXSpellEvent domain.CharacterXSpellEvent) (domain.CharacterXSpellEvent, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.CharacterXSpellEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		characterXSpellEvent.CharacterId,
		characterXSpellEvent.SpellEventId,
	)
	if err != nil {
		return domain.CharacterXSpellEvent{}, err
	}

	return characterXSpellEvent, nil
}

func (r *characterXSpellEventRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDelete, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}