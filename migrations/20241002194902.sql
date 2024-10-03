-- Create "musics" table
CREATE TABLE "musics" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NOT NULL,
  "artist" text NOT NULL,
  "group" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_musics_deleted_at" to table: "musics"
CREATE INDEX "idx_musics_deleted_at" ON "musics" ("deleted_at");
