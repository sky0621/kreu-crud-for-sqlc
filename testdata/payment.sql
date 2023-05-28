-- name: ListPaymentAndCustomer :many
WITH test1 AS
         (
             SELECT customer_id, SUM(amount) AS total_payment FROM payment
             GROUP BY customer_id
         )

SELECT test1.customer_id, test1.total_payment, customer.first_name
FROM test1
         INNER JOIN customer
                    ON test1.customer_id=customer.customer_id
WHERE test1.total_payment>150;
