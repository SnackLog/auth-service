TRUNCATE TABLE revoked_tokens;
TRUNCATE TABLE users;
INSERT INTO users (username, display_name, password_hash) VALUES
('foo', 'FOO', '$argon2id$v=19$m=65536,t=2,p=1$VkE0clVVL2ZaYlI1d05nV1U0cW9vdz09$8/0wJxQmno62Q5x7FhgX++zwSLGV8F+ySU9aNTLDj4Q'); -- Creds: User:foo PW:foo
