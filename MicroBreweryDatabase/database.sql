CREATE TYPE SI_unit as ENUM('mg','g','dag','kg','t');

CREATE TABLE "ingredient_category" (
  "id" SERIAL PRIMARY KEY,
  "category_name" varchar,
  "created_at" timestamp
);

CREATE TABLE "ingredient" (
  "id" SERIAL PRIMARY KEY,
  "ingredient_name" varchar(32),
  "unit" SI_unit NOT NULL DEFAULT 'g',
  "quantity" numeric,
  "created_at" timestamp,
  "ingredient_category_id" int
);

CREATE TABLE "recipe_category" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp
);

CREATE TABLE "recipe" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "ibu" numeric,
  "density" numeric,
  "created_at" timestamp,
  "recipe_category_id" int
);

CREATE TABLE "recipe_ingredient_list" (
  "quantity" numeric,
  "unit" SI_unit NOT NULL DEFAULT 'g',
  "created_at" timestamp,
  "recipe_id" int,
  "ingredient_id" int
);

CREATE TABLE "mash_stage" (
  "id" SERIAL PRIMARY KEY,
  "stage_time" int,
  "temperature" int,
  "pump_work" boolean,
  "created_at" timestamp,
  "recipe_id" int
);

CREATE TABLE "mash_tunes" (
  "id" SERIAL PRIMARY KEY,
  "mash_tune_code" varchar,
  "busy" boolean
);

CREATE TABLE "mesh_history" (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp,
  "start_at" timestamp,
  "end_at" timestamp,
  "recipe_id" int,
  "mash_tunes_id" int
);

ALTER TABLE "ingredient" ADD FOREIGN KEY ("ingredient_category_id") REFERENCES "ingredient_category" ("id");

ALTER TABLE "recipe" ADD FOREIGN KEY ("recipe_category_id") REFERENCES "recipe_category" ("id");

ALTER TABLE "recipe_ingredient_list" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "recipe_ingredient_list" ADD FOREIGN KEY ("ingredient_id") REFERENCES "ingredient" ("id");

ALTER TABLE "mash_stage" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "mesh_history" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("id");

ALTER TABLE "mesh_history" ADD FOREIGN KEY ("mash_tunes_id") REFERENCES "mash_tunes" ("id");

-- TABLE ingredient category:
insert into ingredient_category (category_name,created_at) values('',NOW());
insert into ingredient_category (category_name,created_at) values('chmiel',NOW());
insert into ingredient_category (category_name,created_at) values('drożdż',NOW());
-- TABLE ingredient
insert into ingredient (ingredient_name,unit,quantity,created_at,ingredient_category_id) values('Chmiel #1','g',0,NOW(),1);
insert into ingredient (ingredient_name,unit,quantity,created_at,ingredient_category_id) values('Chmiel cytrusowy','kg',125,NOW(),2);
insert into ingredient (ingredient_name,unit,quantity,created_at,ingredient_category_id) values('Drożdże','kg',125,NOW(),3);
-- TABLE recipe category:
insert into recipe_category(name,created_at) values('IPA',NOW());
-- TABLE recipe:
insert into recipe(name,created_at,recipe_category_id) values('IPA #1',NOW(),1);
-- TABLE recipe ingredient list:
insert into recipe_ingredient_list (quantity,unit,created_at,recipe_id,ingredient_id) values(5,'kg',NOW(),1,2);
insert into recipe_ingredient_list (quantity,unit,created_at,recipe_id,ingredient_id) values(5,'kg',NOW(),1,3);
--TABLE mash tun:
insert into mash_stage(temperature,created_at,recipe_id,pump_work,stage_time) values(60,NOW(),1,true,3600000);
insert into mash_stage(temperature,created_at,recipe_id,pump_work,stage_time) values(60,NOW(),1,true,600000);
