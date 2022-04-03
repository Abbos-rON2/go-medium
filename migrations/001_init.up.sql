create table if not exists users (
  id serial primary key,
  name varchar(255) not null,
  email varchar(255) unique not null,
  password varchar(255) not null,
  created_at timestamp default now(),
  updated_at timestamp default now()
);

create table if not exists subscriptions (
  id serial primary key,
  user_id int not null references users(id),
  subscriber_id int not null references users(id),
  created_at timestamp default now(),
  updated_at timestamp default now()
);

alter table subscriptions add constraint unique_subcriptions unique (user_id, subscriber_id);

create table if not exists posts (
  id serial primary key,
  title varchar(255) not null,
  content text not null,
  author_id int not null references users(id),
  created_at timestamp default now(),
  updated_at timestamp default now()
);

create table if not exists likes (
  id serial primary key,
  user_id int not null references users(id),
  post_id int not null references posts(id),
  created_at timestamp default now()
);

alter table likes add constraint unique_likes unique (user_id, post_id);

create table if not exists comments (
  id serial primary key,
  content text not null,
  reply_id int references comments(id),
  author_id int not null references users(id),
  post_id int not null references posts(id),
  created_at timestamp default now()
);

create table if not exists chats (
  id serial primary key,
  user_id int not null references users(id),
  content text not null,
  created_at timestamp default now()
);

create table if not exists chat_users (
  chat_id int not null references chats(id),
  user_id int not null references users(id)
);

create table if not exists messages (
  id serial primary key,
  chat_id int not null references chats(id),
  user_id int not null references users(id),
  content text not null,
  created_at timestamp default now(),
  updated_at timestamp default now()
);