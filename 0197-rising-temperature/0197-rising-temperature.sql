# Write your MySQL query statement below
SELECT w1.id
FROM Weather w1
INNER JOIN Weather w2
ON TO_DAYS(w1.recordDate) - TO_DAYS(w2.recordDate) = 1 and w1.temperature > w2.temperature