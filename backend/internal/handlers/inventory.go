package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weaver/api/internal/models"
	"github.com/weaver/api/internal/utils"
)

// RawSilkHandler handles raw silk purchase operations
type RawSilkHandler struct {
	db *sql.DB
}

func NewRawSilkHandler(db *sql.DB) *RawSilkHandler {
	return &RawSilkHandler{db: db}
}

// GetRawSilkPurchases retrieves all raw silk purchases
func (h *RawSilkHandler) GetRawSilkPurchases(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT id, purchase_reference, supplier_id, product_id, quantity, unit,
		       rate_per_unit, total_amount, purchase_date, expected_delivery_date,
		       actual_delivery_date, invoice_number, payment_status, amount_paid,
		       amount_pending, status, created_at
		FROM raw_silk_purchases ORDER BY created_at DESC LIMIT 100
	`)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch raw silk purchases")
		return
	}
	defer rows.Close()

	purchases := []models.RawSilkPurchase{}
	for rows.Next() {
		var p models.RawSilkPurchase
		err := rows.Scan(&p.ID, &p.PurchaseReference, &p.SupplierID, &p.ProductID,
			&p.Quantity, &p.Unit, &p.RatePerUnit, &p.TotalAmount, &p.PurchaseDate,
			&p.ExpectedDeliveryDate, &p.ActualDeliveryDate, &p.InvoiceNumber,
			&p.PaymentStatus, &p.AmountPaid, &p.AmountPending, &p.Status, &p.CreatedAt)
		if err != nil {
			continue
		}
		purchases = append(purchases, p)
	}

	utils.OKResponse(c, "Raw silk purchases retrieved", purchases)
}

// CreateRawSilkPurchase creates a new raw silk purchase
func (h *RawSilkHandler) CreateRawSilkPurchase(c *gin.Context) {
	var req models.RawSilkPurchase
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	var id int
	err := h.db.QueryRow(`
		INSERT INTO raw_silk_purchases (purchase_reference, supplier_id, product_id,
		                               quantity, unit, rate_per_unit, total_amount,
		                               purchase_date, expected_delivery_date, invoice_number,
		                               payment_status, amount_paid, amount_pending, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id
	`, req.PurchaseReference, req.SupplierID, req.ProductID, req.Quantity, req.Unit,
		req.RatePerUnit, req.TotalAmount, req.PurchaseDate, req.ExpectedDeliveryDate,
		req.InvoiceNumber, "Pending", req.AmountPaid, req.TotalAmount, "Pending").Scan(&id)

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to create raw silk purchase")
		return
	}

	req.ID = id
	utils.CreatedResponse(c, "Raw silk purchase created successfully", req)
}

// ColouringHandler handles colouring batch operations
type ColouringHandler struct {
	db *sql.DB
}

func NewColouringHandler(db *sql.DB) *ColouringHandler {
	return &ColouringHandler{db: db}
}

// GetColouringBatches retrieves all colouring batches
func (h *ColouringHandler) GetColouringBatches(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT id, colouring_reference, raw_silk_batch_id, colour_id, quantity_sent, unit,
		       colour_factory_supplier_id, date_sent, expected_return_date, date_received,
		       quantity_returned, wastage_quantity, colouring_charges, transportation_cost,
		       total_processing_cost, status, created_at
		FROM colouring_batches ORDER BY created_at DESC LIMIT 100
	`)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch colouring batches")
		return
	}
	defer rows.Close()

	batches := []models.ColouringBatch{}
	for rows.Next() {
		var b models.ColouringBatch
		err := rows.Scan(&b.ID, &b.ColouringReference, &b.RawSilkBatchID, &b.ColourID,
			&b.QuantitySent, &b.Unit, &b.ColourFactorySupplierID, &b.DateSent,
			&b.ExpectedReturnDate, &b.DateReceived, &b.QuantityReturned, &b.WastageQuantity,
			&b.ColouringCharges, &b.TransportationCost, &b.TotalProcessingCost, &b.Status, &b.CreatedAt)
		if err != nil {
			continue
		}
		batches = append(batches, b)
	}

	utils.OKResponse(c, "Colouring batches retrieved", batches)
}

// CreateColouringBatch creates a new colouring batch
func (h *ColouringHandler) CreateColouringBatch(c *gin.Context) {
	var req models.ColouringBatch
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	var id int
	err := h.db.QueryRow(`
		INSERT INTO colouring_batches (colouring_reference, raw_silk_batch_id, colour_id,
		                              quantity_sent, unit, colour_factory_supplier_id,
		                              date_sent, expected_return_date, colouring_charges,
		                              transportation_cost, total_processing_cost, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`, req.ColouringReference, req.RawSilkBatchID, req.ColourID, req.QuantitySent,
		req.Unit, req.ColourFactorySupplierID, req.DateSent, req.ExpectedReturnDate,
		req.ColouringCharges, req.TransportationCost, req.TotalProcessingCost, "Sent").Scan(&id)

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to create colouring batch")
		return
	}

	req.ID = id
	utils.CreatedResponse(c, "Colouring batch created successfully", req)
}

// InventoryHandler handles inventory operations
type InventoryHandler struct {
	db *sql.DB
}

func NewInventoryHandler(db *sql.DB) *InventoryHandler {
	return &InventoryHandler{db: db}
}

// GetInventoryStock retrieves current inventory stock
func (h *InventoryHandler) GetInventoryStock(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT DISTINCT ib.id, ib.batch_reference, ib.product_id, ib.colour_id,
		       ib.quantity_received, ib.unit, ib.current_location_id, ib.status
		FROM inventory_batches ib
		WHERE ib.status = 'In Stock'
		ORDER BY ib.created_at DESC
		LIMIT 100
	`)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch inventory stock")
		return
	}
	defer rows.Close()

	batches := []models.InventoryBatch{}
	for rows.Next() {
		var b models.InventoryBatch
		err := rows.Scan(&b.ID, &b.BatchReference, &b.ProductID, &b.ColourID,
			&b.QuantityReceived, &b.Unit, &b.CurrentLocationID, &b.Status)
		if err != nil {
			continue
		}
		batches = append(batches, b)
	}

	utils.OKResponse(c, "Inventory stock retrieved", batches)
}

// GetInventoryMovements retrieves inventory movements
func (h *InventoryHandler) GetInventoryMovements(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT id, batch_id, movement_type, quantity, unit, from_location_id,
		       to_location_id, reference_type, created_at
		FROM inventory_movements
		ORDER BY created_at DESC
		LIMIT 100
	`)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch inventory movements")
		return
	}
	defer rows.Close()

	movements := []models.InventoryMovement{}
	for rows.Next() {
		var m models.InventoryMovement
		err := rows.Scan(&m.ID, &m.BatchID, &m.MovementType, &m.Quantity, &m.Unit,
			&m.FromLocationID, &m.ToLocationID, &m.ReferenceType, &m.CreatedAt)
		if err != nil {
			continue
		}
		movements = append(movements, m)
	}

	utils.OKResponse(c, "Inventory movements retrieved", movements)
}
