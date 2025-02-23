CREATE TABLE settlements (
    id SERIAL PRIMARY KEY,
    owner VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    survival_limit INTEGER NOT NULL,
    departing_survival INTEGER NOT NULL,
    collective_cog INTEGER NOT NULL,
    year INTEGER NOT NULL
);

CREATE TABLE survivors (
    id SERIAL PRIMARY KEY,
    settlement_id INTEGER REFERENCES settlements(id),
    name VARCHAR(255) NOT NULL,
    birth INTEGER NOT NULL,
    gender VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    hunt_xp INTEGER NOT NULL,
    survival INTEGER NOT NULL,
    movement INTEGER NOT NULL,
    accuracy INTEGER NOT NULL,
    strength INTEGER NOT NULL,
    evasion INTEGER NOT NULL,
    luck INTEGER NOT NULL,
    speed INTEGER NOT NULL,
    insanity INTEGER NOT NULL,
    systemic_pressure INTEGER NOT NULL,
    torment INTEGER NOT NULL,
    lumi INTEGER NOT NULL,
    courage INTEGER NOT NULL,
    understanding INTEGER NOT NULL
);

-- Add indexes for common queries
CREATE INDEX idx_survivors_settlement ON survivors(settlement_id);
CREATE INDEX idx_settlements_owner ON settlements(owner);