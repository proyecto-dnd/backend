package session

var (
	QueryCreateSession = `
		INSERT INTO session (start, end, description, campaign_id)
		VALUES (?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * FROM session;
	`

	QueryGetById = `
		SELECT * FROM session WHERE session_id = ?;
	`

	QueryGetByCampaignId = `
		SELECT * FROM session WHERE campaign_id = ?;
	`

	QueryUpdate = `
		UPDATE session
		SET start = ?, end = ?, description = ?, campaign_id = ?
		WHERE session_id = ?;
	`

	QueryDelete = `
		DELETE FROM session WHERE session_id = ?;
	`

	QueryGetCampaignById = `
		SELECT * FROM campaign WHERE campaign_id = ?;
	`
)
