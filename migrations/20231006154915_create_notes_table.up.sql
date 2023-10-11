CREATE TABLE notes (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    body VARCHAR(255) NOT NULL,
    id_category INT NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (id_category) REFERENCES categories (id) ON DELETE CASCADE
);

INSERT INTO notes (title, body, id_category) VALUES
    ("Title A", "Body A", 1),
    ("Title B", "Body B", 2),
    ("Title C", "Body C", 3),
    ("Title D", "Body D", 4);