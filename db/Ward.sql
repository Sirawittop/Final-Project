-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
CREATE Table Ward (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    hospital_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (hospital_id) REFERENCES Hospitals(id)
);