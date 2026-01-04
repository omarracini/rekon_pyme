package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omarracini/rekon_pyme/src/banking/application"
)

// Argumentos
type MovementHandler struct {
	movementuseCase   *application.CreateMovementUseCase
	invoiceUseCase    *application.CreateInvoiceUseCase
	conciliateUseCase *application.ConciliateUseCase
	pendingUseCase    *application.GetPendingItemsUseCase
	getPendingUC      *application.GetPendingMovementsUseCase
	getPendingInvUC   *application.GetPendingInvoicesUseCase
	getDashboardUC    *application.GetDashboardUseCase
}

// Constructor
func NewMovementHandler(
	muc *application.CreateMovementUseCase,
	iuc *application.CreateInvoiceUseCase,
	cuc *application.ConciliateUseCase,
	puc *application.GetPendingItemsUseCase,
	gpuc *application.GetPendingMovementsUseCase,
	gpiuc *application.GetPendingInvoicesUseCase,
	gduc *application.GetDashboardUseCase) *MovementHandler {
	return &MovementHandler{
		movementuseCase:   muc,
		invoiceUseCase:    iuc,
		conciliateUseCase: cuc,
		pendingUseCase:    puc,
		getPendingUC:      gpuc,
		getPendingInvUC:   gpiuc,
		getDashboardUC:    gduc,
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

func (h *MovementHandler) Conciliate(c *gin.Context) {
	fmt.Println(">>> PETICION RECIBIDA EN HANDLER CONCILIATE") // <--- LOG 1

	var req application.ConciliateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("ERROR JSON: %v\n", err) // <--- LOG 2
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de conciliación inválidos"})
		return
	}

	fmt.Printf("INTENTANDO CONCILIAR: Mov=%s, Inv=%s\n", req.MovementID, req.InvoiceID) // <--- LOG 3

	if err := h.conciliateUseCase.Execute(req); err != nil {
		fmt.Printf("ERROR EN USE CASE: %v\n", err)                          // <--- ESTO ES LO QUE NECESITAMOS SABER
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Devolvemos el error real
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movimiento y factura conciliados con éxito"})
}

func (h *MovementHandler) GetPending(c *gin.Context) {
	movements, err := h.getPendingUC.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se logró obtener los movimientos"})
		return
	}
	c.JSON(http.StatusOK, movements)
}

func (h *MovementHandler) GetPendingInvoices(c *gin.Context) {
	invoices, err := h.getPendingInvUC.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se lograron obtener las facturas pendientes"})
		return
	}
	c.JSON(http.StatusOK, invoices)
}

func (h *MovementHandler) GetDashboard(c *gin.Context) {
	summary, err := h.getDashboardUC.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
