CREATE TABLE IF NOT EXISTS _votes(
    user_id TEXT NOT NULL DEFAULT '',
    vote_type SMALLINT CHECK(vote_type > 0),
    UNIQUE(user_id,vote_type)
)