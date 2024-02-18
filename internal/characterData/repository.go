package characterdata

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

type CharacterDataMySqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) Create(character domain.CharacterData) (domain.CharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateCharacter)
	if err != nil {
		return domain.CharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		character.UserId,
		character.Name,
		character.Class,
		character.Race,
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
		return domain.CharacterData{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.CharacterData{}, ErrGettingLastInsertId
	}
	character.CharacterId = lastId
	return character, nil
}

// Delete implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) Delete(id int) error {
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

// GetAll implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetAll() ([]domain.CharacterData, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.CharacterData{}, nil
	}
	defer rows.Close()

	var characters []domain.CharacterData


	for rows.Next() {
		var character domain.CharacterData
	
		err := rows.Scan(
			&character.CharacterId,
			&character.UserId,
			&character.Name,
			&character.Class,
			&character.Race,
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
			return []domain.CharacterData{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.CharacterData{}, err
	}

	return characters, nil
}

// GetByCampaignId implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetByCampaignId(campaignid int) ([]domain.CharacterData, error) {
	rows, err := r.db.Query(QueryGetByCampaignId, campaignid)
	if err != nil {
		return []domain.CharacterData{}, nil
	}
	defer rows.Close()

	var characters []domain.CharacterData


	for rows.Next() {
		var character domain.CharacterData
	
		err := rows.Scan(
			&character.CharacterId,
			&character.UserId,
			&character.Name,
			&character.Class,
			&character.Race,
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
			return []domain.CharacterData{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.CharacterData{}, err
	}

	return characters, nil
}

// GetById implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetById(id int) (domain.CharacterData, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var character domain.CharacterData

	err := row.Scan(
		&character.CharacterId,
		&character.UserId,
		&character.Name,
		&character.Class,
		&character.Race,
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
		return domain.CharacterData{}, ErrNotFound
	}

	return character, nil
}

// GetByUserId implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetByUserId(userid string) ([]domain.CharacterData, error) {
	rows, err := r.db.Query(QueryGetByUserId, userid)
	if err != nil {
		return []domain.CharacterData{}, nil
	}
	defer rows.Close()

	var characters []domain.CharacterData


	for rows.Next() {
		var character domain.CharacterData
	
		err := rows.Scan(
			&character.CharacterId,
			&character.UserId,
			&character.Name,
			&character.Class,
			&character.Race,
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
			return []domain.CharacterData{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.CharacterData{}, err
	}

	return characters, nil
}

// GetByUserIdAndCampaignId implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetByUserIdAndCampaignId(userid string, campaignid int) ([]domain.CharacterData, error) {
	rows, err := r.db.Query(QueryGetByUserIdAndCampaignId, userid, campaignid)
	if err != nil {
		return []domain.CharacterData{}, nil
	}
	defer rows.Close()

	var characters []domain.CharacterData


	for rows.Next() {
		var character domain.CharacterData
	
		err := rows.Scan(
			&character.CharacterId,
			&character.UserId,
			&character.Name,
			&character.Class,
			&character.Race,
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
			return []domain.CharacterData{}, err
		}
		characters = append(characters, character)
	}
	if err := rows.Err(); err != nil {
		return []domain.CharacterData{}, err
	}

	return characters, nil
}

// Update implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) Update(character domain.CharacterData) (domain.CharacterData, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.CharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(
		character.UserId,
		character.Name,
		character.Class,
		character.Race,
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
		character.CharacterId,
	)
	if err != nil {
		return domain.CharacterData{}, err
	}
	return character, nil
}

func NewCharacterDataRepository(db *sql.DB) RepositoryCharacterData {
	return &CharacterDataMySqlRepository{db}
}
