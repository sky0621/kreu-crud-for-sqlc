-- name: ListBug :many
SELECT COUNT(bp.product_id)                    AS how_many_products,
       COUNT(dev.account_id)                   AS how_many_developers,
       COUNT(b.bug_id) / COUNT(dev.account_id) AS avg_bugs_per_developer,
       COUNT(cust.account_id)                  AS how_many_customers
FROM bugs b
         JOIN bugs_products bp ON (b.bug_id = bp.bug_id)
         JOIN accounts dev ON (b.assigned_to = dev.account_id)
         JOIN accounts cust ON (b.reported_by = cust.account_id)
WHERE cust.email NOT LIKE $1
GROUP BY bp.product_id;