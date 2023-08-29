DROP TABLE IF EXISTS "public"."app_owner";
CREATE TABLE "public"."app_owner" (
    "app_id" uuid NOT NULL,
    "user_id" uuid NOT NULL
);

DROP TABLE IF EXISTS "public"."applications";
CREATE TABLE "public"."applications" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" text NOT NULL,
    "description" text NOT NULL,
    "api_key" text NOT NULL DEFAULT gen_random_uuid(),
    "is_valid" bool NOT NULL
);

DROP TABLE IF EXISTS "public"."buildings";
CREATE TABLE "public"."buildings" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "long_name" text,
    "short_name" text
);

DROP TABLE IF EXISTS "public"."elon_ods";
CREATE TABLE "public"."elon_ods" (
    "elon_id" text NOT NULL,
    "ods_id" uuid NOT NULL
);

DROP TABLE IF EXISTS "public"."users";

DROP TYPE IF EXISTS "public"."affiliation";
CREATE TYPE "public"."affiliation" AS ENUM ('Staff', 'Alumni', 'Student', 'Affiliate');
CREATE TABLE "public"."users" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "given_name" text NOT NULL,
    "family_name" text NOT NULL,
    "email" text NOT NULL,
    "affiliation" "public"."affiliation" NOT NULL
);

CREATE UNIQUE INDEX "api_key_is_unique" ON "applications" USING BTREE ("api_key");

CREATE UNIQUE INDEX "id_is_unique" ON "applications" USING BTREE ("id");

CREATE UNIQUE INDEX "elon_id_is_unique" ON "elon_ods" USING BTREE ("elon_id");

CREATE UNIQUE INDEX "ods_id_is_unique" ON "elon_ods" USING BTREE ("ods_id");

CREATE UNIQUE INDEX "user_id_is_unique" ON "users" USING BTREE ("id");

ALTER TABLE "public"."app_owner" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id");
ALTER TABLE "public"."app_owner" ADD FOREIGN KEY ("app_id") REFERENCES "public"."applications"("id");
ALTER TABLE "public"."elon_ods" ADD FOREIGN KEY ("ods_id") REFERENCES "public"."users"("id");
