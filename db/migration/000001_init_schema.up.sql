CREATE TABLE "fizzbuzz_statistics" (
  "request" jsonb PRIMARY KEY,
  "count" int,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
