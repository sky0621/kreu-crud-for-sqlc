-- name: ListGuest :many
SELECT answer1, answer2
FROM quest
WHERE age >
      (SELECT AVG(age) FROM quest2)
ORDER BY answer1 ASC;
