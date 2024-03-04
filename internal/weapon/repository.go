package weapon

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type weaponMySqlRepository struct {
	db *sql.DB
}

func NewWeaponRepository(db *sql.DB) WeaponRepository {
	return &weaponMySqlRepository{db: db}
}

func (r *weaponMySqlRepository) Create(weapon domain.Weapon) (domain.Weapon, error) {
	statement, err := r.db.Prepare(QueryCreateWeapon)
	if err != nil {
		fmt.Println(err)
		return domain.Weapon{}, ErrPrepareStatement
	}

	defer statement.Close()
	result, err := statement.Exec(
		weapon.WeaponType,
		weapon.Name,
		weapon.Weight,
		weapon.Price,
		weapon.Category,
		weapon.Reach,
		weapon.Description,
		weapon.Damage,
		weapon.VersatileDamage,
		weapon.Ammunition,
		weapon.DamageType,
		weapon.Basic,
	)
	if err != nil {
		return domain.Weapon{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Weapon{}, ErrGettingLastInsertId
	}
	weapon.WeaponID = int(lastId)

	return weapon, nil
}

func (r *weaponMySqlRepository) GetAllWeapons() ([]domain.Weapon, error) {
	rows, err := r.db.Query(QueryGetAllWeapons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weapons []domain.Weapon
	for rows.Next() {
		var weapon domain.Weapon
		if err := rows.Scan(
			&weapon.WeaponID,
			&weapon.WeaponType,
			&weapon.Name,
			&weapon.Weight,
			&weapon.Price,
			&weapon.Category,
			&weapon.Reach,
			&weapon.Description,
			&weapon.Damage,
			&weapon.VersatileDamage,
			&weapon.Ammunition,
			&weapon.DamageType,
			&weapon.Basic,
		); err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}
	return weapons, nil
}

func (r *weaponMySqlRepository) GetWeaponById(id int) (domain.Weapon, error) {
	var weapon domain.Weapon
	err := r.db.QueryRow(QueryGetWeaponById, id).Scan(
		&weapon.WeaponID,
		&weapon.WeaponType,
		&weapon.Name,
		&weapon.Weight,
		&weapon.Price,
		&weapon.Category,
		&weapon.Reach,
		&weapon.Description,
		&weapon.Damage,
		&weapon.VersatileDamage,
		&weapon.Ammunition,
		&weapon.DamageType,
		&weapon.Basic,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Weapon{}, errors.New("weapon not found")
		}
		return domain.Weapon{}, err
	}
	return weapon, nil
}

func (r *weaponMySqlRepository) UpdateWeapon(weapon domain.Weapon, id int) (domain.Weapon, error) {
	statement, err := r.db.Prepare(QueryUpdateWeapon)
	if err != nil {
		return domain.Weapon{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		weapon.WeaponType,
		weapon.Name,
		weapon.Weight,
		weapon.Price,
		weapon.Category,
		weapon.Reach,
		weapon.Description,
		weapon.Damage,
		weapon.VersatileDamage,
		weapon.Ammunition,
		weapon.DamageType,
		weapon.Basic,
		id,
	)
	if err != nil {
		return domain.Weapon{}, err
	}

	weapon.WeaponID = id
	return weapon, nil
}

func (r *weaponMySqlRepository) DeleteWeapon(id int) error {
	statement, err := r.db.Prepare(QueryDeleteWeapon)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
