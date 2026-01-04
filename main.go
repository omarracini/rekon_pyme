package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/omarracini/rekon_pyme/docs"
	"github.com/omarracini/rekon_pyme/src/banking/application"
	"github.com/omarracini/rekon_pyme/src/banking/infrastructure"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// Metadatos para Swagger

// @title           Rekon Pyme API
// @version         1.0
// @description     MVP de conciliación bancaria para prueba técnica.
// @host            localhost:8080
// @BasePath        /

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=user_pyme password=password_pyme dbname=conciliacion_db sslmode=disable")
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	// Inyección de dependencias
	repo := infrastructure.NewPostgresBankRepository(db)
	aiClient := infrastructure.NewAIClient()

	// Casos de uso
	movUseCase := application.NewCreateMovementUseCase(repo)
	invUseCase := application.NewCreateInvoiceUseCase(repo)
	concUseCase := application.NewConciliateUseCase(repo)
	pendUseCase := application.NewGetPendingItemsUseCase(repo)
	getPendingUC := application.NewGetPendingMovementsUseCase(repo)
	getPendingInvUC := application.NewGetPendingInvoicesUseCase(repo)
	getDashboardUC := application.NewGetDashboardUseCase(repo)
	suggestCategoryUC := application.NewSuggestCategoryUseCase(aiClient)

	// Handlers
	handler := infrastructure.NewMovementHandler(
		movUseCase,
		invUseCase,
		concUseCase,
		pendUseCase,
		getPendingUC,
		getPendingInvUC,
		getDashboardUC,
		suggestCategoryUC)

	// Configuración de Gin
	r := gin.Default()

	// Rutas
	r.POST("/movements", handler.CreateMovement)
	r.POST("/invoices", handler.CreateInvoice)
	r.GET("/invoices", handler.ListInvoices)
	r.POST("/conciliations", handler.Conciliate)
	r.GET("/pending", handler.GetPending)
	r.GET("/movements/pending", handler.GetPending)
	r.GET("/invoices/pending", handler.GetPendingInvoices)
	r.GET("/dashboard", handler.GetDashboard)
	r.GET("/health", handler.HealthCheck)
	r.GET("/ai/suggest-category", handler.SuggestCategory)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Servidor iniciado en http://localhost:8080")

	r.Run(":8080")
}
