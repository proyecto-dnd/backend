package domain

type Item struct{
	ItemId int64 `json:"iditem"`
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	CampaignId int64 `json:"campaign"`
	Equipable bool `json:" equipable"`
}