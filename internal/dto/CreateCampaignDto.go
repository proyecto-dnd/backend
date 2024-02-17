package dto

type CreateCampaignDto struct {
	DungeonMaster      int    `json:"dungeon_master"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Image              int    `json:"image"`
}
