SELECT 
	tag_id,
	user_id,
	tag_name,
	color,
    created_at,
    deleted_at
FROM user_tag
WHERE user_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC;
