package weapon

var (
	QueryCreateWeapon = `
		INSERT INTO weapon (weapon_type, name, weight, price, category, reach, description, damage, versatile_damage, ammunition, damage_type, basic)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	QueryGetAllWeapons = `
		SELECT * FROM weapon;
	`

	QueryGetWeaponById = `
		SELECT * FROM weapon WHERE weapon_id = ?;
	`

	QueryUpdateWeapon = `
		UPDATE weapon
		SET weapon_type = ?, name = ?, weight = ?, price = ?, category = ?, reach = ?, description = ?, damage = ?, versatile_damage = ?, ammunition = ?, damage_type = ?, basic = ?
		WHERE weapon_id = ?;
	`

	QueryDeleteWeapon = `
		DELETE FROM weapon WHERE weapon_id = ?;
	`
)
