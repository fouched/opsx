CREATE TABLE remote_services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(128) NOT NULL UNIQUE,
    type VARCHAR(128) NOT NULL,
    host VARCHAR(128) NOT NULL,
    port INT NOT NULL,
    status VARCHAR(128) NOT NULL DEFAULT 'Down',
    health_url VARCHAR(128),
    command_start TEXT NOT NULL,
    command_stop TEXT NOT NULL,
    command_restart TEXT NOT NULL,
    last_checked TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    UNIQUE (host, port)
);