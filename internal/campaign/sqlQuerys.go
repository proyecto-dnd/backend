package campaign

var (
	QueryCreateCampaign = `
		INSERT INTO campaign (dungeon_master, name, description, image, notes, status)
		VALUES (?, ?, ?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * FROM campaign;
	`

	QueryGetById = `
		SELECT * FROM campaign WHERE campaign_id = ?;
	`

	QueryGetByUserId = `
		SELECT c.*
		FROM campaign c
		JOIN user_campaign uc ON c.campaign_id = uc.campaign_id
		WHERE uc.user_id = ?;
	
	`

	QueryUpdate = `
		UPDATE campaign
		SET dungeon_master = ?, name = ?, description = ?, image = ? , notes = ?, status = ?
		WHERE campaign_id = ?;
	`

	QueryDelete = `
		DELETE FROM campaign WHERE campaign_id = ?;
	`
)
