DROP TRIGGER IF EXISTS users_set_updated_at ON users.users;
DROP FUNCTION IF EXISTS users.set_updated_at();
DROP TABLE IF EXISTS users.users;