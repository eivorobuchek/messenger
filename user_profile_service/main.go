package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type UserProfile struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}

var profiles []UserProfile

func updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	var profile UserProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, p := range profiles {
		if p.Username == profile.Username {
			profiles[i] = profile
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(profile)
			return
		}
	}
	profiles = append(profiles, profile)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

func searchProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	for _, profile := range profiles {
		if profile.Username == username {
			json.NewEncoder(w).Encode(profile)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/profile/update", updateProfileHandler)
	http.HandleFunc("/profile/search", searchProfileHandler)
	log.Println("User Profile Service is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
