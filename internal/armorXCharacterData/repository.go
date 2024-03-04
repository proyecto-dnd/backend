package armorXCharacterData

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatementArmor    = errors.New("error preparing statement for armor")
	ErrGettingLastInsertIdArmor = errors.New("error getting last insert id for armor")
	ErrNotFoundArmor            = errors.New("armor not found")
)

type armorXCharacterDataSqlRepository struct {
	db *sql.DB
}

func (r *armorXCharacterDataSqlRepository) CreateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateCharacterArmor)
	if err != nil {
		return domain.ArmorXCharacterData{}, ErrPrepareStatementArmor
	}
	defer statement.Close()

	result, err := statement.Exec(
		data.CharacterData_Id,
		data.Armor_Id,
		data.Equipped,
	)

	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.ArmorXCharacterData{}, ErrGettingLastInsertIdArmor
	}

	data.ArmorXCharacterData_Id = lastId

	return data, nil
}

// DeleteArmorXCharacterData implements RepositoryArmorXCharacterData.
func (r *armorXCharacterDataSqlRepository) DeleteArmorXCharacterData(id int64) error {
	result, err := r.db.Exec(QueryDeleteCharacterArmor, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFoundArmor
	}
	return nil
}

func (r *armorXCharacterDataSqlRepository) DeleteByCharacterDataIdArmor(id int64) error {
	result, err := r.db.Exec(QueryDeleteByCharacterIdCharacterArmor, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFoundArmor
	}
	return nil
}

func (r *armorXCharacterDataSqlRepository) GetAllArmorXCharacterData() ([]domain.ArmorXCharacterData, error) {
	rows, err := r.db.Query(QueryGetByIdCharacterArmor)

	if err != nil {
		return []domain.ArmorXCharacterData{}, err
	}

	defer rows.Close()

	var armorXCharacterDataList []domain.ArmorXCharacterData

	for rows.Next() {
		var data domain.ArmorXCharacterData
		err := rows.Scan(
			&data.ArmorXCharacterData_Id,
			&data.CharacterData_Id,
			&data.Armor_Id,
			&data.Equipped,
		)
		if err != nil {
			return []domain.ArmorXCharacterData{}, err
		}
		armorXCharacterDataList = append(armorXCharacterDataList, data)
	}
	if err := rows.Err(); err != nil {
		return []domain.ArmorXCharacterData{}, err
	}
	return armorXCharacterDataList, nil
}

func (r *armorXCharacterDataSqlRepository) GetByIdArmorXCharacterData(id int64) (domain.ArmorXCharacterData, error) {
	row := r.db.QueryRow(QueryGetByCharacterIdCharacterArmor, id)
	var data domain.ArmorXCharacterData
	err := row.Scan(
		&data.ArmorXCharacterData_Id,
		&data.CharacterData_Id,
		&data.Armor_Id,
		&data.Equipped,
	)
	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}

	return data, nil
}

func (r *armorXCharacterDataSqlRepository) GetByCharacterDataIdArmor(id int64) ([]domain.ArmorXCharacterData, error) {
	rows, err := r.db.Query(QueryGetByCharacterIdCharacterArmor, id)

	if err != nil {
		return []domain.ArmorXCharacterData{}, err
	}

	defer rows.Close()

	var armorXCharacterDataList []domain.ArmorXCharacterData

	for rows.Next() {
		var data domain.ArmorXCharacterData
		err := rows.Scan(
			&data.ArmorXCharacterData_Id,
			&data.CharacterData_Id,
			&data.Armor_Id,
			&data.Equipped,
		)
		if err != nil {
			return []domain.ArmorXCharacterData{}, err
		}
		armorXCharacterDataList = append(armorXCharacterDataList, data)
	}
	if err := rows.Err(); err != nil {
		return []domain.ArmorXCharacterData{}, err
	}
	return armorXCharacterDataList, nil
}

func (r *armorXCharacterDataSqlRepository) UpdateArmorXCharacterData(data domain.ArmorXCharacterData) (domain.ArmorXCharacterData, error) {
	statement, err := r.db.Prepare(QueryUpdateCharacterArmor)
	if err != nil {
		return domain.ArmorXCharacterData{}, ErrPrepareStatementArmor
	}
	defer statement.Close()

	_, err = statement.Exec(
		data.CharacterData_Id,
		data.Armor_Id,
		data.Equipped,
		data.ArmorXCharacterData_Id,
	)

	if err != nil {
		return domain.ArmorXCharacterData{}, err
	}

	return data, nil
}

func NewArmorXCharacterDataSqlRepository(db *sql.DB) RepositoryArmorXCharacterData {
	return &armorXCharacterDataSqlRepository{db}
}
