package dto

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"time"
)

type CreateAttackEventDto struct {
	Type               string     `json:"type"`
	Environment        string     `json:"environment"`
	Session_id         int        `json:"session_id"`
	EventProtagonistId int        `json:"event_protagonist_id"`
	EventResolution    string     `json:"event_resolution"`
	Weapon             *int       `json:"weapon"`
	Spell              *int       `json:"spell"`
	DmgType            *string    `json:"dmg_type"`
	Description        *string    `json:"description"`
}

type RepositoryResponseAttackEvent struct {
	AttackEventId            int        `json:"attack_event_id"`
	Type                     string     `json:"type"`
	Environment              string     `json:"environment"`
	Session_id               int        `json:"session_id"`
	EventProtagonistId       int        `json:"event_protagonist_id"`
	EventResolution          string     `json:"event_resolution"`
	DmgType                  *string    `json:"dmg_type"`
	Weapon                   *int       `json:"weapon"`
	Spell                    *int       `json:"spell"`
	Description              *string    `json:"description"`
	TimeStamp                *time.Time `json:"time_stamp"`
	SessionSessionId         int        `json:"session_session_id"`
	Start                    time.Time  `json:"start"`
	End                      time.Time  `json:"end"`
	SessionDescription       string     `json:"session_description"`
	SessionCampaignId        int        `json:"session_campaign_id"`
	SessionCurrentEnviroment *string    `json:"session_current_enviroment"`
}

type ResponseEventDto struct {
	AttackEventId    int                 `json:"attack_event_id"`
	Type             string              `json:"type"`
	Environment      string              `json:"environment"`
	Session          domain.Session      `json:"session"`
	EventProtagonist CharacterCardDto   `json:"event_protagonist"`
	EventResolution  string              `json:"event_resolution"`
	Weapon           *int                `json:"weapon"`
	Spell            *int                `json:"spell"`
	DmgType          *string             `json:"dmg_type"`
	Description      *string             `json:"description"`
	TimeStamp        *time.Time          `json:"time_stamp"`
	Affected         []CharacterCardDto `json:"affected"`
}
