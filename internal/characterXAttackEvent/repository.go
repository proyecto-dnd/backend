package characterxattackevent

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

type characterXAttackEventRepository struct {
	db *sql.DB
}

func NewCharacterXAttackEventRepository(db *sql.DB) CharacterXAttackEventRepository {
	return &characterXAttackEventRepository{db: db}
}

func (r *characterXAttackEventRepository) GetAll() ([]domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetAll)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	var characterXSpellEvents []domain.CharacterXAttackEvent
	for rows.Next() {
		var characterXSpellEvent domain.CharacterXAttackEvent
		err := rows.Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
		if err != nil {
			return nil, err
		}
		characterXSpellEvents = append(characterXSpellEvents, characterXSpellEvent)
	}
	return characterXSpellEvents, nil
}

func (r *characterXAttackEventRepository) GetById(id int) (domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetById)
	if err != nil {
		return domain.CharacterXAttackEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	var characterXSpellEvent domain.CharacterXAttackEvent
	err = statement.QueryRow(id).Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
	if err != nil {
		return domain.CharacterXAttackEvent{}, err
	}
	return characterXSpellEvent, nil
}

func (r *characterXAttackEventRepository) GetByCharacterId(characterId int) ([]domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetByCharacterId)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(characterId)
	if err != nil {
		return nil, err
	}

	var characterXSpellEvents []domain.CharacterXAttackEvent
	for rows.Next() {
		var characterXSpellEvent domain.CharacterXAttackEvent
		err := rows.Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
		if err != nil {
			return nil, err
		}
		characterXSpellEvents = append(characterXSpellEvents, characterXSpellEvent)
	}
	return characterXSpellEvents, nil
}

func (r *characterXAttackEventRepository) GetBySpellEventId(spellEventId int) ([]domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryGetBySpellEventId)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(spellEventId)
	if err != nil {
		return nil, err
	}

	var characterXSpellEvents []domain.CharacterXAttackEvent
	for rows.Next() {
		var characterXSpellEvent domain.CharacterXAttackEvent
		err := rows.Scan(&characterXSpellEvent.CharacterId, &characterXSpellEvent.SpellEventId)
		if err != nil {
			return nil, err
		}
		characterXSpellEvents = append(characterXSpellEvents, characterXSpellEvent)
	}
	return characterXSpellEvents, nil
}

func (r *characterXAttackEventRepository) Create(characterXSpellEvent domain.CharacterXAttackEvent) (domain.CharacterXAttackEvent, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.CharacterXAttackEvent{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		characterXSpellEvent.CharacterId,
		characterXSpellEvent.SpellEventId,
	)
	if err != nil {
		return domain.CharacterXAttackEvent{}, err
	}

	return characterXSpellEvent, nil
}

func (r *characterXAttackEventRepository) Delete(id int) error {
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