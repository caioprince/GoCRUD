package main

import (
	"GoCRUD/store"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/google/uuid"
)

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func main() {
	userStore := store.New()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var u store.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			writeJSONError(w, http.StatusBadRequest, "Invalid JSON payload")
			return
		}

		if err := validateUser(u); err != nil {
			writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		createdUser := userStore.Insert(u)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdUser)
	})

	mux.HandleFunc("GET /api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		users := userStore.FindAll()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	})

	mux.HandleFunc("GET /api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			writeJSONError(w, http.StatusNotFound, "User not found")
			return
		}

		user, err := userStore.FindById(id)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				writeJSONError(w, http.StatusNotFound, "User not found")
				return
			}
			writeJSONError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	})

	mux.HandleFunc("PUT /api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			writeJSONError(w, http.StatusNotFound, "User not found")
			return
		}

		var u store.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			writeJSONError(w, http.StatusBadRequest, "Invalid JSON payload")
			return
		}

		if err := validateUser(u); err != nil {
			writeJSONError(w, http.StatusBadRequest, err.Error())
			return
		}

		updatedUser, err := userStore.Update(id, u)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				writeJSONError(w, http.StatusNotFound, "User not found")
				return
			}
			writeJSONError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedUser)
	})

	mux.HandleFunc("DELETE /api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := r.PathValue("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			writeJSONError(w, http.StatusNotFound, "User not found")
			return
		}

		deletedUser, err := userStore.Delete(id)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				writeJSONError(w, http.StatusNotFound, "User not found")
				return
			}
			writeJSONError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(deletedUser)
	})

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}

func validateUser(u store.User) error {
	fnLen := utf8.RuneCountInString(u.FirstName)
	if fnLen < 2 || fnLen > 20 {
		return errors.New("first_name must be between 2 and 20 characters")
	}

	lnLen := utf8.RuneCountInString(u.LastName)
	if lnLen < 2 || lnLen > 20 {
		return errors.New("last_name must be between 2 and 20 characters")
	}

	bioLen := utf8.RuneCountInString(u.Biography)
	if bioLen < 20 || bioLen > 450 {
		return errors.New("biography must be between 20 and 450 characters")
	}

	return nil
}
