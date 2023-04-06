CREATE DATABASE cost_per_use_tracker;

\c cost_per_use_tracker
CREATE TABLE tags (
    tag_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tag_name text
);


CREATE TABLE categories(
    category_id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_name text
    );

CREATE TABLE items (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text,
    price numeric,
    uses int,
    date_bought DATE,
    notes text,
    active boolean,
	category integer REFERENCES categories (category_id)
);

CREATE TABLE tag_associations (
  match_id integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  tag integer REFERENCES tags (tag_id),
  item integer REFERENCES items (id)
);

/* Insertions for tests*/
INSERT INTO categories (category_name) VALUES ('Cooking');
INSERT INTO tags (tag_name) VALUES ('Spontankauf');
INSERT INTO tags (tag_name) VALUES ('Lebensdauer');
INSERT INTO items (name, price, uses, date_bought, notes, category) VALUES ('Pan', 10.15, 10, '2020-10-10', 'gekauft blbalabllablabliblub', 1);
INSERT INTO tag_associations (tag, item) VALUES (1,1);

/*
select tags.tag_name FROM items i, tag_associations ta, tags WHERE i.id = ta.item and ta.tag = tags.tag_id AND i.id = 1; for showing the tags of an item by id.

SELECT json_agg(items) from items;
*/
