CREATE TABLE "fizzbuzz_statistics" (
  "request" jsonb PRIMARY KEY,
  "count" bigint NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
