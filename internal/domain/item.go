package domain

type Item struct{
	Iditem int `json:"iditem"`
	Name string `json:"name"`
	Type string `json:"type"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	Character_id int `json:"character_id"`
	Equipped bool `json:"equipped"`
}