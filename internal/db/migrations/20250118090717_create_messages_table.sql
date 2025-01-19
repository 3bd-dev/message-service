-- migrate:up
CREATE TYPE message_status AS ENUM ('pending', 'approved', 'rejected');

CREATE TABLE messages (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    content TEXT NOT NULL,
    status message_status DEFAULT 'pending',
    recipient VARCHAR(255) NOT NULL,
    sender_id UUID REFERENCES users(id) ON DELETE CASCADE,
    approvals_required INT DEFAULT 2,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE IF EXISTS messages;
DROP TYPE IF EXISTS message_status;