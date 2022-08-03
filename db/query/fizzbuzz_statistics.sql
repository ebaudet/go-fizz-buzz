-- name: GetMostUsedRequest :one
SELECT * FROM fizzbuzz_statistics
ORDER BY count DESC
LIMIT 1;

-- name: IncrementRequest :one
INSERT INTO fizzbuzz_statistics (
  request, count
) VALUES (
  $1, 1
) ON CONFLICT (request) DO
UPDATE SET
  count = fizzbuzz_statistics.count + 1,
  updated_at = now()
RETURNING *;
