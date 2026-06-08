-- Create tables for users, authors, genres, books, and loans
CREATE DATABASE users (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE DATABASE authors (
    ID SERIAL PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL
);

CREATE DATABASE genres (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(50) NOT NULL
);

CREATE DATABASE books (
    ID SERIAL PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    AuthorID INT REFERENCES AUTHORS(ID),
    GenreID INT REFERENCES GENRES(ID),
    PublishedDate DATE,
    AvailableCopies INT DEFAULT 0
);

CREATE DATABASE loans (
    ID SERIAL PRIMARY KEY,
    UserID INT REFERENCES USERS(ID),
    BookID INT REFERENCES BOOKS(ID),
    LoanDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ReturnDate TIMESTAMP
);

-- Insert sample data into authors, genres, and books tables
INSERT INTO authors (FirstName, LastName) VALUES
('J.K.', 'Rowling'),
('George', 'Orwell'),
('Harper', 'Lee');

INSERT INTO genres (Name) VALUES
('Fantasy'),
('Dystopian'),
('Classic');

INSERT INTO books (Title, AuthorID, GenreID, PublishedDate, AvailableCopies) VALUES
('Harry Potter and the Sorcerer''s Stone', 1, 1, '1997-06-26', 5),
('1984', 2, 2, '1949-06-08', 3),
('To Kill a Mockingbird', 3, 3, '1960-07-11', 4);
