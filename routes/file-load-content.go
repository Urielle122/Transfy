package routes

import (
	//"context"
	//"context"
	"encoding/json"
	"net/http"
	"transfy/core"
	"transfy/logs"
	"transfy/models"

	"github.com/google/uuid"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
  var body models.User

  // Décoder le corps de la requête
  if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
      logs.Errorf("Erreur lors du décodage du JSON: %v", err)
      w.WriteHeader(http.StatusBadRequest)
      json.NewEncoder(w).Encode(models.Response{
          Success: false,
          Message: "Format de données invalide",
      })
      return
  }

  // Générer un UUID pour l'ID
  body.ID = uuid.New().String()

  // Exécuter la requête d'insertion
  _, err := core.MysqlDb.Exec(
      "INSERT INTO users (id, nom, prenom, age, contact, password) VALUES (?, ?, ?, ?, ?, ?)",
      body.ID, body.Nom, body.Prenom, body.Age, body.Contact, body.Password,
  )
  if err != nil {
      logs.Errorf("Erreur lors de l'insertion: %v", err)
      w.WriteHeader(http.StatusInternalServerError)
      json.NewEncoder(w).Encode(models.Response{
          Success: false,
          Message: "Erreur lors de l'ajout de l'utilisateur",
      })
      return
  }

  // Répondre avec succès
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(models.Response{
      Success: true,
      Message: "Utilisateur ajouté avec succès",
      Data:    &body,
  })
}