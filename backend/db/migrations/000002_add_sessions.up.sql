CREATE TABLE "Sessions" (
  "id" uuid PRIMARY KEY,
  "email" varchar(255) NOT NULL,
  "employee_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "Sessions" ADD FOREIGN KEY ("employee_id") REFERENCES "Employees" ("id");
