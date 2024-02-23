package domain

type Item struct{
	Item_Id int64 `json:"iditem"`
	Name string `json:"name"`
    Weight int64 `json:"weight"`
	Price int64 `json:"price"`
	Description string `json:"description"`
	Campaign_Id int64 `json:"campaign"`
}