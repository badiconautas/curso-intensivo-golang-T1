SELECT 
    notebook_id, user_id, name, description, icon, image, created_at, updated_at
FROM notebooks
WHERE user_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC