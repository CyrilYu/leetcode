select e1.Name as Employee
from (select Name, ManagerId, Salary from Employee where ManagerId IS NOT NULL) as e1
left join Employee as e2
on e1.ManagerId = e2.Id
where e1.Salary > e2.Salary

