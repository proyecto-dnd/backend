package characterdata

var (
	QueryCreateCharacter = "INSERT INTO character_data (user_id, campaign_id, race_id, class_id, background_id, name, story, alignment, age, hair, eyes, skin, height, weight, img_url, str, dex, `int`, con, wiz, cha, hitpoints, hit_dice, speed, armor_class, level, exp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	QueryGetAll = `SELECT character_data.character_id, character_data.user_id, character_data.campaign_id,
	race.race_id, race.name, race.description, race.speed, race.str, race.dex, race.int, race.con, race.wiz, race.cha, 
	class.class_id, class.name, class.description, class.proficiency_bonus, class.hit_dice, class.armor_proficiencies, class.weapon_proficiencies, class.tool_proficiencies, class.spellcasting_ability, 
	background.background_id, background.name, background.languages, background.personality_traits, background.ideals, background.bonds, background.flaws, background.trait, background.tool_proficiencies,
	character_data.name, character_data.story, character_data.alignment, character_data.age, character_data.hair, character_data.eyes, character_data.skin, character_data.height, character_data.weight, character_data.img_url, character_data.str, character_data.dex, character_data.int, character_data.con, character_data.wiz, character_data.cha, character_data.hitpoints, character_data.hit_dice, character_data.speed, character_data.armor_class, character_data.level, character_data.exp
	FROM character_data left join race on character_data.race_id = race.race_id left join class on character_data.class_id = class.class_id left join background on character_data.background_id = background.background_id;`
	QueryGetById = ``
	QueryGetByUserId = ``
	QueryGetByUserIdAndCampaignId = ``
	QueryGetByCampaignId = ``
	QueryUpdate = ``
	QueryDelete = ``
)