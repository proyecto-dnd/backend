package domain

type Event struct {
	EventId            int    `json:"event_id"`
	Type               string `json:"type"`
	EventDescription   string `json:"event_description"`
	Environment        string `json:"environment"`
	Session_id         int    `json:"session_id"`
	Character_involved int    `json:"character_id"`
	Dice_roll          bool   `json:"dice_roll"`
	Difficulty_Class   bool   `json:"difficulty_class"`
}
