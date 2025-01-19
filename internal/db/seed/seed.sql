INSERT INTO users (id, username, role) VALUES
    ('550e8400-e29b-41d4-a716-446655440000', 'maker1', 'maker')
ON CONFLICT (username) DO NOTHING;

INSERT INTO users (id, username, role) VALUES
    ('550e8400-e29b-41d4-a716-446655440001', 'checker1', 'checker')
ON CONFLICT (username) DO NOTHING;

INSERT INTO users (id, username, role) VALUES
    ('550e8400-e29b-41d4-a716-446655440002', 'checker2', 'checker')
ON CONFLICT (username) DO NOTHING;
