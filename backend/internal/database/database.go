package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	fmt.Println("✓ Database connection successful")
	return db, nil
}

func RunMigrations(db *sql.DB) error {
	migrations := []string{
		createUsersTable,
		createSuppliersTable,
		createWeaversTable,
		createBuyersTable,
		createProductsTable,
		createColoursTable,
		createLocationsTable,
		createDesignsTable,
		createInventoryBatchesTable,
		createInventoryMovementsTable,
		createRawSilkPurchasesTable,
		createColouringBatchesTable,
		createDoliBatchesTable,
		createWindingBatchesTable,
		createBobbiInventoryTable,
		createSupplierTransactionsTable,
		createAuditLogsTable,
	}

	for i, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			// Check if table already exists
			if isTableExistsError(err) {
				continue
			}
			return fmt.Errorf("migration %d failed: %v", i+1, err)
		}
	}

	fmt.Println("✓ Migrations completed")
	return nil
}

func isTableExistsError(err error) bool {
	// Simple check - in production, handle this more carefully
	return err != nil && (err.Error() == "pq: relation already exists" || 
		err.Error() == "pq: duplicate key value violates unique constraint \"pg_type_typname_nsp_index\"")
}

// ... Migration SQL strings follow ...

const createUsersTable = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(200),
    role VARCHAR(50),
    status VARCHAR(20) DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createSuppliersTable = `
CREATE TABLE IF NOT EXISTS suppliers (
    id SERIAL PRIMARY KEY,
    supplier_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(150),
    address TEXT,
    city VARCHAR(50),
    material_type VARCHAR(100),
    payment_terms VARCHAR(100),
    bank_account_number VARCHAR(50),
    bank_name VARCHAR(100),
    opening_balance DECIMAL(12, 2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'Active',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createWeaversTable = `
