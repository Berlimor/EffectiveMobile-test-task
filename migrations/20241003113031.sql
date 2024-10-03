-- Modify "musics" table
ALTER TABLE "musics" DROP COLUMN "artist", DROP COLUMN "text";
-- Create "song_details" table
CREATE TABLE "song_details" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "music_id" bigint NOT NULL,
  "release_date" text NULL,
  "text" text NULL,
  "link" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_song_details_music_id" UNIQUE ("music_id")
);
-- Create index "idx_song_details_deleted_at" to table: "song_details"
CREATE INDEX "idx_song_details_deleted_at" ON "song_details" ("deleted_at");
