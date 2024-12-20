CREATE TABLE users(
                      id uuid PRIMARY KEY,
                      nickname varchar(70) not null,
                      name varchar(100) not null,
                      email varchar(120) unique not null,
                      verified boolean default false
);
CREATE UNIQUE INDEX idx_users_email ON users (email);
CREATE TABLE email_tokens(
                             id SERIAL primary key ,
                             email varchar(120) unique REFERENCES  users(email),
                             token varchar(255) unique not null,
                             expires_at timestamp default  current_timestamp+INTERVAL '5 hours',
                             created_at timestamp default current_timestamp
);
CREATE UNIQUE INDEX idx_email_tokens_token ON email_tokens (token);
