CREATE TABLE _candidate
(
    id          VARCHAR(28) PRIMARY KEY,
    name        VARCHAR(70) NOT NULL DEFAULT '',
    course_code VARCHAR(10)  NOT NULL DEFAULT '',
    position VARCHAR(50) NOT NULL DEFAULT '',
    type VARCHAR(50) NOT NULL DEFAULT '',
    department VARCHAR(50)NOT NULL DEFAULT '',
    thumbnail TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)