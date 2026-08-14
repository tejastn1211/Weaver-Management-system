package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/weaver/api/internal/models"
	"github.com/weaver/api/internal/utils"
)

// WeaverHandler handles weaver operations
type WeaverHandler struct {
	db *sql.DB
}

func NewWeaverHandler(db *sql.DB) *WeaverHandler {
	return &WeaverHandler{db: db}
}

// GetWeavers retrieves all weavers
func (h *WeaverHandler) GetWeavers(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT id, weaver_code, name, phone, email, address, village,
		       joining_date, bank_account_number, bank_name, opening_balance, status, notes, created_at
		FROM weavers ORDER BY created_at DESC LIMIT 100
	`)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch weavers")
		return
	}
	defer rows.Close()

	weavers := []models.Weaver{}
	for rows.Next() {
		var w models.Weaver
		err := rows.Scan(&w.ID, &w.WeaverCode, &w.Name, &w.Phone, &w.Email, &w.Address,
			&w.Village, &w.JoiningDate, &w.BankAccount, &w.BankName, &w.OpeningBalance,
			&w.Status, &w.Notes, &w.CreatedAt)
		if err != nil {
			continue
		}
		weavers = append(weavers, w)
	}

	utils.OKResponse(c, "Weavers retrieved", weavers)
}

// CreateWeaver creates a new weaver
func (h *WeaverHandler) CreateWeaver(c *gin.Context) {
	var req models.Weaver
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	var id int
	err := h.db.QueryRow(`
		INSERT INTO weavers (weaver_code, name, phone, email, address, village, joining_date,
		                    bank_account_number, bank_name, opening_balance, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`, req.WeaverCode, req.Name, req.Phone, req.Email, req.Address, req.Village,
		req.JoiningDate, req.BankAccount, req.BankName, req.OpeningBalance, "Active").Scan(&id)

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to create weaver")
		return
	}

	req.ID = id
	utils.CreatedResponse(c, "Weaver created successfully", req)
}

// BuyerHandler handles buyer operations
type BuyerHandler struct {
	db *sql.DB
}

func NewBuyerHandler(db *sql.DB) *BuyerHandler {
	return &BuyerHandler{db: db}
}

// GetBuyers retrieves all buyers
func (h *BuyerHandler) GetBuyers(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT id, buyer_code, name, business_name, phone, email, address, city,
		       gst_number, credit_limit, opening_balance, status, created_at
		FROM buyers ORDER BY created_at DESC LIMIT 100
	`)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch buyers")
		return
	}
	defer rows.Close()

	buyers := []models.Buyer{}
	for rows.Next() {
		var b models.Buyer
		err := rows.Scan(&b.ID, &b.BuyerCode, &b.Name, &b.BusinessName, &b.Phone, &b.Email,
			&b.Address, &b.City, &b.GSTNumber, &b.CreditLimit, &b.OpeningBalance,
			&b.Status, &b.CreatedAt)
		if err != nil {
			continue
		}
		buyers = append(buyers, b)
	}

	utils.OKResponse(c, "Buyers retrieved", buyers)
}

// CreateBuyer creates a new buyer
func (h *BuyerHandler) CreateBuyer(c *gin.Context) {
	var req models.Buyer
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	var id int
	err := h.db.QueryRow(`
		INSERT INTO buyers (buyer_code, name, business_name, phone, email, address, city,
		                   gst_number, credit_limit, opening_balance, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`, req.BuyerCode, req.Name, req.BusinessName, req.Phone, req.Email, req.Address,
		req.City, req.GSTNumber, req.CreditLimit, req.OpeningBalance, "Active").Scan(&id)

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to create buyer")
		return
	}

	req.ID = id
	utils.CreatedResponse(c, "Buyer created successfully", req)
}
