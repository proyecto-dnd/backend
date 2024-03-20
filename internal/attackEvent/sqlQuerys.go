package attackEvent

var (
	QueryCreateAttackEvent = `
		INSERT INTO attack_event (type, environment, session_id, event_protagonist_id, event_resolution, weapon, spell, dmg_type, description, timestamp)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * from attack_event ae;
	`

	QueryGetById = `
		SELECT ae.*, s.*
		FROM attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		WHERE ae.event_id = ?;
	`

	QueryGetBySessionId = `
		SELECT ae.*, s.*
		FROM attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		WHERE ae.session_id = ?;
	`

	QueryGetByProtagonistId = `
		SELECT ae.*, s.*
		FROM attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		WHERE ae.event_protagonist_id = ?;
	`

	QueryGetByAffectedId = `
		SELECT ae.*, s.* from attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		INNER JOIN character_attack_event cae ON ae.event_id = cae.event_id 
		WHERE cae.character_id = ?
	`

	QueryGetByProtagonistIdAndAffectedId = `
		SELECT ae.*, s.* from attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		INNER JOIN character_attack_event cae ON ae.event_id = cae.event_id 
		WHERE cae.character_id = ? AND ae.event_protagonist_id = ?;
	`

	QueryUpdate = `
		UPDATE attack_event
		SET type = ?, environment = ?, session_id = ?, event_protagonist_id = ?, event_resolution = ?, dmg_type = ?, weapon = ?, spell = ?, description = ?, timestamp = ?
		WHERE event_id = ?;
	`

	QueryDelete = `
		DELETE FROM attack_event WHERE event_id = ?;
	`
)
