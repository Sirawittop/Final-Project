-- Add a new table for ward data
CREATE TABLE IF NOT EXISTS ward_data (
    Ward_ID INT PRIMARY KEY AUTO_INCREMENT,
    hospital_code INT,
    ward_name VARCHAR(255) NOT NULL,
    FOREIGN KEY (hospital_code) REFERENCES hospitals(hospital_code)
);

-- Sample data for initial testing
INSERT INTO ward_data (ward_name) VALUES
    ('งานการพยาบาลผู้ป่วยใน');
