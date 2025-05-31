--────────────────────────────────────
-- Drop tables in reverse dependency order
--────────────────────────────────────

--────────────────────────────────────
-- Video Age Rating Category - Content categories
--────────────────────────────────────

DROP TABLE IF EXISTS video_age_rating_category CASCADE;

--────────────────────────────────────
-- Video Metadata table - AI-generated data
--────────────────────────────────────

DROP INDEX IF EXISTS idx_video_metadata_video;
DROP TABLE IF EXISTS video_metadata CASCADE;

--────────────────────────────────────
-- Video table - Video files and URLs
--────────────────────────────────────

DROP INDEX IF EXISTS idx_video_updated_at;
DROP INDEX IF EXISTS idx_video_episode;
DROP INDEX IF EXISTS idx_video_movie;
DROP TABLE IF EXISTS video CASCADE;

--────────────────────────────────────
-- Episode table - Individual episodes
--────────────────────────────────────

DROP INDEX IF EXISTS idx_episode_thumbnail;
DROP INDEX IF EXISTS idx_episode_season_number;
DROP INDEX IF EXISTS idx_episode_season;
DROP TABLE IF EXISTS episode CASCADE;

--────────────────────────────────────
-- Season table - TV show seasons
--────────────────────────────────────

DROP INDEX IF EXISTS idx_season_tvshow_number;
DROP INDEX IF EXISTS idx_season_tvshow;
DROP TABLE IF EXISTS season CASCADE;

--────────────────────────────────────
-- TV Show table - TV show-specific data
--────────────────────────────────────

DROP INDEX IF EXISTS idx_tvshow_thumbnail;
DROP INDEX IF EXISTS idx_tvshow_content;
DROP TABLE IF EXISTS tv_show CASCADE;

--────────────────────────────────────
-- Movie table - Movie-specific data
--────────────────────────────────────

DROP INDEX IF EXISTS idx_movie_thumbnail;
DROP INDEX IF EXISTS idx_movie_content;
DROP TABLE IF EXISTS movie CASCADE;

--────────────────────────────────────
-- Thumbnail table - Image references
--────────────────────────────────────

DROP TABLE IF EXISTS thumbnail CASCADE;

--────────────────────────────────────
-- Content table - Base content info
--────────────────────────────────────

DROP INDEX IF EXISTS idx_content_updated_at;
DROP INDEX IF EXISTS idx_content_release_date;
DROP INDEX IF EXISTS idx_content_title;
DROP INDEX IF EXISTS idx_content_type;
DROP TABLE IF EXISTS content CASCADE;

--────────────────────────────────────
-- Drop custom types
--────────────────────────────────────

DROP TYPE IF EXISTS content_type_enum CASCADE;