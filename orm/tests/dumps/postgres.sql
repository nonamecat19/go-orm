DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS roles CASCADE;

CREATE TABLE roles
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    name       TEXT NOT NULL
);

INSERT INTO roles (created_at, name)
VALUES ('2023-01-01 10:00:00', 'admin'),
       ('2023-01-01 10:00:00', 'moderator'),
       ('2023-01-01 10:00:00', 'user'),
       ('2023-01-01 10:00:00', 'owner');

CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    gender     VARCHAR(10)  NOT NULL,
    role_id    BIGINT,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE SET NULL
);

INSERT INTO users (created_at, name, email, gender, role_id)
VALUES ('2023-01-01 10:00:00', 'Alice Johnson', 'alice.johnson@example.com', 'female', 1),
       ('2023-01-01 10:00:00', 'Bob Smith', 'bob.smith@example.com', 'male', 3),
       ('2023-01-01 10:00:00', 'Charlie Brown', 'charlie.brown@example.com', 'male', 3),
       ('2023-01-01 10:00:00', 'Diana Prince', 'diana.prince@example.com', 'female', 3),
       ('2023-01-01 10:00:00', 'Edward King', 'edward.king@example.com', 'male', 3),
       ('2023-01-01 10:00:00', 'Fiona White', 'fiona.white@example.com', 'female', 2),
       ('2023-01-01 10:00:00', 'George Hall', 'george.hall@example.com', 'male', 3),
       ('2023-01-01 10:00:00', 'Hannah Wright', 'hannah.wright@example.com', 'female', 4),
       ('2023-01-01 10:00:00', 'Ivy Green', 'ivy.green@example.com', 'female', 2),
       ('2023-01-01 10:00:00', 'Jack Black', 'jack.black@example.com', 'male', 2),
       ('2023-01-01 10:00:00', 'Karen Hill', 'karen.hill@example.com', 'female', 3),
       ('2023-01-01 10:00:00', 'Liam Adams', 'liam.adams@example.com', 'male', 1),
       ('2023-01-01 10:00:00', 'Marie Clark', 'marie.clark@example.com', 'female', 3),
       ('2023-01-01 10:00:00', 'Nathan Bell', 'nathan.bell@example.com', 'male', 1),
       ('2023-01-01 10:00:00', 'Olivia Wood', 'olivia.wood@example.com', 'female', 3),
       ('2023-01-01 10:00:00', 'Patrick Moore', 'patrick.moore@example.com', 'male', 4),
       ('2023-01-01 10:00:00', 'Quinn Baker', 'quinn.baker@example.com', 'female', 4),
       ('2023-01-01 10:00:00', 'Ruby Fox', 'ruby.fox@example.com', 'female', 2),
       ('2023-01-01 10:00:00', 'Sam Hunter', 'sam.hunter@example.com', 'male', 3),
       ('2023-01-01 10:00:00', 'Tina Hall', 'tina.hall@example.com', 'female', 2);

CREATE TABLE orders
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    count      INT       NOT NULL,
    user_id    BIGINT,
    order_date TIMESTAMP NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);

INSERT INTO orders (created_at, count, user_id, order_date)
VALUES ('2023-01-01 10:00:00', 5, 1, '2023-01-01 10:00:00'),
       ('2023-01-01 10:00:00', 3, 2, '2023-01-02 12:00:00'),
       ('2023-01-01 10:00:00', 2, 3, '2023-01-03 14:00:00'),
       ('2023-01-01 10:00:00', 7, 4, '2023-01-04 16:00:00'),
       ('2023-01-01 10:00:00', 4, 5, '2023-01-05 18:00:00'),
       ('2023-01-01 10:00:00', 10, 6, '2023-01-06 20:00:00'),
       ('2023-01-01 10:00:00', 8, NULL, '2023-01-07 22:00:00'),
       ('2023-01-01 10:00:00', 6, 8, '2023-01-08 08:00:00'),
       ('2023-01-01 10:00:00', 9, 9, '2023-01-09 09:00:00'),
       ('2023-01-01 10:00:00', 1, NULL, '2023-01-10 10:00:00'),
       ('2023-01-01 10:00:00', 15, 11, '2023-01-11 11:00:00'),
       ('2023-01-01 10:00:00', 12, 12, '2023-01-12 12:00:00'),
       ('2023-01-01 10:00:00', 20, 13, '2023-01-13 13:00:00'),
       ('2023-01-01 10:00:00', 18, 14, '2023-01-14 14:00:00'),
       ('2023-01-01 10:00:00', 25, 15, '2023-01-15 15:00:00'),
       ('2023-01-01 10:00:00', 30, 16, '2023-01-16 16:00:00'),
       ('2023-01-01 10:00:00', 22, 17, '2023-01-17 17:00:00'),
       ('2023-01-01 10:00:00', 17, 18, '2023-01-18 18:00:00'),
       ('2023-01-01 10:00:00', 13, 19, '2023-01-19 19:00:00'),
       ('2023-01-01 10:00:00', 8, 20, '2023-01-20 20:00:00');