CREATE TABLE IF NOT EXISTS weavers (
    id SERIAL PRIMARY KEY,
    weaver_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(150),
    address TEXT,
    village VARCHAR(100),
    joining_date DATE NOT NULL,
    bank_account_number VARCHAR(50),
    bank_name VARCHAR(100),
    opening_balance DECIMAL(12, 2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'Active',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createBuyersTable = `
CREATE TABLE IF NOT EXISTS buyers (
    id SERIAL PRIMARY KEY,
    buyer_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    business_name VARCHAR(200),
    phone VARCHAR(20),
    email VARCHAR(150),
    address TEXT,
    city VARCHAR(50),
    gst_number VARCHAR(30),
    credit_limit DECIMAL(12, 2),
    opening_balance DECIMAL(12, 2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'Active',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createProductsTable = `
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    product_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    category VARCHAR(50),
    unit VARCHAR(20),
    description TEXT,
    status VARCHAR(20) DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createColoursTable = `
CREATE TABLE IF NOT EXISTS colours (
    id SERIAL PRIMARY KEY,
    colour_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    hex_code VARCHAR(7),
    description TEXT,
    status VARCHAR(20) DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createLocationsTable = `
CREATE TABLE IF NOT EXISTS locations (
    id SERIAL PRIMARY KEY,
    location_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    type VARCHAR(50),
    address TEXT,
    description TEXT,
    status VARCHAR(20) DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createDesignsTable = `
CREATE TABLE IF NOT EXISTS designs (
    id SERIAL PRIMARY KEY,
    design_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    saree_type VARCHAR(100),
    status VARCHAR(20) DEFAULT 'Active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createInventoryBatchesTable = `
CREATE TABLE IF NOT EXISTS inventory_batches (
    id SERIAL PRIMARY KEY,
    batch_reference VARCHAR(50) UNIQUE NOT NULL,
    product_id INTEGER NOT NULL REFERENCES products(id),
    colour_id INTEGER REFERENCES colours(id),
    supplier_id INTEGER REFERENCES suppliers(id),
    quantity_received DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    cost_per_unit DECIMAL(10, 2),
    total_cost DECIMAL(12, 2),
    current_location_id INTEGER REFERENCES locations(id),
    received_date DATE,
    status VARCHAR(20) DEFAULT 'Available',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createInventoryMovementsTable = `
CREATE TABLE IF NOT EXISTS inventory_movements (
    id SERIAL PRIMARY KEY,
    batch_id INTEGER NOT NULL REFERENCES inventory_batches(id),
    movement_type VARCHAR(50) NOT NULL,
    quantity DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    from_location_id INTEGER REFERENCES locations(id),
    to_location_id INTEGER REFERENCES locations(id),
    reference_type VARCHAR(50),
    reference_id INTEGER,
    created_by INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);
CREATE INDEX IF NOT EXISTS idx_inventory_movements_batch_id ON inventory_movements(batch_id);
CREATE INDEX IF NOT EXISTS idx_inventory_movements_type ON inventory_movements(movement_type);
`

const createRawSilkPurchasesTable = `
CREATE TABLE IF NOT EXISTS raw_silk_purchases (
    id SERIAL PRIMARY KEY,
    purchase_reference VARCHAR(50) UNIQUE NOT NULL,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(id),
    product_id INTEGER NOT NULL REFERENCES products(id),
    quantity DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    rate_per_unit DECIMAL(10, 2) NOT NULL,
    total_amount DECIMAL(12, 2) NOT NULL,
    purchase_date DATE NOT NULL,
    expected_delivery_date DATE,
    actual_delivery_date DATE,
    invoice_number VARCHAR(50),
    payment_status VARCHAR(20) DEFAULT 'Pending',
    amount_paid DECIMAL(12, 2) DEFAULT 0,
    amount_pending DECIMAL(12, 2),
    status VARCHAR(20) DEFAULT 'Confirmed',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createColouringBatchesTable = `
CREATE TABLE IF NOT EXISTS colouring_batches (
    id SERIAL PRIMARY KEY,
    colouring_reference VARCHAR(50) UNIQUE NOT NULL,
    raw_silk_batch_id INTEGER NOT NULL REFERENCES inventory_batches(id),
    colour_id INTEGER NOT NULL REFERENCES colours(id),
    quantity_sent DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    colour_factory_supplier_id INTEGER NOT NULL REFERENCES suppliers(id),
    date_sent DATE NOT NULL,
    expected_return_date DATE,
    date_received DATE,
    quantity_returned DECIMAL(10, 2),
    wastage_quantity DECIMAL(10, 2) DEFAULT 0,
    colouring_charges DECIMAL(10, 2),
    transportation_cost DECIMAL(10, 2) DEFAULT 0,
    total_processing_cost DECIMAL(10, 2),
    status VARCHAR(20) DEFAULT 'Sent',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createDoliBatchesTable = `
CREATE TABLE IF NOT EXISTS doli_batches (
    id SERIAL PRIMARY KEY,
    doli_reference VARCHAR(50) UNIQUE NOT NULL,
    source_coloured_silk_batch_id INTEGER NOT NULL REFERENCES inventory_batches(id),
    colour_id INTEGER NOT NULL REFERENCES colours(id),
    quantity_created DECIMAL(10, 2) NOT NULL,
    saree_capacity INTEGER,
    weight DECIMAL(10, 2),
    warping_charges DECIMAL(10, 2),
    warping_date DATE NOT NULL,
    current_location_id INTEGER REFERENCES locations(id),
    status VARCHAR(20) DEFAULT 'Available',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createWindingBatchesTable = `
CREATE TABLE IF NOT EXISTS winding_batches (
    id SERIAL PRIMARY KEY,
    winding_reference VARCHAR(50) UNIQUE NOT NULL,
    source_coloured_silk_batch_id INTEGER NOT NULL REFERENCES inventory_batches(id),
    colour_id INTEGER NOT NULL REFERENCES colours(id),
    quantity_wound DECIMAL(10, 2) NOT NULL,
    bobbins_produced DECIMAL(10, 2),
    pirns_produced DECIMAL(10, 2),
    winding_charges DECIMAL(10, 2),
    machine_operator VARCHAR(100),
    winding_date DATE NOT NULL,
    status VARCHAR(20) DEFAULT 'Completed',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createBobbiInventoryTable = `
CREATE TABLE IF NOT EXISTS bobbin_inventory (
    id SERIAL PRIMARY KEY,
    bobbin_batch_reference VARCHAR(50) UNIQUE NOT NULL,
    colour_id INTEGER NOT NULL REFERENCES colours(id),
    source_winding_batch_id INTEGER REFERENCES winding_batches(id),
    quantity DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    current_location_id INTEGER REFERENCES locations(id),
    status VARCHAR(20) DEFAULT 'Available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createSupplierTransactionsTable = `
CREATE TABLE IF NOT EXISTS supplier_transactions (
    id SERIAL PRIMARY KEY,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(id),
    transaction_type VARCHAR(50),
    description VARCHAR(300),
    amount DECIMAL(12, 2) NOT NULL,
    reference_type VARCHAR(50),
    reference_id INTEGER,
    transaction_date DATE NOT NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

const createAuditLogsTable = `
CREATE TABLE IF NOT EXISTS audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    action VARCHAR(100),
    entity_type VARCHAR(50),
    entity_id INTEGER,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
