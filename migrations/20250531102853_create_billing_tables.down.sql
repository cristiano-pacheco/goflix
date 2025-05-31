--────────────────────────────────────
-- 1. Drop tables (in reverse order of creation)
--────────────────────────────────────

-- Drop subscription table first (has foreign key dependency)
DROP TABLE IF EXISTS subscription;

-- Drop plan table
DROP TABLE IF EXISTS plan;

--────────────────────────────────────
-- 2. Drop enums
--────────────────────────────────────

DROP TYPE IF EXISTS subscription_status_enum;
DROP TYPE IF EXISTS plan_interval_enum;