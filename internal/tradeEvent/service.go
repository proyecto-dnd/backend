package tradeevent

import (
	"errors"
	"github.com/proyecto-dnd/backend/internal/armorXCharacterData"
	charactertrade "github.com/proyecto-dnd/backend/internal/characterTrade"
	"github.com/proyecto-dnd/backend/internal/domain"
	itemxcharacterdata "github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
)

var (
	ErrCannotBeNegative = errors.New("cannot be negative")
)

type serviceTradeEvent struct {
	tradeEventRepo          RepositoryTradeEvent
	characterTradeService   charactertrade.ServiceCharacterTrade
	weaponXCharacterService weaponxcharacterdata.ServiceWeaponXCharacterData
	armorXCharacterService  armorXCharacterData.ServiceArmorXCharacterData
	itemXCharacterService   itemxcharacterdata.ServiceItemXCharacterData
}

// Create implements ServiceTradeEvent.
//TO DO: Implement method in itemXCharacterData to search by characterData_Id and Item_Id, must also be implemented in itemXcharacterData's create to prevent duping of entries
func (s *serviceTradeEvent) Create(tradeEvent domain.TradeEvent) (domain.TradeEvent, error) {
	newTradeEvent, err := s.tradeEventRepo.Create(tradeEvent)
	if err != nil {
		return domain.TradeEvent{}, err
	}

	for i, tradingItems := range tradeEvent.TradingItems {
		tradeEvent.TradingItems[i].TradeEvent_Id = newTradeEvent.TradeEvent_Id
		tradeEvent.TradingItems[i].ItemReciever = tradeEvent.Receiver
		tradeEvent.TradingItems[i].ItemOwner = tradeEvent.Sender
		if tradingItems.Weapon != nil {
			err = s.weaponXCharacterService.UpdateOwnership(domain.WeaponXCharacterData{Character_Weapon_Id: *tradingItems.Weapon, CharacterData_Id: tradeEvent.Receiver, Weapon: domain.Weapon{}, Equipped: false})
			if err != nil {
				return domain.TradeEvent{}, err
			}
		}
		if tradingItems.Armor != nil {
			err = s.armorXCharacterService.UpdateOwnership(domain.ArmorXCharacterData{ArmorXCharacterData_Id: *tradingItems.Armor, Armor: domain.Armor{}, CharacterData_Id: tradeEvent.Receiver, Equipped: false})
			if err != nil {
				return domain.TradeEvent{}, err
			}
		}

		if tradingItems.Item != nil {
			itemXCharacterToUpdate, err := s.itemXCharacterService.GetById(*tradingItems.Item)
			if err != nil {
				return domain.TradeEvent{}, err
			}
			if itemXCharacterToUpdate.Quantity < *tradingItems.Quantity {
				return domain.TradeEvent{}, ErrCannotBeNegative
			}
			if itemXCharacterToUpdate.Quantity == *tradingItems.Quantity {
				err = s.itemXCharacterService.UpdateOwnership(domain.ItemXCharacterData{Character_Item_Id: *tradingItems.Item, CharacterData_Id: tradeEvent.Receiver, Item: domain.Item{Item_Id: itemXCharacterToUpdate.Item.Item_Id}, Quantity: itemXCharacterToUpdate.Quantity})
				if err != nil {
					return domain.TradeEvent{}, err
				}
			} else {
				_,err = s.itemXCharacterService.Update(domain.ItemXCharacterData{Character_Item_Id: *tradingItems.Item, CharacterData_Id: tradeEvent.Sender, Item: domain.Item{Item_Id: itemXCharacterToUpdate.Item.Item_Id}, Quantity: itemXCharacterToUpdate.Quantity - *tradingItems.Quantity})
				if err != nil {
					return domain.TradeEvent{}, err
				}
				_ ,err = s.itemXCharacterService.Create(domain.ItemXCharacterData{Character_Item_Id: *tradingItems.Item, CharacterData_Id: tradeEvent.Receiver, Item: domain.Item{Item_Id: itemXCharacterToUpdate.Item.Item_Id}, Quantity: *tradingItems.Quantity})
				if err != nil {
					return domain.TradeEvent{}, err
				}
			}

		}
	}
	err = s.characterTradeService.BulkCreateCharacterTrade(tradeEvent.TradingItems)
	if err != nil {
		return domain.TradeEvent{}, err
	}
	newTradeEvent.TradingItems, err = s.characterTradeService.GetByTradeEventId(newTradeEvent.TradeEvent_Id)
	if err != nil {
		return domain.TradeEvent{}, err
	}
	return newTradeEvent, nil
}

// Delete implements ServiceTradeEvent.
func (s *serviceTradeEvent) Delete(id int) error {
	s.characterTradeService.DeleteByTradeEventId(id)
	return s.tradeEventRepo.Delete(id)
}

// GetByReceiver implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetByReceiver(receiver int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetByReceiver(receiver)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

// GetBySender implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetBySender(sender int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetBySender(sender)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

// GetBySenderOrReciever implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetBySenderOrReciever(id int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetBySenderOrReciever(id)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

// GetBySessionId implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetBySessionId(sessionId int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetBySessionId(sessionId)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

func NewTradeEventService(tradeEventRepo RepositoryTradeEvent, characterTradeService charactertrade.ServiceCharacterTrade, weaponService weaponxcharacterdata.ServiceWeaponXCharacterData, armorService armorXCharacterData.ServiceArmorXCharacterData, itemService itemxcharacterdata.ServiceItemXCharacterData) ServiceTradeEvent {
	return &serviceTradeEvent{tradeEventRepo, characterTradeService, weaponService, armorService, itemService}
}
