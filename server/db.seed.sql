-- SQL Seeding file for generating 50 entries for the 'users' table

-- Truncate the table to start fresh
TRUNCATE TABLE users;

-- Insert 50 entries with real email providers and current timestamp
INSERT INTO users (name, email, created_at) VALUES
('John Doe', 'john.doe@gmail.com', NOW()),
('Jane Smith', 'jane.smith@yahoo.com', NOW()),
('Michael Johnson', 'michael.johnson@hotmail.com', NOW()),
('Emily Davis', 'emily.davis@aol.com', NOW()),
('Chris Wilson', 'chris.wilson@outlook.com', NOW()),
('Sarah Brown', 'sarah.brown@icloud.com', NOW()),
('David Lee', 'david.lee@protonmail.com', NOW()),
('Amanda Garcia', 'amanda.garcia@mail.com', NOW()),
('Kevin Martinez', 'kevin.martinez@yandex.com', NOW()),
('Jennifer Rodriguez', 'jennifer.rodriguez@zoho.com', NOW()),
('James Taylor', 'james.taylor@inbox.com', NOW()),
('Jessica Hernandez', 'jessica.hernandez@live.com', NOW()),
('Daniel Gonzalez', 'daniel.gonzalez@fastmail.com', NOW()),
('Lisa Martinez', 'lisa.martinez@rocketmail.com', NOW()),
('Brian Walker', 'brian.walker@iinet.net.au', NOW()),
('Nicole Perez', 'nicole.perez@att.net', NOW()),
('Matthew Turner', 'matthew.turner@shaw.ca', NOW()),
('Ashley King', 'ashley.king@sbcglobal.net', NOW()),
('Joshua Young', 'joshua.young@verizon.net', NOW()),
('Kimberly Mitchell', 'kimberly.mitchell@cox.net', NOW()),
('Andrew Adams', 'andrew.adams@ntlworld.com', NOW()),
('Alyssa Evans', 'alyssa.evans@talktalk.net', NOW()),
('Ryan Campbell', 'ryan.campbell@btinternet.com', NOW()),
('Lauren Hill', 'lauren.hill@comcast.net', NOW()),
('Justin Green', 'justin.green@virginmedia.com', NOW()),
('Megan Baker', 'megan.baker@earthlink.net', NOW()),
('Brandon Wright', 'brandon.wright@orange.fr', NOW()),
('Stephanie Carter', 'stephanie.carter@sky.com', NOW()),
('Timothy Hall', 'timothy.hall@ntlworld.com', NOW()),
('Rebecca Torres', 'rebecca.torres@cox.net', NOW()),
('Jordan Nelson', 'jordan.nelson@comcast.net', NOW()),
('Mary White', 'mary.white@verizon.net', NOW()),
('Paul Murphy', 'paul.murphy@earthlink.net', NOW()),
('Samantha Rivera', 'samantha.rivera@ntlworld.com', NOW()),
('Nathan Thomas', 'nathan.thomas@ntlworld.com', NOW()),
('Victoria Collins', 'victoria.collins@ntlworld.com', NOW()),
('Eric Brooks', 'eric.brooks@ntlworld.com', NOW()),
('Rachel Sanchez', 'rachel.sanchez@ntlworld.com', NOW()),
('Tyler Ross', 'tyler.ross@ntlworld.com', NOW()),
('Hannah Ramirez', 'hannah.ramirez@ntlworld.com', NOW()),
('Adam Price', 'adam.price@ntlworld.com', NOW()),
('Olivia Scott', 'olivia.scott@ntlworld.com', NOW()),
('Peter Morris', 'peter.morris@ntlworld.com', NOW()),
('Alexa Stewart', 'alexa.stewart@ntlworld.com', NOW()),
('Benjamin Ward', 'benjamin.ward@ntlworld.com', NOW()),
('Diana Powell', 'diana.powell@ntlworld.com', NOW()),
('Patrick Foster', 'patrick.foster@ntlworld.com', NOW()),
('Catherine Simmons', 'catherine.simmons@ntlworld.com', NOW());
