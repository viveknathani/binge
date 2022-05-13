create table if not exists videos(
    id uuid primary key,
    "videoId" varchar not null unique
);

create table if not exists shows(
    id uuid primary key,
    name varchar not null
);

create table if not exists episodes(
    id uuid primary key,
    "episodeNumber" int not null,
    name varchar not null,
    season int not null,
    "showId" uuid references shows("id") on delete cascade,
    "videoId" varchar references videos("videoId") on delete cascade
);

create table if not exists movies(
    id uuid primary key,
    name varchar not null,
    "videoId" varchar references videos("videoId") on delete cascade
);
