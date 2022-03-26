-- +migrate Up
CREATE TABLE wss_tokens (
    token uuid NOT NULL DEFAULT uuid_generate_v4 (),
    valid_until timestamptz NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT wss_tokens_pkey PRIMARY KEY (token)
);

CREATE INDEX idx_wss_tokens_fk_user_uid ON wss_tokens USING btree (user_id);

ALTER TABLE wss_tokens
    ADD CONSTRAINT wss_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS wss_tokens;

