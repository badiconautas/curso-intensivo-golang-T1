UPDATE notebooks
SET
  updated_at = CURRENT_TIMESTAMP,
  name = CASE WHEN $3::text IS NOT NULL AND $3::text <> '' THEN $3::text ELSE name END,
  description = CASE WHEN $4::text IS NOT NULL AND $4::text <> '' THEN $4::text ELSE description END,
  icon = CASE WHEN $5::text IS NOT NULL AND $5::text <> '' THEN $5::text ELSE icon END,
  image = CASE WHEN $6::text IS NOT NULL AND $6::text <> '' THEN $6::text ELSE image END
WHERE notebook_id = $1 AND user_id = $2;
