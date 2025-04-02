CREATE TABLE users
(
    id         INT IDENTITY(1,1) PRIMARY KEY,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    name       NVARCHAR(255) NOT NULL,
    email      NVARCHAR(255) NOT NULL UNIQUE,
    gender     NVARCHAR(10)  NOT NULL
);

INSERT INTO users (name, email, gender)
VALUES ('Alice Johnson', 'alice.johnson@example.com', 'female'),
       ('Bob Smith', 'bob.smith@example.com', 'male'),
       ('Charlie Brown', 'charlie.brown@example.com', 'male'),
       ('Diana Prince', 'diana.prince@example.com', 'female'),
       ('Edward King', 'edward.king@example.com', 'male'),
       ('Fiona White', 'fiona.white@example.com', 'female'),
       ('George Hall', 'george.hall@example.com', 'male'),
       ('Hannah Wright', 'hannah.wright@example.com', 'female'),
       ('Ivy Green', 'ivy.green@example.com', 'female'),
       ('Jack Black', 'jack.black@example.com', 'male'),
       ('Karen Hill', 'karen.hill@example.com', 'female'),
       ('Liam Adams', 'liam.adams@example.com', 'male'),
       ('Marie Clark', 'marie.clark@example.com', 'female'),
       ('Nathan Bell', 'nathan.bell@example.com', 'male'),
       ('Olivia Wood', 'olivia.wood@example.com', 'female'),
       ('Patrick Moore', 'patrick.moore@example.com', 'male'),
       ('Quinn Baker', 'quinn.baker@example.com', 'female'),
       ('Ruby Fox', 'ruby.fox@example.com', 'female'),
       ('Sam Hunter', 'sam.hunter@example.com', 'male'),
       ('Tina Hall', 'tina.hall@example.com', 'female');

CREATE TABLE orders
(
    id         INT IDENTITY(1,1) PRIMARY KEY,
    count      INT NOT NULL,
    user_id    INT NULL,
    order_date DATETIME NOT NULL,
    created_at DATETIME DEFAULT GETDATE(),
    deleted_at DATETIME NULL,
    updated_at DATETIME NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO orders (count, user_id, order_date)
VALUES (5, 1, '2023-01-01 10:00:00'),
       (3, 2, '2023-01-02 12:00:00'),
       (2, 3, '2023-01-03 14:00:00'),
       (7, 4, '2023-01-04 16:00:00'),
       (4, 5, '2023-01-05 18:00:00'),
       (10, 6, '2023-01-06 20:00:00'),
       (8, NULL, '2023-01-07 22:00:00'),
       (6, 8, '2023-01-08 08:00:00'),
       (9, 9, '2023-01-09 09:00:00'),
       (1, NULL, '2023-01-10 10:00:00'),
       (15, 11, '2023-01-11 11:00:00'),
       (12, 12, '2023-01-12 12:00:00'),
       (20, 13, '2023-01-13 13:00:00'),
       (18, 14, '2023-01-14 14:00:00'),
       (25, 15, '2023-01-15 15:00:00'),
       (30, 16, '2023-01-16 16:00:00'),
       (22, 17, '2023-01-17 17:00:00'),
       (17, 18, '2023-01-18 18:00:00'),
       (13, 19, '2023-01-19 19:00:00'),
       (8, 20, '2023-01-20 20:00:00');
