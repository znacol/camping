create table site
(
	id bigserial not null,
	latitude DECIMAL(10, 8) not null,
	longitude DECIMAL(11, 8) not null,
	national_forest_id int not null,
	district_id int not null,
	altitude int null,
	notes text null
);
comment on table site is 'Camping sites';

create unique index coordinates_uindex
	on site (latitude, longitude);

create table trip
(
	id bigserial not null,
	site_id int not null,
	start_date timestamp not null,
	length_of_trip int null,
	notes text null,
	rating int null,
	dog_friendly bool null
);
comment on table trip is 'Visit to a single site';
comment on column trip.length_of_trip is 'Number of days';
comment on column trip.rating is 'Rating from 1 to 5';

create table national_forest
(
	id bigserial not null,
	name text not null
        constraint national_forest_name_key
            unique,
	website text null
);
comment on table national_forest is 'National Forest or BLM region';

create table district
(
	id bigserial not null,
	national_forest_id int not null,
	name text not null,
	map_location text null
);
comment on table district is 'National forest district';

insert into national_forest (name) values
    ('Arapahoe National Forest');

insert into district (national_forest_id, name) values
    (1, 'Canyon Lakes North');
