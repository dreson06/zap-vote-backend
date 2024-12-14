CREATE TABLE _special_election
(
    id           VARCHAR(28) PRIMARY KEY,
    course_code VARCHAR(8)  NOT NULL,
    candidate_id      VARCHAR(28)  NOT NULL,
    slogan       VARCHAR(150) NOT NULL,
    votes        INTEGER      NOT NULL DEFAULT 0,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE _faculty_election
(
    id           VARCHAR(28) PRIMARY KEY,
    name VARCHAR(15)  NOT NULL,
    candidate_id      VARCHAR(28)  NOT NULL,
    slogan       VARCHAR(150) NOT NULL,
    votes        INTEGER      NOT NULL DEFAULT 0,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)