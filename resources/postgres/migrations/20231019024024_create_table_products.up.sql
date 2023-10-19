create table products (
	id serial CONSTRAINT rmr_remisier_pk PRIMARY KEY,
	item VARCHAR(255) null,
	price bigint null,
	is_active bool null,
	deleted_at timestamp null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp default CURRENT_TIMESTAMP null
);