CREATE TABLE "Employees" (
  "id" serial PRIMARY KEY,
  "first_name" varchar(50) NOT NULL,
  "last_name" varchar(50) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "phone" varchar(20) UNIQUE NOT NULL,
  "address" varchar(255) NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role_id" int NOT NULL,
  "is_admin" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Roles" (
  "id" serial PRIMARY KEY,
  "name" varchar(50) NOT NULL
);

CREATE TABLE "Shifts" (
  "id" serial PRIMARY KEY,
  "employee_id" int NOT NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL
);

CREATE TABLE "ClockIns" (
  "id" serial PRIMARY KEY,
  "employee_id" int NOT NULL,
  "clock_in_time" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ClockOuts" (
  "id" serial PRIMARY KEY,
  "employee_id" int NOT NULL,
  "clock_in_id" serial NOT NULL,
  "clock_out_time" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Shifts" ("employee_id");

CREATE INDEX ON "ClockIns" ("employee_id");

CREATE INDEX ON "ClockOuts" ("employee_id");

ALTER TABLE "Employees" ADD FOREIGN KEY ("role_id") REFERENCES "Roles" ("id");

ALTER TABLE "Shifts" ADD FOREIGN KEY ("employee_id") REFERENCES "Employees" ("id");

ALTER TABLE "ClockIns" ADD FOREIGN KEY ("employee_id") REFERENCES "Employees" ("id");

ALTER TABLE "ClockOuts" ADD FOREIGN KEY ("employee_id") REFERENCES "Employees" ("id");

ALTER TABLE "ClockOuts" ADD FOREIGN KEY ("clock_in_id") REFERENCES "ClockIns" ("id");
