package dto

type EventDto struct {
	Idevent int `json:"idevent"`
	Type string `json:"type"`
	EventDescription string `json:"event_description"`
	Enviroment string `json:"enviroment"`
	Session_id int `json:"session_id"`
	Character_involved int `json:"character_id"`
	Dice_roll bool `json:"dice_roll"`
	Difficulty_Class bool `json:"difficulty_class"`
}