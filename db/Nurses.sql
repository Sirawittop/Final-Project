-- Active: 1701774805198@@127.0.0.1@8889@NursePlan
CREATE TABLE Nurses (
    id INT NOT NULL AUTO_INCREMENT,
    MedicalP_ID INT,
    CitizenI_Num VARCHAR(20) NOT NULL,
    ProfessionalM_license_Num VARCHAR(20) NOT NULL,
    Firstname VARCHAR(255) NOT NULL,
    Lastname VARCHAR(255) NOT NULL,
    Birthday DATE NOT NULL,
    Position VARCHAR(50) NOT NULL,
    WardID INT,
    Employment_Year INT,
    HospitalID INT,
    Salary DECIMAL(10, 2),
    Gender VARCHAR(10) NOT NULL,
    Phone_Number VARCHAR(15),
    Number_Children INT,
    MaritalS_ID INT,
    Religion VARCHAR(50),
    Email VARCHAR(255),
    Line_ID VARCHAR(255),
    RelativeP_Num VARCHAR(15),
    Village VARCHAR(255),
    House_number VARCHAR(20),
    Zip_Code VARCHAR(10),
    Province VARCHAR(255),
    District VARCHAR(255),
    Subdistrict VARCHAR(255),
    PRIMARY KEY (id),
    FOREIGN KEY (WardID) REFERENCES Ward(id),
    FOREIGN KEY (HospitalID) REFERENCES Hospitals(id)
);
