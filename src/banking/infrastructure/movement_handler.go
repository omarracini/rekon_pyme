package infrastructure

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/omarracini/rekon_pyme/src/banking/application"
)

type MovementHandler struct {
	useCase *application.CreateMovementUseCase
}

func NewMovementHandler(uc *application.CreateMovementUseCase) *MovementHandler {
	return &MovementHandler{useCase: uc}
}

func (h *MovementHandler) CreateMovement(c *gin.Context) {
	var req application.CreateMovementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.useCase.Execute(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Movimiento registrado con éxito"})
}