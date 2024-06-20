CREATE TABLE "users" (
  "user_id" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "middle_name" varchar,
  "last_name" varchar NOT NULL,
  "profile_picture" varchar,
  "phone_number" varchar UNIQUE,
  "currency" varchar NOT NULL,
  "is_verified" boolean NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT current_timestamp,
  "updated_at" timestamp NOT NULL DEFAULT current_timestamp
);

CREATE INDEX ON "users" ("user_id");
