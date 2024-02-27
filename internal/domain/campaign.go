package domain

type Campaign struct {
	CampaignId    int    `json:"campaign_id"`
	DungeonMaster int    `json:"dungeon_master"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Image         string `json:"image"`
}
