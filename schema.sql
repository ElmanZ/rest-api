create table "user" (
  user_id BIGSERIAL NOT NULL PRIMARY KEY,
  username VARCHAR(250) NOT NULL,
  created_at DATE default now()
);

create table chat (
  chat_id BIGSERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  users int REFERENCES "user" (user_id),
  created_at DATE default now()
);