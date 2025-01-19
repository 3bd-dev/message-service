INSERT INTO users (id, username, role) VALUES
    (uuid_generate_v4(), 'checker1', 'checker')
ON CONFLICT (username) DO NOTHING;

INSERT INTO users (id, username, role) VALUES
    (uuid_generate_v4(), 'checker2', 'checker')
ON CONFLICT (username) DO NOTHING;

INSERT INTO users (id, username, role) VALUES
    (uuid_generate_v4(), 'maker1', 'maker')
ON CONFLICT (username) DO NOTHING;