package dto

type CharacterCardDto struct {
	CharacterId int    `json:"character_id"`
	UserId      int    `json:"user_id"`
	CampaignID  int    `json:"campaign_id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Class       string `json:"class"`
	Level       int    `json:"level"`
	HitPoints   int    `json:"hit_points"`
}
