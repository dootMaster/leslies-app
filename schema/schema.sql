--D:\leslies-app\db\schema.sql;
-- DROP DATABASE IF EXISTS lesliesapp;
-- CREATE DATABASE lesliesapp;

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    pw VARCHAR(255) NOT NULL,
    salt VARCHAR(255) NOT NULL,
);

DROP TABLE IF EXISTS user_session CASCADE;
CREATE TABLE user_session (
    user_id INT NOT NULL,
    session_key VARCHAR(255) UNIQUE NOT NULL,
    date_created TIMESTAMP WITH TIME ZONE DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE INDEX idx_session_key ON user_session (session_key);

DROP TABLE IF EXISTS global_song_list CASCADE;
CREATE TABLE global_song_list (
    global_song_id VARCHAR(255) PRIMARY KEY NOT NULL,
    global_title VARCHAR(255) NOT NULL,
    global_artist VARCHAR(255) NOT NULL
);

CREATE INDEX idx_song_id ON global_song_list (global_song_id);
CREATE INDEX idx_title ON global_song_list (global_title);
CREATE INDEX idx_artist ON global_song_list (global_artist);

DROP TABLE IF EXISTS main_list CASCADE;
CREATE TABLE main_list (
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    song_id VARCHAR(255) PRIMARY KEY,
    FOREIGN KEY (song_id) REFERENCES global_song_list (global_song_id),
    title VARCHAR(255),
    artist VARCHAR(255),
    main_list_resource VARCHAR(255),
    main_list_subtext VARCHAR(255)
);

CREATE INDEX idx_main_list_user ON main_list (user_id);

DROP TABLE IF EXISTS setlists CASCADE;
CREATE TABLE setlists (
    setlist_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    setlist_title VARCHAR(255) NOT NULL,
    setlist_title_subtext VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS setlist_songs CASCADE;
CREATE TABLE setlist_songs (
    id SERIAL PRIMARY KEY,
    setlist_id INT NOT NULL,
    song_id VARCHAR(255),
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255),
    song_order INT NOT NULL,
    setlist_subtext VARCHAR(255),
    FOREIGN KEY (setlist_id) REFERENCES setlists (setlist_id) ON DELETE CASCADE,
    FOREIGN KEY (song_id) REFERENCES main_list (song_id) ON DELETE CASCADE
);