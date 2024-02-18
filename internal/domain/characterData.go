package domain


type CharacterData struct {
	CharacterId int64  `json:"characterid"`
	UserId      string `json:"userid"`
	Name        string `json:"name"`
	Class     int    `json:"classid"`
	Race      int    `json:"raceid"`
	Background  string `json:"background"`
	Hitpoints   int    `json:"hitpoints"`
	Speed       int    `json:"speed"`
	ArmorClass  int    `json:"armorclass"`
	Level       int    `json:"level"`
	Exp         int    `json:"exp"`
	CampaignId  int    `json:"campaignid"`
	Str         int    `json:"str"`
	Dex         int    `json:"dex"`
	Int         int    `json:"int"`
	Wiz         int    `json:"wiz"`
	Con         int    `json:"con"`
	Cha         int    `json:"cha"`
}