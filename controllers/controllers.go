package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"teste/database"
	"teste/models"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		getAllUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No response")
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, nome FROM usuarios")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error querying database: %v", err)
		return
	}
	defer rows.Close()

	var users []models.Usuario
	for rows.Next() {
		var usuario models.Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nome); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error scanning rows: %v", err)
			return
		}
		users = append(users, usuario)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error encoding JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
