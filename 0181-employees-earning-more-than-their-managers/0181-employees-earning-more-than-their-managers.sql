# Write your MySQL query statement below
SELECT e1.name as Employee
FROM Employee e1
WHERE e1.managerId is not null and e1.salary > (select salary from Employee e2 where e1.managerId = e2.id)