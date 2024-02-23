package domain

type Item struct{
	ItemId int64 `json:"iditem"`
	Name string `json:"name"`
    Weight int64 `json:"weight"`
	Price int64 `json:"price"`
	Description string `json:"description"`
	CampaignId int64 `json:"campaign"`
}