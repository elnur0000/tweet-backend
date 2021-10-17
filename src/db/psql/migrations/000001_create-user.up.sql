CREATE TABLE "user" (
 id bigserial primary key,
 created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
 avatar text,
 email citext UNIQUE NOT NULL,
 password text NOT NULL,
 version integer NOT NULL DEFAULT 1
);