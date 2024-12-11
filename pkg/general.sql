--- this is a simple script to automate adding a trigger function for each table with updated_at
--- so that on every update it should automatically set updated_at to now
CREATE FUNCTION update_updated_at() RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$;

--- this script automates adding the trigger to all tables that contains column 'updated_at'
--- it creates the trigger only once for each table, tables with the trigger won't be affected
DO $$
    DECLARE
        table_name text;
    BEGIN
        FOR table_name IN (
            SELECT DISTINCT c.table_name
            FROM information_schema.columns c
                     LEFT JOIN pg_trigger t ON t.tgrelid = c.table_name::regclass
                AND t.tgname = 'update_updated_at_trigger'
            WHERE c.column_name = 'updated_at'
              AND t.tgname IS NULL
        )
            LOOP
                EXECUTE format('CREATE TRIGGER update_updated_at_trigger
                        BEFORE UPDATE ON %I FOR EACH ROW
                        EXECUTE FUNCTION update_updated_at();', table_name);
            END LOOP;
    END $$;
