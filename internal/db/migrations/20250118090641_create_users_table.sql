-- migrate:up
-- Create enum type for user roles
CREATE TYPE user_role AS ENUM ('admin', 'checker', 'maker');

-- Create users table
CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) UNIQUE NOT NULL,
    role user_role NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS user_role;