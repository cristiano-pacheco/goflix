--────────────────────────────────────
-- 1. Enums
--────────────────────────────────────

CREATE TYPE content_type_enum AS ENUM ('MOVIE', 'TV_SHOW');

--────────────────────────────────────
-- Content table - Base content info
--────────────────────────────────────

CREATE TABLE content (
    id BIGSERIAL PRIMARY KEY,
    type                content_type_enum NOT NULL,
    title               TEXT   NOT NULL,
    description         TEXT   NOT NULL,
    age_recommendation  SMALLINT,
    release_date        DATE,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Indexes for content table
CREATE INDEX idx_content_type ON content(type);
CREATE INDEX idx_content_title ON content(title);
CREATE INDEX idx_content_release_date ON content(release_date);
CREATE INDEX idx_content_updated_at ON content(updated_at);

--────────────────────────────────────
-- Thumbnail table - Image references
--────────────────────────────────────

CREATE TABLE thumbnail (
    id BIGSERIAL PRIMARY KEY,
    url        TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

--────────────────────────────────────
-- Movie table - Movie-specific data
--────────────────────────────────────

CREATE TABLE movie (
    id BIGSERIAL PRIMARY KEY,
    external_rating DOUBLE PRECISION,
    content_id   BIGINT NOT NULL REFERENCES content(id) ON DELETE CASCADE,
    thumbnail_id BIGINT REFERENCES thumbnail(id) ON DELETE SET NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (content_id)
);

-- Indexes for movie table
CREATE INDEX idx_movie_content ON movie(content_id);
CREATE INDEX idx_movie_thumbnail ON movie(thumbnail_id);

--────────────────────────────────────
-- TV Show table - TV show-specific data
--────────────────────────────────────

CREATE TABLE tv_show (
    id BIGSERIAL PRIMARY KEY,
    content_id   BIGINT NOT NULL REFERENCES content(id) ON DELETE CASCADE,
    thumbnail_id BIGINT REFERENCES thumbnail(id) ON DELETE SET NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (content_id)
);

-- Indexes for tv_show table
CREATE INDEX idx_tvshow_content ON tv_show(content_id);
CREATE INDEX idx_tvshow_thumbnail ON tv_show(thumbnail_id);

--────────────────────────────────────
-- Season table - TV show seasons
--────────────────────────────────────

CREATE TABLE season (
    id BIGSERIAL PRIMARY KEY,
    tv_show_id    BIGINT NOT NULL REFERENCES tv_show(id) ON DELETE CASCADE,
    season_number SMALLINT NOT NULL,
    title         TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (tv_show_id, season_number)
);

-- Indexes for season table
CREATE INDEX idx_season_tvshow ON season(tv_show_id);
CREATE INDEX idx_season_tvshow_number ON season(tv_show_id, season_number);

--────────────────────────────────────
-- Episode table - Individual episodes
--────────────────────────────────────

CREATE TABLE episode (
    id BIGSERIAL PRIMARY KEY,
    season_id      BIGINT NOT NULL REFERENCES season(id) ON DELETE CASCADE,
    title          TEXT   NOT NULL,
    description    TEXT   NOT NULL,
    episode_number SMALLINT NOT NULL,
    thumbnail_id   BIGINT REFERENCES thumbnail(id) ON DELETE SET NULL,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (season_id, episode_number)
);

-- Indexes for episode table
CREATE INDEX idx_episode_season ON episode(season_id);
CREATE INDEX idx_episode_season_number ON episode(season_id, episode_number);
CREATE INDEX idx_episode_thumbnail ON episode(thumbnail_id);

--────────────────────────────────────
-- Video table - Video files and URLs
--────────────────────────────────────

CREATE TABLE video (
    id BIGSERIAL PRIMARY KEY,
    url         TEXT NOT NULL,
    size_in_kb  BIGINT,
    duration    INT,
    movie_id    BIGINT REFERENCES movie(id) ON DELETE CASCADE,
    episode_id  BIGINT REFERENCES episode(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (movie_id),
    UNIQUE (episode_id),
    -- Ensure a video belongs to either a movie OR an episode, not both
    CONSTRAINT video_belongs_to_one_content CHECK (
        (movie_id IS NOT NULL AND episode_id IS NULL) OR 
        (movie_id IS NULL AND episode_id IS NOT NULL)
    )
);

-- Indexes for video table
CREATE INDEX idx_video_movie ON video(movie_id);
CREATE INDEX idx_video_episode ON video(episode_id);
CREATE INDEX idx_video_updated_at ON video(updated_at);

--────────────────────────────────────
-- Video Metadata table
--────────────────────────────────────

CREATE TABLE video_metadata (
    id BIGSERIAL PRIMARY KEY,
    auto_generated_description TEXT,
    transcript                 TEXT,
    age_rating                 SMALLINT,
    age_rating_explanation     TEXT,
    video_id BIGINT NOT NULL REFERENCES video(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (video_id)
);

-- Indexes for video_metadata table
CREATE INDEX idx_video_metadata_video ON video_metadata(video_id);

--────────────────────────────────────
-- Video Age Rating Category - Content categories
--────────────────────────────────────

CREATE TABLE video_age_rating_category (
    video_id BIGINT NOT NULL REFERENCES video(id) ON DELETE CASCADE,
    category TEXT   NOT NULL,
    PRIMARY KEY (video_id, category)
);