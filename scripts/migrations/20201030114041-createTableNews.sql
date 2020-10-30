
-- +migrate Up
CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    site_id int,
    node_id int,
    title VARCHAR(255),
    link text UNIQUE,
    CONSTRAINT fk_sites
      FOREIGN KEY(site_id)
	     REFERENCES sites(id),
     CONSTRAINT fk_nodes
       FOREIGN KEY(node_id)
 	     REFERENCES nodes(id)
);
-- +migrate Down
ALTER TABLE IF EXISTS news
DROP CONSTRAINT IF EXISTS fk_nodes;

DROP TABLE IF EXISTS news;
