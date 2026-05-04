CREATE SCHEMA rep;

CREATE TABLE rep.users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(16) NOT NULL UNIQUE,
    password varchar(64) NOT NULL,
    first_name varchar(32) NOT NULL,
    last_name varchar(32) NOT NULL,
    address text,
    email varchar(32),
    phone_number varchar(16) CHECK (
        phone_number ~ '^\+[0-9]+$' AND 
        char_length(phone_number) BETWEEN 11 AND 16 
    ),
    created_at date NOT NULL DEFAULT CURRENT_DATE,
    updated_at timestamp,
    rating smallint NOT NULL DEFAULT 1000,
    role varchar(16) DEFAULT 'user',
    image_url text
);

CREATE OR REPLACE FUNCTION rep.set_updated_at() 
RETURNS TRIGGER AS $$ 
BEGIN
    NEW.updated_at := NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_update_users 
BEFORE UPDATE
    ON rep.users FOR EACH ROW 
    EXECUTE FUNCTION rep.set_updated_at();

CREATE TABLE rep.ivents (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(32) NOT NULL DEFAULT 'ivent',
    description text,
    address varchar(80),
    datetime timestamp NOT NULL DEFAULT TIMESTAMP '2002-11-21 11:20:00',
    price int NOT NULL DEFAULT 0,
    created_at date NOT NULL DEFAULT CURRENT_DATE,
    updated_at date,
    creator varchar(16) NOT NULL REFERENCES rep.users(username) ON UPDATE CASCADE,
    updater varchar(16) REFERENCES rep.users(username) ON UPDATE CASCADE,
    image_url text
);

CREATE TRIGGER before_update_ivents 
BEFORE UPDATE
    ON rep.ivents FOR EACH ROW 
    EXECUTE FUNCTION rep.set_updated_at();

CREATE TYPE rep.game_type AS ENUM('online', 'offline');

CREATE TYPE rep.game_result AS ENUM ('win_w', 'win_b', 'draw');

CREATE TABLE rep.games (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    fen_start varchar(64) NOT NULL DEFAULT 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1',
    fen_end varchar(64) NOT NULL DEFAULT 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1',
    created_at date NOT NULL DEFAULT CURRENT_DATE,
    player_w varchar(16) NOT NULL REFERENCES rep.users(username) ON UPDATE CASCADE,
    player_b varchar(16) NOT NULL REFERENCES rep.users(username) ON UPDATE CASCADE,
    type rep.game_type NOT NULL DEFAULT 'online',
    mode varchar(16) NOT NULL,
    result rep.game_result NOT NULL,
    history json
);

CREATE OR REPLACE FUNCTION uuid_or_null(str text)
RETURNS uuid AS $$
BEGIN
  RETURN str::uuid;
EXCEPTION WHEN invalid_text_representation THEN
  RETURN NULL;
END;
$$ LANGUAGE plpgsql;