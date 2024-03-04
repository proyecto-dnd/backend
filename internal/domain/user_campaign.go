package domain

type UserCampaign struct {
	UserCampaignId int    `json:"user_campaign_id"`
	CampaignId     int    `json:"campaign_id"`
	UserId         string `json:"user_id"`
	IsDm           int    `json:"is_dm"`
	IsOwner        int    `json:"is_owner"`
}
