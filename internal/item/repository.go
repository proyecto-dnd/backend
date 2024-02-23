package item

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

type itemMySqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryItem.
func (r *itemMySqlRepository) Create(item domain.Item) (domain.Item, error) {
	statement, err := r.db.Prepare(QueryCreateItem)
	if err != nil {
		return domain.Item{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		item.Name,
		item.Weight,
		item.Price,
		item.Description,
		item.CampaignId,
	)

	if err != nil {
		return domain.Item{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Item{}, ErrGettingLastInsertId
	}

	item.ItemId = lastId

	return item, nil
}

// Delete implements RepositoryItem.
func (r *itemMySqlRepository) Delete(id int64) error {
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

// GetAll implements RepositoryItem.
func (r *itemMySqlRepository) GetAll() ([]domain.Item, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Item{}, err
	}
	defer rows.Close()
	var items []domain.Item

	for rows.Next() {
		var item domain.Item
		err := rows.Scan(
			&item.ItemId,
			&item.Name,
			&item.Weight,
			&item.Price,
			&item.Description,
			&item.CampaignId,
		)
		if err != nil {
			return []domain.Item{}, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return []domain.Item{}, err
	}
	return items, nil
}

// GetByCampaignId implements RepositoryItem.
func (r *itemMySqlRepository) GetByCampaignId(campaignId int64) ([]domain.Item, error) {
	rows, err := r.db.Query(QueryGetByCampaignId, campaignId)
	if err != nil {
		return []domain.Item{}, err
	}
	defer rows.Close()
	var items []domain.Item

	for rows.Next() {
		var item domain.Item
		err := rows.Scan(
			&item.ItemId,
			&item.Name,
			&item.Weight,
			&item.Price,
			&item.Description,
			&item.CampaignId,
		)
		if err != nil {
			return []domain.Item{}, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return []domain.Item{}, err
	}
	return items, nil
}

// GetById implements RepositoryItem.
func (r *itemMySqlRepository) GetById(id int64) (domain.Item, error) {
	row := r.db.QueryRow(QueryGetById, id)

	var item domain.Item
	err := row.Scan(
		&item.ItemId,
		&item.Name,
		&item.Weight,
		&item.Price,
		&item.Description,
		&item.CampaignId,
	)
	if err != nil {
		return domain.Item{}, ErrNotFound
	}

	return item, nil
}

// Update implements RepositoryItem.
func (r *itemMySqlRepository) Update(item domain.Item) (domain.Item, error) {
	statement, err := r.db.Prepare(QueryCreateItem)
	if err != nil {
		return domain.Item{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		item.Name,
		item.Weight,
		item.Price,
		item.Description,
		item.CampaignId,
		item.ItemId,
	)

	if err != nil {
		return domain.Item{}, err
	}

	return item, nil
}

func NewItemRepository(db *sql.DB) RepositoryItem {
	return &itemMySqlRepository{db}
}
