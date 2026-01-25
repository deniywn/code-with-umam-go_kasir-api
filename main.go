package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Krayon", Description: "Berbagai jenis produk kertas Krayon"},
	{ID: 2, Name: "Kertas A4", Description: "Berbagai jenis produk kertas A4."},
}

func categoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID request", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		for _, c := range categories {
			if c.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(c)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)

	case http.MethodPut:
		var updatedCategory Category
		if err := json.NewDecoder(r.Body).Decode(&updatedCategory); err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		for i, c := range categories {
			if c.ID == id {
				updatedCategory.ID = id
				categories[i] = updatedCategory
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(updatedCategory)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)
	case http.MethodDelete:
		for i, c := range categories {
			if c.ID == id {
				categories = append(categories[:i], categories[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		http.Error(w, "Category not found", http.StatusNotFound)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/category/", categoryByIDHandler)
	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
		case "POST":
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			newCategory.ID = len(categories) + 1
			categories = append(categories, newCategory)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCategory)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Api running",
			"status":  "ok",
		})
	})
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
