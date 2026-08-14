package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weaver/api/internal/models"
	"github.com/weaver/api/internal/utils"
)

// SupplierHandler handles supplier operations
type SupplierHandler struct {
	db *sql.DB
}

func NewSupplierHandler(db *sql.DB) *SupplierHandler {
	return &SupplierHandler{db: db}
}

// GetSuppliers retrieves all suppliers
func (h *SupplierHandler) GetSuppliers(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT id, supplier_code, name, phone, email, address, city, material_type,
		       payment_terms, opening_balance, status, notes, created_at, updated_at
		FROM suppliers
		ORDER BY created_at DESC
		LIMIT 100
	`)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch suppliers")
		return
	}
	defer rows.Close()

	suppliers := []models.Supplier{}
	for rows.Next() {
		var s models.Supplier
		err := rows.Scan(&s.ID, &s.SupplierCode, &s.Name, &s.Phone, &s.Email, &s.Address,
			&s.City, &s.MaterialType, &s.PaymentTerms, &s.OpeningBalance, &s.Status,
			&s.Notes, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			continue
		}
		suppliers = append(suppliers, s)
	}

	utils.OKResponse(c, "Suppliers retrieved", suppliers)
}

// GetSupplier retrieves a specific supplier
func (h *SupplierHandler) GetSupplier(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var s models.Supplier

	err := h.db.QueryRow(`
		SELECT id, supplier_code, name, phone, email, address, city, material_type,
		       payment_terms, opening_balance, status, notes, created_at, updated_at
		FROM suppliers WHERE id = $1
	`, id).Scan(&s.ID, &s.SupplierCode, &s.Name, &s.Phone, &s.Email, &s.Address,
		&s.City, &s.MaterialType, &s.PaymentTerms, &s.OpeningBalance, &s.Status,
		&s.Notes, &s.CreatedAt, &s.UpdatedAt)

	if err == sql.ErrNoRows {
		utils.NotFoundResponse(c, "Supplier not found")
		return
	}
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch supplier")
		return
	}

	utils.OKResponse(c, "Supplier retrieved", s)
}

// CreateSupplier creates a new supplier
func (h *SupplierHandler) CreateSupplier(c *gin.Context) {
	var req models.Supplier
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	var id int
	err := h.db.QueryRow(`
		INSERT INTO suppliers (supplier_code, name, phone, email, address, city,
		                       material_type, payment_terms, opening_balance, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`, req.SupplierCode, req.Name, req.Phone, req.Email, req.Address, req.City,
		req.MaterialType, req.PaymentTerms, req.OpeningBalance, "Active").Scan(&id)

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to create supplier")
		return
	}

	req.ID = id
	utils.CreatedResponse(c, "Supplier created successfully", req)
}

// UpdateSupplier updates a supplier
func (h *SupplierHandler) UpdateSupplier(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req models.Supplier
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	_, err := h.db.Exec(`
		UPDATE suppliers
		SET name = $1, phone = $2, email = $3, address = $4, city = $5,
		    material_type = $6, payment_terms = $7, status = $8, notes = $9
		WHERE id = $10
	`, req.Name, req.Phone, req.Email, req.Address, req.City,
		req.MaterialType, req.PaymentTerms, req.Status, req.Notes, id)

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to update supplier")
		return
	}

	req.ID = id
	utils.OKResponse(c, "Supplier updated successfully", req)
}

// DeleteSupplier deletes a supplier
func (h *SupplierHandler) DeleteSupplier(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.db.Exec("DELETE FROM suppliers WHERE id = $1", id)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to delete supplier")
		return
	}

	utils.OKResponse(c, "Supplier deleted successfully", nil)
}
