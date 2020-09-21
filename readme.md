
# 依赖包
```go
go get github.com/Unknwon/goconfig
go get github.com/bmizerany/pq
go get github.com/gorhill/cronexpr
```




```sql
create or replace PROCEDURE sp_t1()
language plpgsql
as $$
declare
v_sec varchar;
begin
select pg_sleep(5) into v_sec;
insert into t1(id,ins_date,proname,stat) values(1,now(),'t1',1);
raise notice '执行完毕！'; 
end;
$$;

create or replace PROCEDURE sp_t2()
language plpgsql
as $$
declare
v_sec varchar;
begin
select pg_sleep(12000) into v_sec;
raise notice '执行完毕！'; 
end;
$$;


create table t1(
id int,
ins_date timestamp,
proname varchar,
stat int
);

-- job基本任务配置
create table ds_job(
jobno int,
nexttime timestamp,
interval varchar,
what varchar,
stat int default 0,
primary key(jobno)
);

delete from ds_job;
insert into ds_job(jobno,nexttime,interval,what) values(1,'2020-09-01','*/30 * * * * * *','call sp_t1()');
insert into ds_job(jobno,nexttime,interval,what) values(2,'2020-09-01','71 * * * * * *','call sp_t1()');


```