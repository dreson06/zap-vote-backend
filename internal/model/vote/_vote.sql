CREATE TABLE _votes(
    user_id TEXT NOT NULL DEFAULT '',
    election_id TEXT NOT NULL DEFAULT '',
    UNIQUE(user_id,election_id)
);

CREATE OR REPLACE FUNCTION increment_total_votes()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE _election
    SET total_votes = total_votes + 1
    WHERE id = NEW.election_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Then create the trigger that uses this function
CREATE TRIGGER increment_votes_trigger
    AFTER INSERT ON _votes
    FOR EACH ROW
EXECUTE FUNCTION increment_total_votes();