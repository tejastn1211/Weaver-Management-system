package models

import "time"

// User represents a system user
type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
	Role     string    `json:"role"`
	Status   string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Supplier represents a supplier/vendor
type Supplier struct {
	ID              int       `json:"id"`
	SupplierCode    string    `json:"supplier_code"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Address         string    `json:"address"`
	City            string    `json:"city"`
	MaterialType    string    `json:"material_type"`
	PaymentTerms    string    `json:"payment_terms"`
	OpeningBalance  float64   `json:"opening_balance"`
	Status          string    `json:"status"`
	Notes           string    `json:"notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Weaver represents a weaver
type Weaver struct {
	ID              int       `json:"id"`
	WeaverCode      string    `json:"weaver_code"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Address         string    `json:"address"`
	Village         string    `json:"village"`
	JoiningDate     string    `json:"joining_date"`
	BankAccount     string    `json:"bank_account_number"`
	BankName        string    `json:"bank_name"`
	OpeningBalance  float64   `json:"opening_balance"`
	Status          string    `json:"status"`
	Notes           string    `json:"notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Buyer represents a buyer/customer
type Buyer struct {
	ID              int       `json:"id"`
	BuyerCode       string    `json:"buyer_code"`
	Name            string    `json:"name"`
	BusinessName    string    `json:"business_name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Address         string    `json:"address"`
	City            string    `json:"city"`
	GSTNumber       string    `json:"gst_number"`
	CreditLimit     float64   `json:"credit_limit"`
	OpeningBalance  float64   `json:"opening_balance"`
	Status          string    `json:"status"`
	Notes           string    `json:"notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Product represents a product
type Product struct {
	ID          int       `json:"id"`
	ProductCode string    `json:"product_code"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Unit        string    `json:"unit"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// Colour represents a colour
type Colour struct {
	ID       int       `json:"id"`
	Code     string    `json:"colour_code"`
	Name     string    `json:"name"`
	HexCode  string    `json:"hex_code"`
	Status   string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// Location represents a location/warehouse
type Location struct {
	ID          int       `json:"id"`
	Code        string    `json:"location_code"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// Design represents a saree design
type Design struct {
	ID        int       `json:"id"`
	Code      string    `json:"design_code"`
	Name      string    `json:"name"`
	Description string  `json:"description"`
	SareeType string    `json:"saree_type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// InventoryBatch represents an inventory batch
type InventoryBatch struct {
	ID              int       `json:"id"`
	BatchReference  string    `json:"batch_reference"`
	ProductID       int       `json:"product_id"`
	ColourID        *int      `json:"colour_id"`
	SupplierID      *int      `json:"supplier_id"`
	QuantityReceived float64  `json:"quantity_received"`
	Unit            string    `json:"unit"`
	CostPerUnit     float64   `json:"cost_per_unit"`
	TotalCost       float64   `json:"total_cost"`
	CurrentLocationID *int    `json:"current_location_id"`
	ReceivedDate    string    `json:"received_date"`
	Status          string    `json:"status"`
	Notes           string    `json:"notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// InventoryMovement represents a movement in inventory
type InventoryMovement struct {
	ID              int       `json:"id"`
	BatchID         int       `json:"batch_id"`
	MovementType    string    `json:"movement_type"`
	Quantity        float64   `json:"quantity"`
	Unit            string    `json:"unit"`
	FromLocationID  *int      `json:"from_location_id"`
	ToLocationID    *int      `json:"to_location_id"`
	ReferenceType   string    `json:"reference_type"`
	ReferenceID     *int      `json:"reference_id"`
	CreatedBy       *int      `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	Notes           string    `json:"notes"`
}

// RawSilkPurchase represents a raw silk purchase
type RawSilkPurchase struct {
	ID                    int       `json:"id"`
	PurchaseReference     string    `json:"purchase_reference"`
	SupplierID            int       `json:"supplier_id"`
	ProductID             int       `json:"product_id"`
	Quantity              float64   `json:"quantity"`
	Unit                  string    `json:"unit"`
	RatePerUnit           float64   `json:"rate_per_unit"`
	TotalAmount           float64   `json:"total_amount"`
	PurchaseDate          string    `json:"purchase_date"`
	ExpectedDeliveryDate  string    `json:"expected_delivery_date"`
	ActualDeliveryDate    string    `json:"actual_delivery_date"`
	InvoiceNumber         string    `json:"invoice_number"`
	PaymentStatus         string    `json:"payment_status"`
	AmountPaid            float64   `json:"amount_paid"`
	AmountPending         float64   `json:"amount_pending"`
	Status                string    `json:"status"`
	Notes                 string    `json:"notes"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

// ColouringBatch represents a colouring batch
type ColouringBatch struct {
	ID                        int       `json:"id"`
	ColouringReference        string    `json:"colouring_reference"`
	RawSilkBatchID            int       `json:"raw_silk_batch_id"`
	ColourID                  int       `json:"colour_id"`
	QuantitySent              float64   `json:"quantity_sent"`
	Unit                      string    `json:"unit"`
	ColourFactorySupplierID   int       `json:"colour_factory_supplier_id"`
	DateSent                  string    `json:"date_sent"`
	ExpectedReturnDate        string    `json:"expected_return_date"`
	DateReceived              string    `json:"date_received"`
	QuantityReturned          float64   `json:"quantity_returned"`
	WastageQuantity           float64   `json:"wastage_quantity"`
	ColouringCharges          float64   `json:"colouring_charges"`
	TransportationCost        float64   `json:"transportation_cost"`
	TotalProcessingCost       float64   `json:"total_processing_cost"`
	Status                    string    `json:"status"`
	Notes                     string    `json:"notes"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}

// DoliBatch represents a doli (warping) batch
type DoliBatch struct {
	ID                          int       `json:"id"`
	DoliReference               string    `json:"doli_reference"`
	SourceColouredSilkBatchID   int       `json:"source_coloured_silk_batch_id"`
	ColourID                    int       `json:"colour_id"`
	QuantityCreated             float64   `json:"quantity_created"`
	SareeCapacity               int       `json:"saree_capacity"`
	Weight                      float64   `json:"weight"`
	WarpingCharges              float64   `json:"warping_charges"`
	WarpingDate                 string    `json:"warping_date"`
	CurrentLocationID           *int      `json:"current_location_id"`
	Status                      string    `json:"status"`
	Notes                       string    `json:"notes"`
	CreatedAt                   time.Time `json:"created_at"`
}

// WindingBatch represents a winding batch
type WindingBatch struct {
	ID                          int       `json:"id"`
	WindingReference            string    `json:"winding_reference"`
	SourceColouredSilkBatchID   int       `json:"source_coloured_silk_batch_id"`
	ColourID                    int       `json:"colour_id"`
	QuantityWound               float64   `json:"quantity_wound"`
	BobbinsProduced             float64   `json:"bobbins_produced"`
	PirnsProduced               float64   `json:"pirns_produced"`
	WindingCharges              float64   `json:"winding_charges"`
	MachineOperator             string    `json:"machine_operator"`
	WindingDate                 string    `json:"winding_date"`
	Status                      string    `json:"status"`
	CreatedAt                   time.Time `json:"created_at"`
}

// BobbiInventory represents bobbin inventory
type BobbiInventory struct {
	ID                   int       `json:"id"`
	BobbiiBatchReference string    `json:"bobbin_batch_reference"`
	ColourID             int       `json:"colour_id"`
	SourceWindingBatchID *int      `json:"source_winding_batch_id"`
	Quantity             float64   `json:"quantity"`
	Unit                 string    `json:"unit"`
	CurrentLocationID    *int      `json:"current_location_id"`
	Status               string    `json:"status"`
	CreatedAt            time.Time `json:"created_at"`
}

// Response represents a standard API response
type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Timestamp  time.Time   `json:"timestamp"`
}

// LoginRequest represents login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents login response
type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
