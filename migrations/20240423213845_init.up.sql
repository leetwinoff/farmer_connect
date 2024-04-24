create extension citext;

create extension "uuid-ossp";


create table users (
    id uuid primary key default gen_random_uuid (),
    oauth_provider text,
    username citext unique,
    telegram_id bigint unique,
    role text,
    created_at timestamptz not null default now()
 );

create table consumer (
    id uuid primary key default gen_random_uuid (),
    user_id uuid not null references users (id) on delete cascade,
    full_name text,
    profile_pic text,
    description text,
    location text,

    created_at timestamptz not null default now()
 );

create table farmer (
    id uuid primary key default gen_random_uuid (),
    user_id uuid not null references users (id) on delete cascade,
    full_name text,
    company_name text,
    profile_pic text,
    description text,
    rating integer,
    location text,
    created_at timestamptz not null default now()
 );

create table product (
    id uuid primary key default gen_random_uuid (),
    farmer_id uuid not null references farmer (id) on delete cascade,
    name text not null,
    type text,
    image text,
    description text,
    price float not null,
    unit text,
    quantity integer,
    created_at timestamptz not null default now()
 );

create table saved_farmers (
    id uuid primary key default gen_random_uuid (),
    user_id uuid not null references users (id) on delete cascade,
    farmer_id uuid not null references farmer (id) on delete cascade,
    created_at timestamptz not null default now()
 );

create table saved_products (
    id uuid primary key default gen_random_uuid (),
    product_id uuid not null references product (id) on delete cascade,
    farmer_id uuid not null references farmer (id) on delete cascade,
    created_at timestamptz not null default now()
 );

create table review (
    id uuid primary key default gen_random_uuid (),
    consumer_id uuid not null references consumer (id) on delete cascade,
    farmer_id uuid not null references farmer (id) on delete cascade,
    comment text,
    created_at timestamptz not null default now()
 );
)
