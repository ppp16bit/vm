package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := LoadConf()

	db, err := NewDB(cfg)
	if err != nil {
		log.Fatal("failed to initialize db")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("db unreachable")
	}

	repo := NewUserRepository(db)
	handler := NewUserHandler(repo)

	http.HandleFunc("/users", handler.CreateUser)
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetUser(w, r)
		case http.MethodDelete:
			handler.DeleteUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Printf("running on port: %s :>\n", cfg.APIPort)
	log.Fatal(http.ListenAndServe(":"+cfg.APIPort, nil))
}
