package characterXspell

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

type characterXSpellRepository struct {
	db *sql.DB
}

func NewCharacterXSpellRepository(db *sql.DB) CharacterXSpellRepository {
	return &characterXSpellRepository{db: db}
}

func (r characterXSpellRepository) Create(characterXSpell domain.CharacterXSpell) (domain.CharacterXSpell, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.CharacterXSpell{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		characterXSpell.CharacterId,
		characterXSpell.SpellId,
	)
	if err != nil {
		return domain.CharacterXSpell{}, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.CharacterXSpell{}, ErrGettingLastInsertId
	}
	var createdCharacterXSpell = domain.CharacterXSpell{
		CharacterSpellId: int(lastId),
		CharacterId:      characterXSpell.CharacterId,
		SpellId:          characterXSpell.SpellId,
	}
	return createdCharacterXSpell, nil
}

func (r characterXSpellRepository) Delete(id int) error {
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
