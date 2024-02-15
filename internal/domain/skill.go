package domain

type Skill struct{
	Idskill int `json:"idskill"`
	Name string `json:"name"`
	Description string `json:"description"`
	Dharacter_id int `json:"character_id"`
}