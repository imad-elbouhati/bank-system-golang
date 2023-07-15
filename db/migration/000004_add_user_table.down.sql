-- Drop the foreign key constraint on "accounts" table
ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drop the unique index on "accounts" table
DROP INDEX IF EXISTS "accounts_owner_currency_idx";

-- Drop the "users" table
DROP TABLE IF EXISTS "users";
