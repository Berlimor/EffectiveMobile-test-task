-- Modify "song_details" table
ALTER TABLE "song_details" DROP CONSTRAINT "fk_song_details_music", ADD
 CONSTRAINT "fk_song_details_music" FOREIGN KEY ("music_id") REFERENCES "musics" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
