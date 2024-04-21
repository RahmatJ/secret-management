DO $$
BEGIN
    IF EXISTS(SELECT *
              FROM information_schema.columns
              WHERE table_name='secret_management' and column_name='expireddate')
    THEN
        alter table secret_management rename column expireddate to expired_date;
    END IF;

-- Create a temporary TIMESTAMP column
    alter table secret_management ADD COLUMN create_time_holder timestamp without time zone NULL;

-- Copy casted value over to the temporary column
    update secret_management set create_time_holder = expired_date::TIMESTAMP;

-- Modify original column using the temporary column
    alter table secret_management alter column expired_date type TIMESTAMP without time zone USING create_time_holder;

-- Drop the temporary column (after examining altered column values)
    ALTER TABLE secret_management DROP COLUMN create_time_holder;
END $$;