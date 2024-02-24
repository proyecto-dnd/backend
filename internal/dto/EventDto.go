package dto

import "github.com/proyecto-dnd/backend/internal/domain"

type CreateEventDto struct {
	Type               int    `json:"type"`
	Environment        string `json:"environment"`
	Session_id         int    `json:"session_id"`
	EventProtagonistId int    `json:"event_protagonist_id"`
	Dice_rolled        bool   `json:"dice_rolled"`
	Difficulty_Class   bool   `json:"difficulty_class"`
	EventTarget        string `json:"event_target"`
	EventResolution    string `json:"event_resolution"`
}

type EventRepositoryResponseDto struct {
	EventId            int    `json:"event_id"`
	Type               int    `json:"type"`
	Environment        string `json:"environment"`
	Session_id         int    `json:"session_id"`
	EventProtagonistId int    `json:"event_protagonist_id"`
	Dice_rolled        bool   `json:"dice_rolled"`
	Difficulty_Class   bool   `json:"difficulty_class"`
	EventTarget        string `json:"event_target"`
	EventResolution    string `json:"event_resolution"`
	TypeName           string `json:"type_name"`
}

type ResponseEventDto struct {
	EventId            int                `json:"event_id"`
	Type               domain.EventType   `json:"type"`
	Environment        string             `json:"environment"`
	Session_id         int                `json:"session_id"`
	EventProtagonistId int                `json:"event_protagonist_id"`
	Dice_rolled        bool               `json:"dice_rolled"`
	Difficulty_Class   bool               `json:"difficulty_class"`
	EventTarget        string             `json:"event_target"`
	EventResolution    string             `json:"event_resolution"`
	Affected           []domain.Character `json:"affected"`
}
