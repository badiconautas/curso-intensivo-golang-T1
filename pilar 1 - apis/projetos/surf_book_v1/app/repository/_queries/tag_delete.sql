UPDATE user_tag
SET deleted_at = CURRENT_TIMESTAMP
WHERE tag_id = $1 AND user_id = $2;
