
-- +migrate Up
CREATE TABLE sites (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE
);

CREATE TABLE nodes (
    id SERIAL PRIMARY KEY,
    site_id int,
    value VARCHAR(255),
    CONSTRAINT fk_sites
      FOREIGN KEY(site_id)
	     REFERENCES sites(id)
);

-- +migrate Down
ALTER TABLE IF EXISTS news
  DROP CONSTRAINT IF EXISTS fk_sites;

ALTER TABLE IF EXISTS nodes
  DROP CONSTRAINT IF EXISTS fk_sites;

DROP TABLE IF EXISTS nodes;
DROP TABLE IF EXISTS sites;
