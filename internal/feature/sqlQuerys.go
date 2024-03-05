package feature

var (
	QueryCreateFeature = `
		INSERT INTO feature (name, description)
		VALUES (?, ?)
	`

	QueryGetAll = `
		SELECT * FROM feature;
	`

	QueryGetAllByCharacterId = `
		SELECT f.*
		FROM feature f
		INNER JOIN character_feature cf ON f.feature_id = cf.feature_id
		WHERE cf.character_id = ?;	
	`

	QueryGetById = `
		SELECT * FROM feature
		WHERE feature_id = ?;
	`

	QueryUpdate = `
		UPDATE feature
		SET name = ?, description = ?
		WHERE feature_id = ?;
	`

	QueryDelete = `
		DELETE FROM feature WHERE feature_id = ?;
	`
)