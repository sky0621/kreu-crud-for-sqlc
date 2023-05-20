-- name: ListOrder :many
SELECT p.p_name, SUM(o.quantity), SUM(p.price * o.quantity)
FROM order_desc AS o RIGHT JOIN product AS p
                                ON p.p_id = o.p_id
GROUP BY p.p_id, p.p_name ORDER BY SUM(p.price * o.quantity) DESC;
