create table site
(
	id int primary key auto_increment,
	latitude DECIMAL(10, 8) not null,
	longitude DECIMAL(11, 8) not null,
	national_forest_id int not null,
	district_id int not null,
	altitude int null,
	notes text null
)
comment 'Camping sites';

create unique index coordinates_uindex
	on site (latitude, longitude);


create table trip
(
	id int primary key auto_increment,
	site_id int not null,
	start_date timestamp not null comment 'Trip start date',
	length_of_trip int null comment 'Number of days',
	notes text null,
	rating int null comment 'Rating from 1 to 5',
	dog_friendly bool null
)
comment 'Visit to a single site';

create table national_forest
(
	id int primary key auto_increment,
	name text not null,
	website text null
)
comment 'National Forest or BLM region';
create unique index name_uindex on national_forest (name(500));

create table district
(
	id int primary key auto_increment,
	national_forest_id int not null,
	name text not null,
	map_location text null
)
comment 'National forest district';

insert into national_forest (name) values
    ('Arapahoe National Forest');

insert into district (national_forest_id, name) values
    (1, 'Canyon Lakes North');

insert into site (latitude, longitude, national_forest_id, district_id, altitude, notes) values
    (38.252311, -105.673388, 1, 1, 5280, 'abc');
