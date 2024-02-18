package user_campaign

var (
	QueryCreateUserCampaign = `
		INSERT INTO user_campaign (campaign_id, user_id, is_dm, is_owner)
		VALUES (?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * FROM user_campaign;
	`

	QueryGetById = `
		SELECT * FROM user_campaign WHERE user_campaign_id = ?;
	`

	QueryGetByCampaignId = `
		SELECT * FROM user_campaign WHERE campaign_id = ?;
	`

	QueryGetByUserId = `
		SELECT * FROM user_campaign WHERE user_id = ?;
	`

	QueryUpdate = `
		UPDATE user_campaign
		SET campaign_id = ?, user_id = ?, is_dm = ?, is_owner = ?
		WHERE user_campaign_id = ?;
	`

	QueryDelete = `
		DELETE FROM user_campaign WHERE user_campaign_id = ?;
	`
)