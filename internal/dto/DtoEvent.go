package dto

type EventDto struct {
	Type string `json:"type"`
	EventDescription string `json:"event_description"`
	Environment string `json:"Environment"`
	Session_id int `json:"session_id"`
	Character_involved int `json:"character_id"`
	Dice_roll bool `json:"dice_roll"`
	Difficulty_Class bool `json:"difficulty_class"`
}