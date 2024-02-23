package domain

type CharacterData struct {
	CharacterId int64  `json:"characterid"`
	UserId      string `json:"userid"`
	CampaignId  int    `json:"campaignid"`
	Race        int    `json:"raceid"`     // TODO: Change Type from int to Race Struct
	Class       int    `json:"classid"`    // TODO: Change Type from int to Class Struct
	Background  int    `json:"background"` // TODO: Change Type from string to Background Struct
	Name        string `json:"name"`
	Story       string `json:"story"`
	Alignment   string `json:"alignment"`
	Age         int    `json:"age"`
	Hair        string `json:"hair"`
	Eyes        string `json:"eyes"`
	Skin        string `json:"skin"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	ImgUrl      string `json:"img"`
	Str         int    `json:"str"`
	Dex         int    `json:"dex"`
	Int         int    `json:"int"`
	Con         int    `json:"con"`
	Wiz         int    `json:"wiz"`
	Cha         int    `json:"cha"`
	Hitpoints   int    `json:"hitpoints"`
	HitDice     string `json:"hitDice"`
	Speed       int    `json:"speed"`
	ArmorClass  int    `json:"armorclass"`
	Level       int    `json:"level"`
	Exp         int    `json:"exp"`
}