CREATE TABLE IF NOT EXISTS users (
    Id BINARY(16) PRIMARY KEY,
    FirstName VARCHAR(255),
    LastName VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS
  addresses (
    Id binary(16) NOT NULL,
    UserId binary(16) NOT NULL,
    Street varchar(255) DEFAULT NULL,
    City varchar(255) DEFAULT NULL,
    State varchar(255) DEFAULT NULL,
    Zip varchar(255) DEFAULT NULL,
    `Type` varchar(255) DEFAULT NULL,
    PRIMARY KEY (Id),
    KEY addresses_users (UserId),
    CONSTRAINT addresses_users FOREIGN KEY (UserId) REFERENCES users (Id)
  )