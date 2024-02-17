package campaign

var (
	QueryCreateCampaign = `
		INSERT INTO campaign (dungeon_master, name, description, image)
		VALUES (?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * FROM campaign;
	`

	QueryGetById = `
		SELECT * FROM campaign WHERE campaign_id = ?;
	`

	QueryUpdate = `
		UPDATE event
		SET dungeon_master = ?, name = ?, description = ?, image = ?
		WHERE campaign_id = ?;
	`

	QueryDelete = `
		DELETE FROM campaign WHERE campaign_id = ?;
	`
)
