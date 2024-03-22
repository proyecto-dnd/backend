package itemxcharacterdata

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("item not found")
)

type itemXCharacterDataSqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateItemXCharacterData)
	if err != nil {
		return domain.ItemXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		itemXCharacterData.CharacterData_Id,
		itemXCharacterData.Item.Item_Id,
		itemXCharacterData.Quantity,
	)

	if err != nil {
		return domain.ItemXCharacterData{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.ItemXCharacterData{}, ErrGettingLastInsertId
	}

	itemXCharacterData.Character_Item_Id = int(lastId)

	return itemXCharacterData, nil
}

// Delete implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) Delete(id int) error {
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

// DeleteByTableCharacterId implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) DeleteByCharacterDataId(id int) error {
	result, err := r.db.Exec(QueryDeleteByCharacterDataId, id)
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

// GetAll implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) GetAll() ([]domain.ItemXCharacterData, error) {
	rows, err := r.db.Query(QueryGetAll)

	if err != nil {
		return []domain.ItemXCharacterData{}, err
	}

	defer rows.Close()

	itemXCharacterDataList := []domain.ItemXCharacterData{}

	for rows.Next() {
		var itemXCharacterData domain.ItemXCharacterData
		err := rows.Scan(
			&itemXCharacterData.Character_Item_Id,
			&itemXCharacterData.CharacterData_Id,
			&itemXCharacterData.Item.Item_Id,
			&itemXCharacterData.Item.Name,
			&itemXCharacterData.Item.Weight,
			&itemXCharacterData.Item.Price,
			&itemXCharacterData.Item.Description,
			&itemXCharacterData.Item.Campaign_Id,
			&itemXCharacterData.Quantity,
		)
		if err != nil {
			return []domain.ItemXCharacterData{}, err
		}
		itemXCharacterDataList = append(itemXCharacterDataList, itemXCharacterData)
	}
	if err := rows.Err(); err != nil {
		return []domain.ItemXCharacterData{}, err
	}
	return itemXCharacterDataList, nil
}

// GetById implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) GetById(id int) (domain.ItemXCharacterData, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var itemXCharacterData domain.ItemXCharacterData
	err := row.Scan(
		&itemXCharacterData.Character_Item_Id,
		&itemXCharacterData.CharacterData_Id,
		&itemXCharacterData.Item.Item_Id,
		&itemXCharacterData.Item.Name,
		&itemXCharacterData.Item.Weight,
		&itemXCharacterData.Item.Price,
		&itemXCharacterData.Item.Description,
		&itemXCharacterData.Item.Campaign_Id,
		&itemXCharacterData.Quantity,
	)
	if err != nil {
		return domain.ItemXCharacterData{}, err
	}

	return itemXCharacterData, nil
}

// GetByTableCharacterId implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) GetByCharacterDataId(id int) ([]domain.ItemXCharacterData, error) {
	rows, err := r.db.Query(QueryGetByCharacterDataId, id)

	if err != nil {
		return []domain.ItemXCharacterData{}, err
	}

	defer rows.Close()

	itemXCharacterDataList := []domain.ItemXCharacterData{}

	for rows.Next() {
		var itemXCharacterData domain.ItemXCharacterData
		err := rows.Scan(
			&itemXCharacterData.Character_Item_Id,
			&itemXCharacterData.CharacterData_Id,
			&itemXCharacterData.Item.Item_Id,
			&itemXCharacterData.Item.Name,
			&itemXCharacterData.Item.Weight,
			&itemXCharacterData.Item.Price,
			&itemXCharacterData.Item.Description,
			&itemXCharacterData.Item.Campaign_Id,
			&itemXCharacterData.Quantity,
		)
		if err != nil {
			return []domain.ItemXCharacterData{}, err
		}
		itemXCharacterDataList = append(itemXCharacterDataList, itemXCharacterData)
	}
	if err := rows.Err(); err != nil {
		return []domain.ItemXCharacterData{}, err
	}
	return itemXCharacterDataList, nil
}

// Update implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.ItemXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		itemXCharacterData.CharacterData_Id,
		itemXCharacterData.Item.Item_Id,
		itemXCharacterData.Quantity,
		itemXCharacterData.Character_Item_Id,
	)

	if err != nil {
		return domain.ItemXCharacterData{}, err
	}

	return itemXCharacterData, nil
}

func (r *itemXCharacterDataSqlRepository) UpdateOwnership(itemXCharacterData domain.ItemXCharacterData) error {
	statement, err := r.db.Prepare(QueryUpdateOwnership)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		itemXCharacterData.CharacterData_Id,
		itemXCharacterData.Quantity,
		itemXCharacterData.Character_Item_Id,
	)

	if err != nil {
		return err
	}
	return nil
}

func NewItemXCharacterDataSqlRepository(db *sql.DB) RepositoryItemXCharacterData {
	return &itemXCharacterDataSqlRepository{db}
}
