package event

var (
	QueryCreateEvent = `
		INSERT INTO event (type, environment, session_id, event_protagonist_id, dice_rolled, difficulty_class, event_target, event_resolution)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	QueryGetAll = `
		SELECT * FROM event;
	`

	QueryGetById = `
		SELECT * FROM event WHERE event_id = ?;
	`

	QueryGetByTypeId = `
		SELECT * FROM event WHERE type = ?;
	`

	QueryGetBySessionId = `
		SELECT * FROM event WHERE session_id = ?;
	`

	QueryGetByProtagonistId = `
		SELECT * FROM event WHERE event_protagonist_id = ?;
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
