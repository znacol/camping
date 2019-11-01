insert into national_forest (id, name) values
    (1, 'Arapahoe National Forest');
-- Manually inserting id requires resetting the counter
SELECT setval('national_forest_id_seq', (2));

insert into district (id, national_forest_id, name) values
    (1, 1, 'Canyon Lakes North');
-- Manually inserting id requires resetting the counter
SELECT setval('district_id_seq', (2));

insert into site (latitude, longitude, national_forest_id, district_id, altitude, notes) values
    (38.252311, -105.673388, 1, 1, 5280, 'july 4th');
