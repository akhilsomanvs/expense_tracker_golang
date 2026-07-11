DROP TRIGGER IF EXISTS users_updated_at ON users;
DROP TRIGGER IF EXISTS expenses_updated_at ON expenses;
DROP FUNCTION IF EXISTS update_updated_at_column();