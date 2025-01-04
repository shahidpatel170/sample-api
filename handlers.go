package main

import (
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getClientProfile(w, r)
	case http.MethodPatch:
		updateClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getClientProfile(w http.ResponseWriter, r *http.Request) {
	clientprofile := r.Context().Value("ClientProfile").(ClientProfile)
	response := ClientProfile{
		Email: clientprofile.Email,
		Name:  clientprofile.Name,
		Id:    clientprofile.Id,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func updateClientProfile(w http.ResponseWriter, r *http.Request) {
	clientprofile := r.Context().Value("clientProfile").(ClientProfile)
	var payloadData ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	clientprofile.Email = payloadData.Email
	clientprofile.Name = payloadData.Name
	database[clientprofile.Id] = clientprofile

	w.WriteHeader(http.StatusOK)
}
