package domain

type Skill struct{
	SkillId int64 `json:"idskill"`
	Name string `json:"name"`
	Description string `json:"description"`
	CampaignId int64 `json:"campaign"`
}