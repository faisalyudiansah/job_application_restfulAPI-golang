DROP TABLE IF EXISTS Users CASCADE;
DROP TABLE IF EXISTS UserProfiles CASCADE;
DROP TABLE IF EXISTS Jobs CASCADE;
DROP TABLE IF EXISTS JobApplicants CASCADE;

CREATE TABLE Users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,  
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
  	updated_at TIMESTAMP NOT NULL,
  	deleted_at TIMESTAMP
);

CREATE TABLE UserProfiles (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE NOT NULL,  
    age INT NOT NULL,
    current_job VARCHAR,
    address TEXT,
    created_at TIMESTAMP NOT NULL,
  	updated_at TIMESTAMP NOT NULL,
  	deleted_at TIMESTAMP
);

CREATE TABLE Jobs (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,  
    company VARCHAR NOT NULL,
    is_open BOOLEAN NOT NULL,
    quota BIGINT NOT NULL,
    exp_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
  	updated_at TIMESTAMP NOT NULL,
  	deleted_at TIMESTAMP
);

CREATE TABLE JobApplicants (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,  
    job_id BIGINT NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
  	updated_at TIMESTAMP NOT NULL,
  	deleted_at TIMESTAMP
);

ALTER TABLE UserProfiles 
ADD FOREIGN KEY (user_id) REFERENCES Users(id);

ALTER TABLE JobApplicants 
ADD FOREIGN KEY (user_id) REFERENCES Users(id);

ALTER TABLE JobApplicants 
ADD FOREIGN KEY (job_id) REFERENCES Jobs(id);

INSERT INTO Jobs (title, company, is_open, quota, exp_date, created_at, updated_at)
VALUES
('Accountant', 'Ms.Frieren', TRUE, 2, '2024-07-09 21:00:00', NOW(), NOW()),
('IT Support', 'PT. Frieren Merdeka', FALSE, 2, '2024-07-09 22:00:00', NOW(), NOW()),
('Back End', 'PT. Tata Indah Frieren', TRUE, 2, '2024-07-09 23:00:00', NOW(), NOW()),
('Satpam', 'PT. Tata Indah Frieren', TRUE, 1, '2024-07-09 23:00:00', NOW(), NOW()),
('Front End', 'PT. Frieren Rambut Putih', TRUE, 2, '2024-07-08 09:00:00', NOW(), NOW());

INSERT INTO Users (name, email, password, role, created_at, updated_at)
VALUES
('Ariana Admin', 'admin@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 'admin', NOW(), NOW()),
('Margot Robbie', 'margot@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 'applicant', NOW(), NOW()),
('Emma Watson', 'emma@gmail.com', '$2a$10$D1XNxv0r.zr83R/6b2cyc.o0SoXejOOunh6BJj7NXThVC9aS1T0zC', 'applicant', NOW(), NOW());

INSERT INTO UserProfiles (user_id, age, current_job, address, created_at, updated_at)
VALUES
(1, 22, NULL, 'Bangka Belitung', NOW(), NOW()),
(2, 34, 'Cleaning Services', 'Italy', NOW(), NOW()),
(3, 32, NULL, 'Konoha', NOW(), NOW());

INSERT INTO JobApplicants (user_id, job_id, status, created_at, updated_at)
VALUES
(2, 1, 'reject', NOW(), NOW()),
(2, 3, 'applied', NOW(), NOW()),
(3, 1, 'process', NOW(), NOW()),
(3, 3, 'reject', NOW(), NOW()),
(3, 4, 'applied', NOW(), NOW());