package event_type

var (
	QueryCreateEventType = `
		INSERT INTO event_type (name)
		VALUES (?)
	`

	QueryGetAll = `
		SELECT * FROM event_type;
	`

	QueryGetById = `
		SELECT * FROM event_type
		WHERE event_type_id = ?;
	`

	QueryGetByName = `
		SELECT * FROM event_type
		WHERE name = ?;
	`

	QueryUpdate = `
		UPDATE event_type
		SET name = ?
		WHERE event_type_id = ?;
	`

	QueryDelete = `
		DELETE FROM event_type WHERE event_type_id = ?;
	`
)
