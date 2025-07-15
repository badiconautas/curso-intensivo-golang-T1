INSERT INTO user_tag (
	tag_id,
	user_id,
	tag_name,
    created_at,
    deleted_at,
	color
) VALUES ($1, $2, $3, $4, $5, $6);
