create table if not exists orders (
    id varchar(25) not null,
    body json not null
);

create or replace procedure push_order(
    id varchar(25),
    body json)
language plpgsql 
as $$
declare
begin
    insert into orders values(
    id,
    body);
end;
$$;