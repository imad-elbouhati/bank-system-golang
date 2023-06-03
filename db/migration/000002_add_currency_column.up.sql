ALTER TABLE "accounts" ADD COLUMN "currency" varchar NOT NULL DEFAULT '';
ALTER TABLE "accounts" DROP COLUMN "country_code";