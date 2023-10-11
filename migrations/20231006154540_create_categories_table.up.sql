CREATE TABLE categories (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL,
    
    PRIMARY KEY(id)
);

INSERT INTO categories (name) VALUES 
    ("Category A"), 
    ("Category B"), 
    ("Category C"),
    ("Category D");