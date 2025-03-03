package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v5"

	//"transfy/logs"
	"transfy/core"
	"transfy/models"
)


var jwtKey = []byte()

// Claims représente les données à inclure dans le JWT
type Claims struct {
	Contact string `json:"contact"`
	jwt.RegisteredClaims
}

func UserSession(w http.ResponseWriter, r *http.Request) {
	var session models.Session

	// Décoder le corps de la requête
	bodyReader, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(bodyReader, &session)
	if err != nil {
		http.Error(w, "Erreur lors du décodage JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Vérifier les identifiants de l'utilisateur
	var dbContact, dbPassword string
	row := core.MysqlDb.QueryRow("SELECT contact, password FROM users WHERE contact = ?", session.Contact)
	err = row.Scan(&dbContact, &dbPassword)
	
	
	if err != nil {
		// Utilisateur non trouvé
		http.Error(w, "Identifiants invalides", http.StatusUnauthorized)
		return
	}

	// Vérifier le mot de passe (en production, utilisez bcrypt pour comparer)
	if dbPassword != session.Password {
		http.Error(w, "Identifiants invalides", http.StatusUnauthorized)
		return
	}

	// Générer le JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Contact: session.Contact,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "transfy",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	
	if err != nil {
		http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}

	// Créer la réponse
	response := struct {
		Token   string `json:"token"`
		Message string `json:"message"`
		Statut 	string `json: "statut"`
	}{
		Token:   tokenString,
		Message: "Connexion réussie",
		Statut:  http.StatusOK 

	}

	// Envoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}