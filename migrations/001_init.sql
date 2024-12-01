-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    first_name text,
    middle_name text,
    last_name text,
    description text,
    phone_number text,
    num_of_group text,
    vk_tag text,
    tg_tag text,
    is_admin boolean DEFAULT false,
    is_banned boolean DEFAULT false
);

CREATE TABLE IF NOT EXISTS groups (
    id serial PRIMARY KEY,
    name text,
    description text,
    is_academy boolean DEFAULT false,
    is_deletable boolean DEFAULT true
);


CREATE TABLE IF NOT EXISTS users_groups (
    user_id int REFERENCES users(id),
    group_id int REFERENCES groups(id),
    PRIMARY KEY (user_id, group_id)
);

CREATE TABLE IF NOT EXISTS reservations(
    id serial PRIMARY KEY,
    status text,
    group_id int,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    color text,
    place text,
    is_repeatable boolean DEFAULT false
);

CREATE OR REPLACE FUNCTION creating_user() RETURNS trigger AS $$
    BEGIN
        INSERT INTO "groups"("name", "description", "is_academy", "is_deletable") VALUES (CONCAT(NEW."first_name",' ', NEW."middle_name", ' ', NEW."last_name"),'', false, false);
        INSERT INTO "users_groups"("user_id","group_id") VALUES (NEW."user.id", NEW."group_id");
   RETURN NEW;
   END;
$$
language 'plpgsql';

CREATE TRIGGER new_record AFTER INSERT ON users
EXECUTE FUNCTION creating_user();

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE reservations;
DROP TABLE users_groups;
DROP TABLE groups;
DROP TABLE users;
-- +goose StatementEnd
