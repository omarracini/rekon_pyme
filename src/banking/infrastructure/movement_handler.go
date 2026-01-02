package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omarracini/rekon_pyme/src/banking/application"
)

type MovementHandler struct {
	movementuseCase *application.CreateMovementUseCase
	invoiceUseCase  *application.CreateInvoiceUseCase
}

func NewMovementHandler(muc *application.CreateMovementUseCase, iuc *application.CreateInvoiceUseCase) *MovementHandler {
	return &MovementHandler{
		movementuseCase: muc,
		invoiceUseCase:  iuc,
	}
}

func (h *MovementHandler) CreateMovement(c *gin.Context) {
	var req application.CreateMovementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.movementuseCase.Execute(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Movimiento registrado con éxito"})
}

func (h *MovementHandler) CreateInvoice(c *gin.Context) {
	var req application.CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.invoiceUseCase.Execute(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Factura registrada con éxito"})
}

func (h *MovementHandler) ListInvoices(c *gin.Context) {
	invoices, err := h.invoiceUseCase.ExecuteList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las facturas"})
		return
	}
	c.JSON(http.StatusOK, invoices)
}
