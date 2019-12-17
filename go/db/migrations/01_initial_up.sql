create table national_forest
(
    id bigserial not null
        constraint national_forest_id_pk
            primary key,
    name text not null
        constraint national_forest_name_key
            unique,
    website text not null default '',
    created_at timestamp with time zone not null default now()
);
comment on table national_forest is 'National Forest or BLM region';

create table district
(
    id bigserial not null
        constraint district_id_pk
            primary key,
    national_forest_id int not null
        constraint district_national_forest_id references national_forest,
    name text not null,
    map_location text not null default '',
    created_at timestamp with time zone not null default now()
);
comment on table district is 'National forest district';

create table site
(
	id bigserial not null
        constraint site_id_pk
            primary key,
	latitude DECIMAL(10, 8) not null,
	longitude DECIMAL(11, 8) not null,
	national_forest_id int null
	    constraint site_national_forest_id references national_forest,
	district_id int null
        constraint site_district_id references district,
	altitude int not null default 0,
	notes text not null default '',
    created_at timestamp with time zone not null default now()
);
comment on table site is 'Camping sites';

create unique index coordinates_uindex
	on site (latitude, longitude);

create table trip
(
	id bigserial not null
        constraint trip_id_pk
            primary key,
	site_id int not null
        constraint trip_site_id references site,
	start_date timestamp not null,
	length_of_trip int not null,
	notes text not null default '',
	rating int not null default 0,
	dog_friendly bool not null,
    created_at timestamp with time zone not null default now()
);
comment on table trip is 'Visit to a single site';
comment on column trip.length_of_trip is 'Number of days';
comment on column trip.rating is 'Rating from 1 to 5';

