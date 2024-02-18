package dto

import "github.com/proyecto-dnd/backend/internal/domain"

type ResponseCampaignDto struct {
	DungeonMaster int              `json:"dungeon_master"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Image         string           `json:"image"`
	Sessions      []domain.Session `json:"sessions"`
}
