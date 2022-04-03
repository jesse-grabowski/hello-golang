CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE TABLE entities (
    id int primary key generated always as identity,
    display_name varchar(256)
);

CREATE TABLE samples (
    time timestamp not null default now(),
    entity_id int not null,
    coordinates geography not null,
    CONSTRAINT fk_sample_entity FOREIGN KEY (entity_id) REFERENCES entities(id)
);
CREATE INDEX idx_samples ON samples USING gist (time, coordinates, entity_id)