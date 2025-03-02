package database

import (
	"database/sql"
	"os"
	"transfy/logs"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")

	// Chaîne de connexion corrigée
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logs.Errorf("Erreur lors de l'ouverture de la connexion : %v", err)
		return nil, err
	}

	// Vérifier la connexion
	err = db.Ping()
	if err != nil {
		logs.Errorf("Erreur lors de la connexion à la base de données : %v", err)
		return nil, err
	}

	logs.Info("Connecté à MySQL/MariaDB")
	return db, nil
}