alter table if exists posts add  column if not exists tags text[] not null default '{}';