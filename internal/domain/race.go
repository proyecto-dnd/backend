package domain

type Race struct {
	RaceID      int    `json:"race_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Speed       int    `json:"speed"`
	Str         int    `json:"str"`
	Dex         int    `json:"dex"`
	Inte        int    `json:"inte"`
	Con         int    `json:"con"`
	Wiz         int    `json:"wiz"`
	Cha         int    `json:"cha"`
}