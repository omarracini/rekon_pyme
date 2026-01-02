package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/omarracini/rekon_pyme/src/banking/application"
	"github.com/omarracini/rekon_pyme/src/banking/infrastructure"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=user_pyme password=password_pyme dbname=conciliacion_db sslmode=disable")
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	// Inyección de dependencias
	repo := infrastructure.NewPostgresBankRepository(db)
	useCase := application.NewCreateMovementUseCase(repo)
	handler := infrastructure.NewMovementHandler(useCase)

	r := gin.Default()
	r.POST("/movements", handler.CreateMovement)

	log.Println("Servidor iniciado en http://localhost:8080")
	r.Run(":8080")
}
