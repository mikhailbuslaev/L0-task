create table if not exists orders (
    id varchar(25) not null,
    track_number varchar(25),
    entry_ varchar(25),
    delivery_name varchar(25),
    delivery_phone varchar(25),
    delivery_zip varchar(25),
    delivery_city varchar(25),
    delivery_address varchar(25),
    delivery_region varchar(25),
    delivery_email varchar(25),
    payment_transaction varchar(25),
    payment_request_id varchar(25),
    payment_currency varchar(25),
    payment_provider varchar(25),
    payment_amount int,
    payment_payment_dt int,
    payment_bank varchar(25),
    payment_delivery_cost int,
    payment_goods_total int,
    payment_custom_fee int,
    chrt_id int,
    locale varchar(25),
    internal_signature varchar(25),
    customer_id varchar(25),
    delivery_service varchar(25),
    shardkey varchar(25),
    sm_id int,
    date_created varchar(25),
    oof_shard varchar(25)
);
create table if not exists items (
    chrt_id int,
    track_number varchar(25),
    price int,
    rid varchar(25),
    name_ varchar(25),
    sale int,
    size varchar(25),
    total_price int,
    nm_id int,
    brand varchar(25),
    status_ int
);

create or replace procedure orders()
language plpgsql 
as $$
declare
begin
    select * from orders;
end;
$$;

create or replace procedure items(chrt_id integer)
language plpgsql 
as $$
declare
begin
    select * from items where items.chrt_id = chrt_id;
end;
$$;

create or replace procedure push_order(
    id varchar(25),
    track_number varchar(25),
    entry_ varchar(25),
    delivery_name varchar(25),
    delivery_phone varchar(25),
    delivery_zip varchar(25),
    delivery_city varchar(25),
    delivery_address varchar(25),
    delivery_region varchar(25),
    delivery_email varchar(25),
    payment_transaction varchar(25),
    payment_request_id varchar(25),
    payment_currency varchar(25),
    payment_provider varchar(25),
    payment_amount int,
    payment_payment_dt int,
    payment_bank varchar(25),
    payment_delivery_cost int,
    payment_goods_total int,
    payment_custom_fee int,
    chrt_id int,
    locale varchar(25),
    internal_signature varchar(25),
    customer_id varchar(25),
    delivery_service varchar(25),
    shardkey varchar(25),
    sm_id int,
    date_created varchar(25),
    oof_shard varchar(25))
language plpgsql 
as $$
declare
begin
    insert into orders values(
    id,
    track_number,
    entry_,
    delivery_name,
    delivery_phone,
    delivery_zip,
    delivery_city,
    delivery_address,
    delivery_region,
    delivery_email,
    payment_transaction,
    payment_request_id,
    payment_currency,
    payment_provider,
    payment_amount,
    payment_payment_dt,
    payment_bank,
    payment_delivery_cost,
    payment_goods_total,
    payment_custom_fee,
    chrt_id,
    locale,
    internal_signature,
    customer_id,
    delivery_service,
    shardkey,
    sm_id,
    date_created,
    oof_shard );
end;
$$;

create or replace procedure push_item(
    chrt_id int,
    track_number varchar(25),
    price int,
    rid varchar(25),
    name_ varchar(25),
    sale int,
    size varchar(25),
    total_price int,
    nm_id int,
    brand varchar(25),
    status_ int)
language plpgsql 
as $$
declare
begin
    insert into items values(
    chrt_id,
    track_number,
    price,
    rid,
    name_,
    sale,
    size,
    total_price,
    nm_id,
    brand,
    status_);
end;
$$;