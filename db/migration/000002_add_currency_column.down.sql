ALTER TABLE "accounts" ADD COLUMN "country_code" int NOT NULL DEFAULT 0;
ALTER TABLE "accounts" DROP COLUMN "currency";