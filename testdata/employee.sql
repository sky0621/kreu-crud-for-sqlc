-- name: ListUserAndEmployee :many
SELECT u.l_name_kana, u.f_name_kana
FROM users AS u
UNION
SELECT e.l_name_kana, e.f_name_kana
FROM employee AS e
ORDER BY 1, 2;

-- name: ListEmployeeForAllSystemUses :many
with
/* 全てのシステムを取得して、全ての社員と cross join して直積結合を作ります。*/
    all_app_systems as (
        select
            *
        from
            (select
                 distinct(AppSystem) as AllAppSystems
             from
                 AppSystems) as app_systems
                cross join
            Employees
    ),

/*「全てのシステムを利用する」とは、使わないシステムがないという事なので、
   all_app_systems から使わないシステムがある社員を取得します。*/
    employee_not_using_all_systems as (
        select
            all_app_systems.Employee
        from
            all_app_systems
                left join EmployeeAppSystems as eas
                          on eas.AppSystem = all_app_systems.AllAppSystems
                              and eas.Employee = all_app_systems.Employee
        where
            eas.Employee is null
        group by
            all_app_systems.Employee
    ),

/* 使わないシステムがある社員を`Employee`から除外すれば、
   全てのシステムを利用する社員が取得できます。*/
    employee_uses_all_systems as (
        select
            e.Employee
        from
            Employees as e
                left join employee_not_using_all_systems as n
                          on e.Employee = n.Employee
        where
            n.Employee is null
    )
select * from employee_uses_all_systems;

-- name: ListEmployeeForDataAnalyst :many
with
/* データアナリスト用システムを取得します。*/
    data_analyst_systems as (
        select
            rs.AppSystem
        from
            RoleAppSystems as rs
        where
                rs.Role = 'Data Analyst'
    )

/* データアナリスト用システムのみを利用する社員を取得します。*/
select
    eas.Employee
from
    EmployeeAppSystems as eas
        left outer join data_analyst_systems as das
                        on das.AppSystem = eas.AppSystem
group by
    eas.Employee
having
        count(das.AppSystem) = (select count(*) from data_analyst_systems)
   and
        count(das.AppSystem) = count(*);
