-- update deleted at
UPDATE notebooks
SET 
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE notebook_id = $1 AND user_id = $2;
