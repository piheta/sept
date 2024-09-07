-- Users Table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL UNIQUE, -- Unique identifier for the user, e.g., a UUID or username
    username TEXT NOT NULL,     -- Display name of the user
    ip TEXT,                    -- IP address of the user
    avatar TEXT                 -- URL or path to the user's avatar
);

-- Chats Table
CREATE TABLE chats (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,                  -- Optional name for the chat (could be empty for private chats)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User Chats Table (many-to-many relationship)
CREATE TABLE user_chats (
    user_id INTEGER,
    chat_id INTEGER,
    PRIMARY KEY (user_id, chat_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (chat_id) REFERENCES chats(id)
);

-- Messages Table
CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INTEGER,
    user_id INTEGER,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES chats(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
