package domain

type Class struct {
	ClassId             int `json:"class_id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	ProficiencyBonus    int    `json:"proficiency_bonus"`
	HitDice             string `json:"hit_dice"`
	ArmorProficiencies  string `json:"armor_proficiencies"`
	WeaponProficiencies string `json:"weapon_proficiencies"`
	ToolProficiencies   string `json:"tool_proficiencies"`
	SpellcastingAbility int    `json:"spellcasting_ability"`
}
