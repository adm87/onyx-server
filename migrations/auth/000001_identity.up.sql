CREATE TABLE auth.identities (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email           TEXT NOT NULL,
    password_hash   TEXT NOT NULL,
    auth_provider   TEXT NOT NULL DEFAULT 'password',
    mfa_enabled     BOOLEAN NOT NULL DEFAULT false,
    totp_secret     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX identities_email_idx ON auth.identities (email);

CREATE OR REPLACE FUNCTION auth.set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER identities_set_updated_at
    BEFORE UPDATE ON auth.identities
    FOR EACH ROW
    EXECUTE FUNCTION auth.set_updated_at();