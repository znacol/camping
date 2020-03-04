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

create table national_forest_district
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
comment on table national_forest_district is 'National forest district';

create table site
(
	id bigserial not null
        constraint site_id_pk
            primary key,
	latitude DECIMAL(10, 8) not null,
	longitude DECIMAL(11, 8) not null,
	altitude int not null default 0,
	notes text not null default '',
    created_at timestamp with time zone not null default now()
);
comment on table site is 'Camping sites';

create unique index coordinates_uindex
	on site (latitude, longitude);

create table site_has_national_forest_district
(
   site_id bigint not null
       constraint site_id_fk
           references site,
   national_forest_district_id bigint not null
       constraint national_forest_district_id_fk
           references national_forest_district,
   created_at timestamp with time zone not null default now(),
   constraint site_has_national_forest_district
       primary key (site_id, national_forest_district_id)
);
comment on table site_has_national_forest_district is 'Linking table to assign a national forest district to a site';

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

