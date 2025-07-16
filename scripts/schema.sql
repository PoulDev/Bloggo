CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title char(128) NOT NULL,
    content TEXT NOT NULL,
    contentRaw TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    comment text NOT NULL,
    created_at TIMESTAMP NOT NULL,
    post INTEGER NOT NULL,
    FOREIGN KEY (post) REFERENCES posts(id)
);

CREATE TABLE IF NOT EXISTS post_authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post INTEGER NOT NULL,
    author INTEGER NOT NULL,
    FOREIGN KEY (post) REFERENCES posts(id),
    FOREIGN KEY (author) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name char(32) NOT NULL,
    picture char(512) NOT NULL,
    role SMALLINT NOT NULL DEFAULT 0, -- 0 = author, 1 = admin
    pwd BINARY(64) NOT NULL,
    salt BINARY(32) NOT NULL
);
