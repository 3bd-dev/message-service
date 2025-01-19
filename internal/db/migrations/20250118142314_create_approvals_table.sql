-- migrate:up
CREATE TYPE message_review_action AS ENUM ('approved', 'rejected');

CREATE TABLE message_reviews (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id UUID REFERENCES messages(id) ON DELETE CASCADE,
    reviewer_id UUID REFERENCES users(id) ON DELETE CASCADE,
    action message_review_action NOT NULL, -- Use the enum type here
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE IF EXISTS message_reviews;
DROP TYPE IF EXISTS message_review_action;