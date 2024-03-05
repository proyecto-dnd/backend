package event

var (
	QueryCreateEvent = `
		INSERT INTO event (type, environment, session_id, event_protagonist_id, dice_rolled, difficulty_class, event_target, event_resolution)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	QueryGetAll = `
		SELECT e.*, et.name FROM event e
		JOIN event_type et ON e.type = et.event_type_id;
	`

	QueryGetById = `
		SELECT e.*, et.name FROM event e
		JOIN event_type et ON e.type = et.event_type_id
		WHERE e.event_id = ?;
	`

	QueryGetByTypeId = `
		SELECT e.*, et.name
		FROM event e
		JOIN event_type et ON e.type = ?;
	`

	QueryGetBySessionId = `
		SELECT e.*, et.name FROM event e
		JOIN event_type et ON e.type = et.event_type_id
		WHERE e.session_id = ?;
	`

	QueryGetByProtagonistId = `
		SELECT e.*, et.name FROM event e
		JOIN event_type et ON e.type = et.event_type_id
		WHERE e.event_protagonist_id = ?;
	`

	QueryUpdate = `
		UPDATE event
		SET type = ?, environment = ?, session_id = ?, event_protagonist_id = ?, dice_rolled = ?, difficulty_class = ?, event_target = ?, event_resolution = ?
		WHERE event_id = ?;
	`

	QueryDelete = `
		DELETE FROM event WHERE event_id = ?;
	`
)
