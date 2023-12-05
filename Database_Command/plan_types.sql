-- Create a new database if it doesn't exist
CREATE DATABASE IF NOT EXISTS nurse_scheduler;

-- Switch to the database
USE nurse_scheduler;

-- Create a table for the type of plan
CREATE TABLE IF NOT EXISTS plan_types (
    TypeID INT AUTO_INCREMENT PRIMARY KEY,
    TypeName VARCHAR(10) NOT NULL,
    Morning BOOLEAN,
    Afternoon BOOLEAN,
    Night BOOLEAN,
    X BOOLEAN,
    V BOOLEAN,
    Nleave BOOLEAN,
    C BOOLEAN,
    Part BOOLEAN,
    OtM BOOLEAN,
    OtA BOOLEAN,
    OtN BOOLEAN,
    StatusS BOOLEAN,
    StatusOt12 BOOLEAN,
    StatusOt8 BOOLEAN,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Sample data for initial testing
INSERT INTO plan_types (TypeName, Morning, Afternoon, Night, X, V, NLeave, C, Part, OtM, OtA, OtN, StatusS, StatusOt12, StatusOt8)
VALUES
    ('ช', true, false, false, false, false, false, false, false, false, false, false, false, false, false),
    ('บ', false, true, false, false, false, false, false, false, false, false, false, false, false, false),
    ('ด', false, false, true, false, false, false, false, false, false, false, false, false, false, false),
    ('t', false, false, false, true, false, false, false, false, false, false, false, false, false, false),
    ('v', false, false, false, false, true, false, false, false, false, false, false, false, false, false),
    ('T', true, false, false, false, false, false, false, false, false, false, false, false, false, false),
    ('ล', false, false, false, false, false, true, false, false, false, false, false, false, false, false),
    ('C', false, false, false, false, false, false, true, false, false, false, false, false, false, false),
    ('น', false, false, false, false, false, false, false, true, false, false, false, false, false, false),
    ('pช/บ', false, true, false, false, false, false, false, false, true, false, false, false, false, false),
    ('ช/pบ', true, false, false, false, false, false, false, false, false, true, false, false, false, false),
    ('pบ/ด', false, false, true, false, false, false, false, false, false, true, false, false, false, false),
    ('บ/pด', false, true, false, false, false, false, false, false, false, false, true, false, false, false),
    ('pช/ด', false, false, true, false, false, false, false, false, true, false, false, false, false, false),
    ('ช/pด', true, false, false, false, false, false, false, false, false, true, false, false, false, false),
    ('x/ช', false, false, false, false, false, false, false, false, true, false, false, false, false, false),
    ('x/บ', false, false, false, false, false, false, false, false, false, true, false, false, false, false),
    ('x/ด', false, false, false, false, false, false, false, false, false, false, true, false, false, false),
    ('ช2', true, false, false, false, false, false, false, false, false, false, false, false, false, false);