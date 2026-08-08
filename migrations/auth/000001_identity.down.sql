DROP TRIGGER IF EXISTS identities_set_updated_at ON auth.identities;
DROP FUNCTION IF EXISTS auth.set_updated_at();
DROP TABLE IF EXISTS auth.identities;