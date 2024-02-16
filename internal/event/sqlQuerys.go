package event

var (
	QueryCreateEvent = `
		INSERT INTO event (type, event_description, environment, session_id, character_involved, dice_rolled, difficulty_class)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * FROM event;
	`

	QueryGetById = `
		SELECT * FROM event WHERE event_id = ?;
	`

	QueryGetBySessionId = `
		SELECT * FROM event WHERE session_id = ?;
	`

	QueryGetByCharacterId = `
		SELECT * FROM event WHERE character_involved = ?;
	`

	QueryUpdate = `
		UPDATE event
		SET type = ?, event_description = ?, environment = ?, session_id = ?, character_involved = ?, dice_rolled = ?, difficulty_class = ?
		WHERE event_id = ?;
	`

	QueryDelete = `
		DELETE FROM event WHERE event_id = ?;
	`
)
