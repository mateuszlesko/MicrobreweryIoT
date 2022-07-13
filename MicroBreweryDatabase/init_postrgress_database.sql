DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS ingredients;
DROP TYPE IF EXISTS SI_units;

CREATE TABLE categories(
   category_id INT GENERATED ALWAYS AS IDENTITY,
   category_name VARCHAR(255) NOT NULL,
   PRIMARY KEY(category_id)
);

CREATE TYPE SI_unit as ENUM('mg','g','dag','kg','t');

CREATE TABLE ingredients(
   ingredient_id INT GENERATED ALWAYS AS IDENTITY,
   ingredient_name varchar(64) NOT NULL,
   unit SI_unit NOT NULL DEFAULT 'g',
   quantity REAL NOT NULL,
   description TEXT,
   category_id INT,
   PRIMARY KEY(ingredient_id),
   CONSTRAINT fk_category
      FOREIGN KEY(category_id) 
	  REFERENCES categories(category_id)
);

-- Categories data:
INSERT INTO categories(
    category_name
)
VALUES(
    'Chmiel cytrusowy'
);

INSERT INTO categories(
    category_name
)
VALUES(
    'Chmiel amerykański'
);


-- ingredients data:
INSERT INTO ingredients(
    ingredient_name,
    unit,
    quantity,
    category_id
)
VALUES(
    'Marynka',
    'kg',
    25,
    1
);

INSERT INTO ingredients(
    ingredient_name,
    unit,
    quantity,
    category_id
)
VALUES(
    'Wai-iti',
    'kg',
    25,
    1
);