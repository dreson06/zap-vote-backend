CREATE TABLE _user(
    id VARCHAR(15) PRIMARY KEY,
    device_id TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL DEFAULT '',
    course_code VARCHAR(10) NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);