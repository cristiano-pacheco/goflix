--────────────────────────────────────
-- 1. Enums
--────────────────────────────────────

CREATE TYPE plan_interval_enum AS ENUM ('Day', 'Week', 'Month', 'Year');
CREATE TYPE subscription_status_enum AS ENUM ('Active', 'Inactive', 'Cancelled', 'Expired', 'PastDue');

--────────────────────────────────────
-- 2. Subscription system tables
--────────────────────────────────────

--────────────────────────────────────
-- Plan table - Subscription plans
--────────────────────────────────────

CREATE TABLE plan (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(255),
    amount_cents INTEGER NOT NULL,
    currency VARCHAR(3) NOT NULL,
    interval plan_interval_enum NOT NULL,
    trial_period INTEGER,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

--────────────────────────────────────
-- Subscription table - User subscriptions
--────────────────────────────────────

CREATE TABLE subscription (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    plan_id BIGINT NOT NULL REFERENCES plan(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    status subscription_status_enum NOT NULL DEFAULT 'Inactive',
    start_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    end_date TIMESTAMPTZ,
    auto_renew BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);