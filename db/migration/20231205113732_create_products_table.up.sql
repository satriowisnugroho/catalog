CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY,
  "sku" varchar NOT NULL,
  "title" varchar NOT NULL,
  "category" smallint NOT NULL,
  "condition" smallint NOT NULL,
  "tenant" smallint NOT NULL,
  "qty" integer NOT NULL,
  "price" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "products" ("sku");
CREATE INDEX ON "products" ("title");
CREATE INDEX ON "products" ("category");
CREATE INDEX ON "products" ("condition");
CREATE INDEX ON "products" ("tenant");
CREATE INDEX ON "products" ("created_at");
