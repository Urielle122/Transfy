CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY AUTO INCREMENT,
    nom VARCHAR(255),
    prenom VARCHAR(50),
    age VARCHAR(10),
    contact VARCHAR(10), 
    password VARCHAR(100)
);

admin

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '';
FLUSH PRIVILEGES;

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'MySQL_P@ssw0rd';
FLUSH PRIVILEGES;

ALTER USER 'admin'@'localhost' IDENTIFIED WITH mysql_native_password BY 'MySQL_P@ssw0rdss';