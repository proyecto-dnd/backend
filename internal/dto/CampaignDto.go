package dto

import "github.com/proyecto-dnd/backend/internal/domain"

type ResponseCampaignDto struct {
	CampaignId    int              `json:"campaign_id"`
	DungeonMaster int              `json:"dungeon_master"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Image         string           `json:"image"`
	Notes         *string          `json:"notes"`
	Status        *string          `json:"status"`
	Sessions      []domain.Session `json:"sessions"`
}

type CreateCampaignDto struct {
	DungeonMaster int     `json:"dungeon_master"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Image         string  `json:"image"`
	Notes         *string `json:"notes"`
	Status        *string `json:"status"`
}
