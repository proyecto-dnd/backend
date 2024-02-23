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

type itemXtableCharacterSqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryItemXTableCharacter.
func (r *itemXtableCharacterSqlRepository) Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateItemXTableCharacter)
	if err != nil {
		return domain.ItemXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		itemXCharacterData.Item.Item_Id,
		itemXCharacterData.CharacterData_Id,
		itemXCharacterData.Quantity,
	)

	if err != nil {
		return domain.ItemXCharacterData{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.ItemXCharacterData{}, ErrGettingLastInsertId
	}

	itemXCharacterData.ItemXCharacterData_Id = lastId

	return itemXCharacterData, nil
}

// Delete implements RepositoryItemXTableCharacter.
func (r *itemXtableCharacterSqlRepository) Delete(id int64) error {
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
func (r *itemXtableCharacterSqlRepository) DeleteByCharacterDataId(id int64) error {
	result, err := r.db.Exec(QueryDeleteByTableCharacterId, id)
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
func (r *itemXtableCharacterSqlRepository) GetAll() ([]domain.ItemXCharacterData, error) {
	rows, err := r.db.Query(QueryGetAll)

	if err != nil {
		return []domain.ItemXCharacterData{}, err
	}

	defer rows.Close()

	var itemXCharacterDataList []domain.ItemXCharacterData

	for rows.Next() {
		var itemXCharacterData domain.ItemXCharacterData
		err := rows.Scan(
			&itemXCharacterData.ItemXCharacterData_Id,
			&itemXCharacterData.Item.Item_Id,
			&itemXCharacterData.Item.Name,
			&itemXCharacterData.Item.Weight,
			&itemXCharacterData.Item.Price,
			&itemXCharacterData.Item.Description,
			&itemXCharacterData.Item.Campaign_Id,
			&itemXCharacterData.CharacterData_Id,
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
func (r *itemXtableCharacterSqlRepository) GetById(id int64) (domain.ItemXCharacterData, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var itemXCharacterData domain.ItemXCharacterData
	err := row.Scan(
		&itemXCharacterData.ItemXCharacterData_Id,
		&itemXCharacterData.Item.Item_Id,
		&itemXCharacterData.Item.Name,
		&itemXCharacterData.Item.Weight,
		&itemXCharacterData.Item.Price,
		&itemXCharacterData.Item.Description,
		&itemXCharacterData.Item.Campaign_Id,
		&itemXCharacterData.CharacterData_Id,
		&itemXCharacterData.Quantity,
	)
	if err != nil {
		return domain.ItemXCharacterData{}, err
	}

	return itemXCharacterData, nil
}

// GetByTableCharacterId implements RepositoryItemXTableCharacter.
func (r *itemXtableCharacterSqlRepository) GetByCharacterDataId(id int64) ([]domain.ItemXCharacterData, error) {
	rows, err := r.db.Query(QueryDeleteByTableCharacterId, id)

	if err != nil {
		return []domain.ItemXCharacterData{}, err
	}

	defer rows.Close()

	var itemXCharacterDataList []domain.ItemXCharacterData

	for rows.Next() {
		var itemXCharacterData domain.ItemXCharacterData
		err := rows.Scan(
			&itemXCharacterData.ItemXCharacterData_Id,
			&itemXCharacterData.Item.Item_Id,
			&itemXCharacterData.Item.Name,
			&itemXCharacterData.Item.Weight,
			&itemXCharacterData.Item.Price,
			&itemXCharacterData.Item.Description,
			&itemXCharacterData.Item.Campaign_Id,
			&itemXCharacterData.CharacterData_Id,
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
func (r *itemXtableCharacterSqlRepository) Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateItemXTableCharacter)
	if err != nil {
		return domain.ItemXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		itemXCharacterData.ItemXCharacterData_Id,
		itemXCharacterData.Item.Item_Id,
		itemXCharacterData.CharacterData_Id,
		itemXCharacterData.Quantity,
	)

	if err != nil {
		return domain.ItemXCharacterData{}, err
	}

	return itemXCharacterData, nil
}

func NewItemXtableCharacterSqlRepository(db *sql.DB) RepositoryItemXCharacterData {
	return &itemXtableCharacterSqlRepository{db}
}
