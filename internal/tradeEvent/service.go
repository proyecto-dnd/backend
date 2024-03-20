package tradeevent

import (
	charactertrade "github.com/proyecto-dnd/backend/internal/characterTrade"
	"github.com/proyecto-dnd/backend/internal/domain"
)

type serviceTradeEvent struct {
	tradeEventRepo RepositoryTradeEvent
	characterTradeService charactertrade.ServiceCharacterTrade
}

// Create implements ServiceTradeEvent.
func (s *serviceTradeEvent) Create(tradeEvent domain.TradeEvent) (domain.TradeEvent, error) {
	newTradeEvent, err := s.tradeEventRepo.Create(tradeEvent)
	if err != nil {
		return domain.TradeEvent{}, err
	}

	for i := range tradeEvent.TradingItems{
		tradeEvent.TradingItems[i].TradeEvent_Id = newTradeEvent.TradeEvent_Id
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
	err := s.characterTradeService.DeleteByTradeEventId(id)
	if err != nil {
		return err
	}
	err = s.tradeEventRepo.Delete(id)
	return err
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

func NewTradeEventService(tradeEventRepo RepositoryTradeEvent, characterTradeService charactertrade.ServiceCharacterTrade) ServiceTradeEvent {
	return &serviceTradeEvent{tradeEventRepo, characterTradeService}
}
