package domain

type CharacterXSpellEvent struct {
	CharacterSpellEventId int `json:"character_spell_event_id"`
	CharacterId           int `json:"character_id"`
	SpellEventId          int `json:"spell_event_id"`
}
