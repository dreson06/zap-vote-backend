CREATE TABLE IF NOT EXISTS _classrep(
    id TEXT NOT NULL PRIMARY KEY,
    course_code TEXT NOT NULL DEFAULT '',
    candidate_id TEXT NOT NULL DEFAULT '',
    slogan TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)