package charactertrade

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrNotFound = errors.New("characters not found")
)

type CharacterTradeMySqlRepository struct {
	db *sql.DB
}

func NewCharacterTradeMySqlRepository(db *sql.DB) RepositoryCharacterTrade {
	return &CharacterTradeMySqlRepository{db}
}

// CreateCharacterTrade implements CharacterTradeRepository.
func (c *CharacterTradeMySqlRepository) BulkCreateCharacterTrade(characterTradeList []domain.CharacterTrade) error {
	values := []interface{}{}
	sqlQuery := QueryBulkInsert
	for _, characterTrade := range characterTradeList {
		sqlQuery += "(?, ?, ?, ?, ?, ?, ?, ?, ?),"
		values = append(values, characterTrade.TradeEvent_Id,
			characterTrade.Weapon,
			characterTrade.Item,
			characterTrade.Armor,
			characterTrade.ItemOwner,
			characterTrade.ItemReciever,
			characterTrade.Quantity,
			characterTrade.ItemName,
			characterTrade.ItemType)
	}
	sqlQuery = sqlQuery[:len(sqlQuery)-1]
	statement, err := c.db.Prepare(sqlQuery)
	if err != nil {
		return err
	}

	_, err = statement.Exec(values...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByTradeEventId implements CharacterTradeRepository.
func (c *CharacterTradeMySqlRepository) DeleteByTradeEventId(tradeEventId int) error {
	result, err := c.db.Exec(QueryDeleteByTradeEventId, tradeEventId)
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

// GetByTradeEventId implements CharacterTradeRepository.
func (c *CharacterTradeMySqlRepository) GetByTradeEventId(tradeEventId int) ([]domain.CharacterTrade, error) {
	rows, err := c.db.Query(QueryGetByTradeEventId, tradeEventId)
	if err != nil {
		return []domain.CharacterTrade{}, err
	}
	defer rows.Close()

	characterTrades := []domain.CharacterTrade{}
	for rows.Next() {
		var characterTrade domain.CharacterTrade
		err := rows.Scan(
			&characterTrade.CharacterTrade_Id,
			&characterTrade.TradeEvent_Id,
			&characterTrade.Weapon,
			&characterTrade.Item,
			&characterTrade.Armor,
			&characterTrade.ItemOwner,
			&characterTrade.ItemReciever,
			&characterTrade.Quantity,
			&characterTrade.ItemName,
			&characterTrade.ItemType,
		)
		if err != nil {
			return []domain.CharacterTrade{}, err
		}
		characterTrades = append(characterTrades, characterTrade)
	}
	if err := rows.Err(); err != nil {
		return []domain.CharacterTrade{}, err
	}
	return characterTrades, nil
}
