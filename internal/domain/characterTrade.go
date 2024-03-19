package domain

type CharacterTrade struct {
	CharacterTrade_Id int `json:"character_trade_id"`
	TradeEvent_Id int `json:"trade_event_id"`
	Weapon *int `json:"weapon"`
	Item *int `json:"item"` 
	Armor *int `json:"armor"`
	ItemOwner int `json:"item_owner"`
	ItemReciever int `json:"item_reciever"`
	Quantity *int `json:"quantity"`
	ItemName string `json:"item_name"`
	ItemType string `json:"item_type"`
}