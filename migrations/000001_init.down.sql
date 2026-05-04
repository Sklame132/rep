DROP FUNCTION rep.uuid_or_null ();

DROP TABLE rep.games;

DROP TYPE rep.game_result;

DROP TYPE rep.game_type;

DROP TRIGGER before_update_ivents ON rep.ivents;

DROP TABLE rep.ivents;

DROP TRIGGER before_update_users ON rep.users;

DROP FUNCTION rep.set_updated_at ();

DROP TABLE rep.users;

DROP SCHEMA rep;