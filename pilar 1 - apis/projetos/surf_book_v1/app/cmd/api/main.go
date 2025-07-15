package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/lucasbadico/surfbook-api-v1/foundation"
	"github.com/lucasbadico/surfbook-api-v1/handler"
	"github.com/lucasbadico/surfbook-api-v1/repository"
	"github.com/lucasbadico/surfbook-api-v1/service"
)

type CreateUserDTO struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
}

func createUser(db *sql.DB, id uuid.UUID, name, email string) error {
	query := `INSERT INTO users (name, email, created_at) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, name, email, "2024-07-02")
	return err
}

// Exemplo básico de consulta
func getUserByID(db *sql.DB, id uuid.UUID) (*User, error) {
	user := User{}
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	row := db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		println("err: %e", err)
		return nil, err
	}
	return &user, nil
}
//conexõ com o banco
func main() {
	log.Println("creating a webserver")
	conn := "postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable"

	db, err := foundation.NewPostgresDB(conn)
	if err != nil {
		log.Fatalf("Conexão com PostgreSQL não executada!")
	}
	fmt.Println("Conexão com PostgreSQL estabelecida com sucesso.")

	userRepo := repository.NewUserRepository(db)
	userSrv := service.NewUserService(userRepo)
	fmt.Println("User Service criado com sucesso", userSrv)

	notebookRepo := repository.NewNotebookRepository(db)
	notebookSrv := service.NewNotebookService(notebookRepo)
	fmt.Println("Notebook Service criado com sucesso", notebookSrv)

	tagRepo := repository.NewUserTagRepository(db)
	tagSrv := service.NewUserTagService(tagRepo)
	fmt.Println("Tag Service criado com sucesso", tagSrv)

	r := mux.NewRouter()
	h := handler.New(notebookSrv, tagSrv)

	h.MountHandlers(r)
	r.HandleFunc("/tags", h.CreateTag).Methods("POST")
	r.HandleFunc("/tags", h.ListTag).Methods("GET")
	r.HandleFunc("/tags/{tag_id}", h.DeleteTag).Methods("DELETE")
	
	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idString := mux.Vars(r)["id"]
		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, "falha ao de codificar id", http.StatusInternalServerError)
		}
		user, err := getUserByID(db, id)
		if err != nil {
			http.Error(w, "falha ao recuperar usuario", http.StatusInternalServerError)
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		}
	}).Methods("GET")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var input CreateUserDTO
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
			return
		}
		id := uuid.New()
		err := createUser(db, id, input.Name, input.Email)
		if err != nil {
			http.Error(w, "falha ao recuperar usuario", http.StatusInternalServerError)
		}

		user := User{
			Name:  input.Name,
			Email: input.Email,
			ID:    id,
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		}
	}).Methods("POST")

	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", r)
}
