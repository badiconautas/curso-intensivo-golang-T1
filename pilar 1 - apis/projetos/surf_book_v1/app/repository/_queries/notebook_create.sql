INSERT INTO notebooks (
	notebook_id,
	user_id,
	icon,
	name,
	description,
	image,
	created_at,
	updated_at,
	deleted_at 
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)