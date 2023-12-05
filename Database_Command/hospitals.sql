-- Add a new table for hospitals
CREATE TABLE IF NOT EXISTS hospitals (
    hospital_code INT PRIMARY KEY AUTO_INCREMENT,
    hospital_name VARCHAR(255) NOT NULL
);

-- Sample data for initial testing
INSERT INTO hospitals (hospital_name) VALUES
    ('โรงพยาบาลมหาวิทยาลัยพะเยา');
