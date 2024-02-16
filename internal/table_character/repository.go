package tablecharacter

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("Error preparing statement")
	ErrGettingLastInsertId = errors.New("Error getting last insert id")
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
		character.User_id, 
		character.Name, 
		character.Class, 
		character.Race, 
		character.Background,
		character.Hitpoints,
		character.Speed,
		character.Armor_class,
		character.Level,
		character.Exp,
		character.Campaign_id,
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
	panic("unimplemented")
}

// GetAll implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetAll() ([]domain.TableCharacter, error) {
	panic("unimplemented")
}

// GetByCampaignId implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetByCampaignId(campaignid int) ([]domain.TableCharacter, error) {
	panic("unimplemented")
}

// GetById implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetById(id int) (domain.TableCharacter, error) {
	panic("unimplemented")
}

// GetByUserId implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetByUserId(userid int) ([]domain.TableCharacter, error) {
	panic("unimplemented")
}

// GetByUserIdAndCampaignId implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) GetByUserIdAndCampaignId(userid int, campaignid int) ([]domain.TableCharacter, error) {
	panic("unimplemented")
}

// Update implements RepositoryTableCharacter.
func (r *tableCharacterMySqlRepository) Update(character domain.TableCharacter) {
	panic("unimplemented")
}

func NewTableCharacterRepository(db *sql.DB) RepositoryTableCharacter {
	return &tableCharacterMySqlRepository{db}
}
