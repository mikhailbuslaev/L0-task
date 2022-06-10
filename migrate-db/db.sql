create table if not exists orders (
    id varchar(25) not null,
    track_number varchar(25),
    entry varchar(25),
    delivery_name varchar(50),
    delivery_phone varchar(25),
    delivery_zip varchar(10),
    delivery_city varchar(25),
    delivery_address varchar(50),
    delivery_region varchar(50),
    delivery_email varchar(50),
    payment_transaction varchar(25),
    payment_request_id varchar(25),
    payment_currency varchar(5),
    payment_provider varchar(25),
    payment_amount int,
    payment_payment_dt int,
    payment_bank varchar(25),
    payment_delivery_cost int,
    payment_goods_total int,
    payment_custom_fee int,
    chrt_id int,
    locale varchar(5),
    internal_signature varchar(25),
    customer_id varchar(25),
    delivery_service varchar(25),
    shardkey varchar(5),
    sm_id int,
    date_created varchar(25),
    oof_shard varchar(5)
);
create table if not exists items (
    chrt_id int,
    track_number varchar(25),
    price int,
    rid varchar(25),
    name varchar(25),
    sale int,
    size varchar(10),
    total_price int,
    nm_id int,
    brand varchar(25),
    status int
);

create or replace procedure orders()
language plpgsql 
as $$
declare
begin
    select * from orders;
end;
$$;

create or replace procedure items("chrt_id" integer)
language plpgsql 
as $$
declare
begin
    select * from items where items.chrt_id = "chrt_id";
end;
$$;