package tablecharacter

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("Error preparing statement")
	ErrGettingLastInsertId = errors.New("Error getting last insert id")
	ErrNotFound            = errors.New("characters not found")
)

type tableCharacterMySqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) Create(character domain.TableCharacter) (domain.TableCharacter, error) {
	statement, err := r.db.Prepare(QueryCreateCharacter)
	if err != nil {
		return domain.TableCharacter{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		character.UserId,
		character.Name,
		character.ClassId,
		character.RaceId,
		character.Background,
		character.Hitpoints,
		character.Speed,
		character.ArmorClass,
		character.Level,
		character.Exp,
		character.CampaignId,
		character.Str,
		character.Dex,
		character.Int,
		character.Wiz,
		character.Con,
		character.Cha,
	)
	if err != nil {
		return domain.TableCharacter{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.TableCharacter{}, ErrGettingLastInsertId
	}
	character.Idcharacter = lastId
	return character, nil
}

// Delete implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDelete, id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// GetAll implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetAll() ([]domain.TableCharacter, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.TableCharacter{}, nil
	}
	defer rows.Close()

	var characters []domain.TableCharacter

	for rows.Next() {
		var character domain.TableCharacter
		err := rows.Scan(
			&character.Idcharacter,
			&character.UserId,
			&character.Name,
			&character.ClassId,
			&character.RaceId,
			&character.Background,
			&character.Hitpoints,
			&character.Speed,
			&character.ArmorClass,
			&character.Level,
			&character.Exp,
			&character.CampaignId,
			&character.Str,
			&character.Dex,
			&character.Int,
			&character.Wiz,
			&character.Con,
			&character.Cha,
		)
		if err != nil {
			return []domain.TableCharacter{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.TableCharacter{}, err
	}

	return characters, nil
}

// GetByCampaignId implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetByCampaignId(campaignid int) ([]domain.TableCharacter, error) {
	rows, err := r.db.Query(QueryGetByCampaignId, campaignid)
	if err != nil {
		return []domain.TableCharacter{}, nil
	}
	defer rows.Close()

	var characters []domain.TableCharacter

	for rows.Next() {
		var character domain.TableCharacter
		err := rows.Scan(
			&character.Idcharacter,
			&character.UserId,
			&character.Name,
			&character.ClassId,
			&character.RaceId,
			&character.Background,
			&character.Hitpoints,
			&character.Speed,
			&character.ArmorClass,
			&character.Level,
			&character.Exp,
			&character.CampaignId,
			&character.Str,
			&character.Dex,
			&character.Int,
			&character.Wiz,
			&character.Con,
			&character.Cha,
		)
		if err != nil {
			return []domain.TableCharacter{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.TableCharacter{}, err
	}

	return characters, nil
}

// GetById implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetById(id int) (domain.TableCharacter, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var character domain.TableCharacter
	err := row.Scan(
		&character.Idcharacter,
		&character.UserId,
		&character.Name,
		&character.ClassId,
		&character.RaceId,
		&character.Background,
		&character.Hitpoints,
		&character.Speed,
		&character.ArmorClass,
		&character.Level,
		&character.Exp,
		&character.CampaignId,
		&character.Str,
		&character.Dex,
		&character.Int,
		&character.Wiz,
		&character.Con,
		&character.Cha,
	)
	if err != nil {
		return domain.TableCharacter{}, err
	}

	return character, nil
}

// GetByUserId implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetByUserId(userid string) ([]domain.TableCharacter, error) {
	rows, err := r.db.Query(QueryGetByUserId, userid)
	if err != nil {
		return []domain.TableCharacter{}, nil
	}
	defer rows.Close()

	var characters []domain.TableCharacter

	for rows.Next() {
		var character domain.TableCharacter
		err := rows.Scan(
			&character.Idcharacter,
			&character.UserId,
			&character.Name,
			&character.ClassId,
			&character.RaceId,
			&character.Background,
			&character.Hitpoints,
			&character.Speed,
			&character.ArmorClass,
			&character.Level,
			&character.Exp,
			&character.CampaignId,
			&character.Str,
			&character.Dex,
			&character.Int,
			&character.Wiz,
			&character.Con,
			&character.Cha,
		)
		if err != nil {
			return []domain.TableCharacter{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.TableCharacter{}, err
	}

	return characters, nil
}

// GetByUserIdAndCampaignId implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetByUserIdAndCampaignId(userid string, campaignid int) ([]domain.TableCharacter, error) {
	rows, err := r.db.Query(QueryGetByUserIdAndCampaignId, userid, campaignid)
	if err != nil {
		return []domain.TableCharacter{}, nil
	}
	defer rows.Close()

	var characters []domain.TableCharacter

	for rows.Next() {
		var character domain.TableCharacter
		err := rows.Scan(
			&character.Idcharacter,
			&character.UserId,
			&character.Name,
			&character.ClassId,
			&character.RaceId,
			&character.Background,
			&character.Hitpoints,
			&character.Speed,
			&character.ArmorClass,
			&character.Level,
			&character.Exp,
			&character.CampaignId,
			&character.Str,
			&character.Dex,
			&character.Int,
			&character.Wiz,
			&character.Con,
			&character.Cha,
		)
		if err != nil {
			return []domain.TableCharacter{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.TableCharacter{}, err
	}

	return characters, nil
}

// Update implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) Update(character domain.TableCharacter) (domain.TableCharacter, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.TableCharacter{}, ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(
		character.UserId,
		character.Name,
		character.ClassId,
		character.RaceId,
		character.Background,
		character.Hitpoints,
		character.Speed,
		character.ArmorClass,
		character.Level,
		character.Exp,
		character.CampaignId,
		character.Str,
		character.Dex,
		character.Int,
		character.Wiz,
		character.Con,
		character.Cha,
		character.Idcharacter,
	)
	if err != nil {
		return domain.TableCharacter{}, err
	}
	return character, nil
}

func NewTableCharacterRepository(db *sql.DB) RepositoryTableCharacter {
	return &tableCharacterMySqlRepository{db}
}
