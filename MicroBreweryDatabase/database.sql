CREATE TYPE SI_unit as ENUM('mg','g','dag','kg','t');

CREATE TABLE "ingredient_category" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp
);

CREATE TABLE "ingredient" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(32),
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
