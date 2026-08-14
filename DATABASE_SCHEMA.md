# Weaving Business Management System - Database Schema

---

## Database Overview

This document defines the complete PostgreSQL database schema for both Phase 1 and Phase 2.

**Key Principles:**
1. Immutable transaction ledgers (no direct updates to quantities)
2. Foreign key constraints for data integrity
3. Audit trail for all critical operations
4. Proper indexing for performance

---

---

# PHASE 1: CORE TABLES

---

## 1. Users & Authentication

### users
```sql
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(200),
    phone VARCHAR(20),
    role_id INTEGER REFERENCES roles(role_id),
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive', 'Locked')),
    last_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role_id ON users(role_id);
```

### roles
```sql
CREATE TABLE roles (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default roles
INSERT INTO roles (role_name, description) VALUES
('Admin', 'Full system access'),
('Manager', 'Inventory and production management'),
('Accountant', 'Finance and payment tracking'),
('Weaver', 'Weaver app access - Phase 2+');
```

### permissions (For future use)
```sql
CREATE TABLE permissions (
    permission_id SERIAL PRIMARY KEY,
    role_id INTEGER REFERENCES roles(role_id),
    module VARCHAR(50),
    action VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(role_id, module, action)
);
```

---

## 2. Master Data

### suppliers
```sql
CREATE TABLE suppliers (
    supplier_id SERIAL PRIMARY KEY,
    supplier_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(150),
    address TEXT,
    city VARCHAR(50),
    state VARCHAR(50),
    pincode VARCHAR(10),
    material_type VARCHAR(100),  -- 'Raw Silk', 'Colour Factory', 'Warping', 'Winding', etc.
    payment_terms VARCHAR(100),  -- e.g., 'Net 30', '50% Advance'
    bank_account_number VARCHAR(50),
    bank_name VARCHAR(100),
    ifsc_code VARCHAR(20),
    opening_balance DECIMAL(12, 2) DEFAULT 0,  -- Amount owed to supplier initially
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    notes TEXT,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_suppliers_code ON suppliers(supplier_code);
CREATE INDEX idx_suppliers_material_type ON suppliers(material_type);
CREATE INDEX idx_suppliers_status ON suppliers(status);
```

### weavers
```sql
CREATE TABLE weavers (
    weaver_id SERIAL PRIMARY KEY,
    weaver_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(150),
    address TEXT,
    village VARCHAR(100),
    city VARCHAR(50),
    state VARCHAR(50),
    joining_date DATE NOT NULL,
    bank_account_number VARCHAR(50),
    bank_name VARCHAR(100),
    ifsc_code VARCHAR(20),
    opening_balance DECIMAL(12, 2) DEFAULT 0,  -- Amount owed to weaver initially
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive', 'On Leave')),
    notes TEXT,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_weavers_code ON weavers(weaver_code);
CREATE INDEX idx_weavers_status ON weavers(status);
```

### buyers
```sql
CREATE TABLE buyers (
    buyer_id SERIAL PRIMARY KEY,
    buyer_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    business_name VARCHAR(200),
    phone VARCHAR(20),
    email VARCHAR(150),
    address TEXT,
    city VARCHAR(50),
    state VARCHAR(50),
    pincode VARCHAR(10),
    gst_number VARCHAR(30),
    credit_limit DECIMAL(12, 2),
    payment_terms VARCHAR(100),  -- e.g., 'Net 30', 'COD'
    opening_balance DECIMAL(12, 2) DEFAULT 0,  -- Amount buyer owes initially
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive', 'Blocked')),
    notes TEXT,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_buyers_code ON buyers(buyer_code);
CREATE INDEX idx_buyers_status ON buyers(status);
```

### products
```sql
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,  -- e.g., 'Kachha Silk', 'Coloured Silk', 'Doli', 'Bobbins'
    category VARCHAR(50),  -- 'Raw Material', 'Processed', 'Finished'
    unit VARCHAR(20),  -- 'kg', 'piece', 'unit', etc.
    description TEXT,
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (product_code, name, category, unit) VALUES
('P001', 'Kachha Silk', 'Raw Material', 'kg'),
('P002', 'Coloured Silk', 'Processed', 'kg'),
('P003', 'Doli', 'Processed', 'unit'),
('P004', 'Bobbins', 'Processed', 'unit'),
('P005', 'Pirns', 'Processed', 'unit'),
('P006', 'Saree', 'Finished', 'piece');

CREATE INDEX idx_products_code ON products(product_code);
CREATE INDEX idx_products_category ON products(category);
```

### colours
```sql
CREATE TABLE colours (
    colour_id SERIAL PRIMARY KEY,
    colour_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,  -- e.g., 'Red', 'Blue', 'Green', 'Maroon'
    hex_code VARCHAR(7),  -- For UI display, e.g., '#FF0000'
    description TEXT,
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO colours (colour_code, name, hex_code) VALUES
('C001', 'Red', '#FF0000'),
('C002', 'Blue', '#0000FF'),
('C003', 'Green', '#00AA00'),
('C004', 'Maroon', '#800000'),
('C005', 'Yellow', '#FFFF00'),
('C006', 'Black', '#000000'),
('C007', 'White', '#FFFFFF');

CREATE INDEX idx_colours_code ON colours(colour_code);
```

### locations
```sql
CREATE TABLE locations (
    location_id SERIAL PRIMARY KEY,
    location_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,  -- e.g., 'Main Warehouse', 'Colour Factory', 'Weaver - Ravi'
    type VARCHAR(50),  -- 'Warehouse', 'Supplier', 'Weaver', 'Colour Factory', 'Processing', 'Finished Goods'
    address TEXT,
    description TEXT,
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO locations (location_code, name, type) VALUES
('L001', 'Main Warehouse', 'Warehouse'),
('L002', 'Colour Factory', 'Processing'),
('L003', 'Warping Unit', 'Processing'),
('L004', 'Winding Unit', 'Processing'),
('L005', 'Finished Goods Store', 'Finished Goods');

CREATE INDEX idx_locations_code ON locations(location_code);
CREATE INDEX idx_locations_type ON locations(type);
```

### designs
```sql
CREATE TABLE designs (
    design_id SERIAL PRIMARY KEY,
    design_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,  -- e.g., 'Design A', 'Designer Saree 001'
    description TEXT,
    saree_type VARCHAR(100),  -- e.g., 'Regular', 'Premium', 'Wedding'
    image_url VARCHAR(500),
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_designs_code ON designs(design_code);
```

---

## 3. Inventory System (Core)

### inventory_batches
```sql
CREATE TABLE inventory_batches (
    batch_id SERIAL PRIMARY KEY,
    batch_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'KS-2026-001', 'RED-001'
    product_id INTEGER NOT NULL REFERENCES products(product_id),
    colour_id INTEGER REFERENCES colours(colour_id),  -- NULL for non-coloured items
    
    -- Source information
    supplier_id INTEGER REFERENCES suppliers(supplier_id),  -- For raw materials
    source_processing_id INTEGER,  -- For processed materials (links to colouring_batches, etc.)
    
    -- Physical details
    quantity_received DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    cost_per_unit DECIMAL(10, 2),
    total_cost DECIMAL(12, 2),
    
    -- Location and tracking
    current_location_id INTEGER REFERENCES locations(location_id),
    received_date DATE,
    
    -- Status
    status VARCHAR(20) CHECK (status IN ('Available', 'Reserved', 'In-Process', 'Damaged', 'Consumed')),
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_inventory_batches_reference ON inventory_batches(batch_reference);
CREATE INDEX idx_inventory_batches_product_id ON inventory_batches(product_id);
CREATE INDEX idx_inventory_batches_colour_id ON inventory_batches(colour_id);
CREATE INDEX idx_inventory_batches_supplier_id ON inventory_batches(supplier_id);
CREATE INDEX idx_inventory_batches_location_id ON inventory_batches(current_location_id);
CREATE INDEX idx_inventory_batches_status ON inventory_batches(status);
```

### inventory_movements (IMMUTABLE LEDGER)
```sql
CREATE TABLE inventory_movements (
    movement_id SERIAL PRIMARY KEY,
    batch_id INTEGER NOT NULL REFERENCES inventory_batches(batch_id),
    
    -- Movement details
    movement_type VARCHAR(50) NOT NULL,  
    -- e.g., 'Purchase', 'Sent to Colour', 'Received from Colour', 
    -- 'Warping', 'Winding', 'Issued to Weaver', 'Returned from Weaver', 'Damage', 'Adjustment'
    
    quantity DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    
    -- Location tracking
    from_location_id INTEGER REFERENCES locations(location_id),
    to_location_id INTEGER REFERENCES locations(location_id),
    
    -- Reference tracking
    reference_type VARCHAR(50),  
    -- e.g., 'Purchase', 'Colouring', 'Production', 'Material Issue'
    reference_id INTEGER,  -- Links to purchase_id, colouring_id, production_id, etc.
    
    -- Audit
    created_by INTEGER NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

-- CRITICAL: This table is IMMUTABLE (no updates, only inserts)
-- Create trigger to prevent updates:
-- CREATE TRIGGER prevent_inventory_movement_update 
--    BEFORE UPDATE ON inventory_movements 
--    FOR EACH ROW EXECUTE FUNCTION raise_immutable_error();

CREATE INDEX idx_inventory_movements_batch_id ON inventory_movements(batch_id);
CREATE INDEX idx_inventory_movements_type ON inventory_movements(movement_type);
CREATE INDEX idx_inventory_movements_created_at ON inventory_movements(created_at);
CREATE INDEX idx_inventory_movements_from_location ON inventory_movements(from_location_id);
CREATE INDEX idx_inventory_movements_to_location ON inventory_movements(to_location_id);
CREATE INDEX idx_inventory_movements_reference ON inventory_movements(reference_type, reference_id);
```

### audit_logs (For compliance and debugging)
```sql
CREATE TABLE audit_logs (
    audit_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(user_id),
    action VARCHAR(100),  -- e.g., 'Created Purchase', 'Issued Material', 'Recorded Payment'
    entity_type VARCHAR(50),  -- e.g., 'Purchase', 'Material Issue', 'Invoice'
    entity_id INTEGER,
    old_values JSONB,  -- For updates
    new_values JSONB,  -- For updates
    ip_address VARCHAR(45),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_entity ON audit_logs(entity_type, entity_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);
```

---

## 4. Raw Silk Purchase

### raw_silk_purchases
```sql
CREATE TABLE raw_silk_purchases (
    purchase_id SERIAL PRIMARY KEY,
    purchase_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'PUR-2026-001'
    supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id),
    product_id INTEGER NOT NULL REFERENCES products(product_id),  -- Usually Kachha Silk
    
    -- Material details
    quantity DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    rate_per_unit DECIMAL(10, 2) NOT NULL,
    total_amount DECIMAL(12, 2) NOT NULL,
    
    -- Dates
    purchase_date DATE NOT NULL,
    expected_delivery_date DATE,
    actual_delivery_date DATE,
    
    -- Invoice and payment
    invoice_number VARCHAR(50),
    payment_status VARCHAR(20) CHECK (payment_status IN ('Pending', 'Partial', 'Paid')),
    amount_paid DECIMAL(12, 2) DEFAULT 0,
    amount_pending DECIMAL(12, 2),
    
    -- Status
    status VARCHAR(20) CHECK (status IN ('Draft', 'Confirmed', 'Received', 'Partially Received', 'Cancelled')),
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_raw_silk_purchases_reference ON raw_silk_purchases(purchase_reference);
CREATE INDEX idx_raw_silk_purchases_supplier_id ON raw_silk_purchases(supplier_id);
CREATE INDEX idx_raw_silk_purchases_status ON raw_silk_purchases(status);
CREATE INDEX idx_raw_silk_purchases_payment_status ON raw_silk_purchases(payment_status);
```

---

## 5. Colouring Process

### colouring_batches
```sql
CREATE TABLE colouring_batches (
    colouring_id SERIAL PRIMARY KEY,
    colouring_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'COL-2026-001'
    raw_silk_batch_id INTEGER NOT NULL REFERENCES inventory_batches(batch_id),
    
    -- Colouring details
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    quantity_sent DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    
    -- Colour factory
    colour_factory_supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id),
    
    -- Dates
    date_sent DATE NOT NULL,
    expected_return_date DATE,
    date_received DATE,
    
    -- Quantity and cost
    quantity_returned DECIMAL(10, 2),
    wastage_quantity DECIMAL(10, 2) DEFAULT 0,
    colouring_charges DECIMAL(10, 2),
    transportation_cost DECIMAL(10, 2) DEFAULT 0,
    total_processing_cost DECIMAL(10, 2),
    
    -- Status
    status VARCHAR(20) CHECK (status IN ('Pending', 'Sent', 'Processing', 'Received', 'Partially Received', 'Cancelled')),
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_colouring_batches_reference ON colouring_batches(colouring_reference);
CREATE INDEX idx_colouring_batches_raw_silk_batch_id ON colouring_batches(raw_silk_batch_id);
CREATE INDEX idx_colouring_batches_colour_id ON colouring_batches(colour_id);
CREATE INDEX idx_colouring_batches_status ON colouring_batches(status);
```

---

## 6. Coloured Silk Stock

### coloured_silk_stock
```sql
-- NOTE: This is tracked via inventory_batches with product_id = Coloured Silk
-- This table is optional - can use inventory_batches view instead
-- If used separately, maintain consistency with inventory_batches

CREATE TABLE coloured_silk_stock (
    coloured_silk_id SERIAL PRIMARY KEY,
    inventory_batch_id INTEGER REFERENCES inventory_batches(batch_id),
    
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    source_raw_silk_batch_id INTEGER REFERENCES inventory_batches(batch_id),
    source_colouring_batch_id INTEGER REFERENCES colouring_batches(colouring_id),
    
    quantity DECIMAL(10, 2) NOT NULL,
    unit VARCHAR(20),
    cost DECIMAL(12, 2),
    received_date DATE,
    current_location_id INTEGER REFERENCES locations(location_id),
    
    status VARCHAR(20) CHECK (status IN ('Available', 'Reserved', 'In-Use', 'Damaged')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Consider: Is this table necessary? Or use inventory_batches with filters?
-- For now, keeping both for clarity
```

---

## 7. Warping / Doli Management

### doli_batches
```sql
CREATE TABLE doli_batches (
    doli_id SERIAL PRIMARY KEY,
    doli_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'DOLI-2026-00042'
    source_coloured_silk_batch_id INTEGER NOT NULL REFERENCES inventory_batches(batch_id),
    
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    
    -- Doli details
    quantity_created DECIMAL(10, 2) NOT NULL,  -- Number of doli units created
    saree_capacity INTEGER,  -- How many sarees can be made from each doli
    weight DECIMAL(10, 2),  -- Weight of doli
    
    -- Processing details
    warping_charges DECIMAL(10, 2),
    warping_date DATE NOT NULL,
    
    -- Location and status
    current_location_id INTEGER REFERENCES locations(location_id),
    status VARCHAR(20) CHECK (status IN ('Available', 'Reserved', 'Issued', 'Returned', 'Damaged')),
    
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_doli_batches_reference ON doli_batches(doli_reference);
CREATE INDEX idx_doli_batches_colour_id ON doli_batches(colour_id);
CREATE INDEX idx_doli_batches_status ON doli_batches(status);
```

---

## 8. Bobbins / Pirns Management

### winding_batches
```sql
CREATE TABLE winding_batches (
    winding_id SERIAL PRIMARY KEY,
    winding_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'WIND-2026-001'
    source_coloured_silk_batch_id INTEGER NOT NULL REFERENCES inventory_batches(batch_id),
    
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    
    -- Winding details
    quantity_wound DECIMAL(10, 2) NOT NULL,
    bobbins_produced DECIMAL(10, 2),
    pirns_produced DECIMAL(10, 2),
    
    -- Processing details
    winding_charges DECIMAL(10, 2),
    machine_operator VARCHAR(100),
    winding_date DATE NOT NULL,
    
    status VARCHAR(20) CHECK (status IN ('Completed', 'In-Use', 'Damaged')),
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_winding_batches_reference ON winding_batches(winding_reference);
CREATE INDEX idx_winding_batches_colour_id ON winding_batches(colour_id);
```

### bobbin_inventory
```sql
CREATE TABLE bobbin_inventory (
    bobbin_id SERIAL PRIMARY KEY,
    bobbin_batch_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'BOB-BLUE-001'
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    
    source_winding_batch_id INTEGER REFERENCES winding_batches(winding_id),
    
    quantity DECIMAL(10, 2) NOT NULL,  -- Number of bobbins
    unit VARCHAR(20),  -- Usually 'piece'
    
    current_location_id INTEGER REFERENCES locations(location_id),
    status VARCHAR(20) CHECK (status IN ('Available', 'Issued', 'Returned', 'Damaged')),
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bobbin_inventory_colour_id ON bobbin_inventory(colour_id);
CREATE INDEX idx_bobbin_inventory_status ON bobbin_inventory(status);
```

---

## 9. Transactions Tracking (Phase 1)

### supplier_transactions
```sql
CREATE TABLE supplier_transactions (
    transaction_id SERIAL PRIMARY KEY,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id),
    
    transaction_type VARCHAR(50),  
    -- e.g., 'Purchase', 'Payment', 'Adjustment', 'Credit Note'
    
    description VARCHAR(300),
    amount DECIMAL(12, 2) NOT NULL,
    reference_type VARCHAR(50),  -- 'Purchase ID', 'Payment ID'
    reference_id INTEGER,
    
    transaction_date DATE NOT NULL,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_supplier_transactions_supplier_id ON supplier_transactions(supplier_id);
CREATE INDEX idx_supplier_transactions_type ON supplier_transactions(transaction_type);
CREATE INDEX idx_supplier_transactions_date ON supplier_transactions(transaction_date);
```

---

---

# PHASE 2: ADDITIONAL TABLES

---

## 1. Production Management

### production_orders
```sql
CREATE TABLE production_orders (
    production_id SERIAL PRIMARY KEY,
    po_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'PO-2026-00042'
    
    weaver_id INTEGER NOT NULL REFERENCES weavers(weaver_id),
    design_id INTEGER NOT NULL REFERENCES designs(design_id),
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    
    -- Material requirements
    doli_required DECIMAL(10, 2),
    bobbins_required DECIMAL(10, 2),
    
    -- Production details
    expected_sarees INTEGER NOT NULL,
    rate_per_saree DECIMAL(10, 2) NOT NULL,
    total_expected_payment DECIMAL(12, 2),
    
    -- Dates
    issued_date DATE NOT NULL,
    expected_completion_date DATE,
    actual_completion_date DATE,
    
    -- Status
    status VARCHAR(20) CHECK (status IN ('Pending', 'Issued', 'In-Progress', 'Completed', 'Cancelled')),
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_production_orders_reference ON production_orders(po_reference);
CREATE INDEX idx_production_orders_weaver_id ON production_orders(weaver_id);
CREATE INDEX idx_production_orders_design_id ON production_orders(design_id);
CREATE INDEX idx_production_orders_status ON production_orders(status);
```

### material_issues
```sql
CREATE TABLE material_issues (
    issue_id SERIAL PRIMARY KEY,
    issue_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'ISS-2026-001'
    
    production_order_id INTEGER NOT NULL REFERENCES production_orders(production_id),
    weaver_id INTEGER NOT NULL REFERENCES weavers(weaver_id),
    
    -- Material issued
    doli_issued DECIMAL(10, 2),
    bobbins_issued DECIMAL(10, 2),
    other_material TEXT,
    
    issued_date DATE NOT NULL,
    expected_return_date DATE,
    
    -- Status
    status VARCHAR(20) CHECK (status IN ('Pending Return', 'Returned', 'Partially Returned')),
    return_date DATE,
    
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_material_issues_reference ON material_issues(issue_reference);
CREATE INDEX idx_material_issues_production_order_id ON material_issues(production_order_id);
CREATE INDEX idx_material_issues_weaver_id ON material_issues(weaver_id);
CREATE INDEX idx_material_issues_status ON material_issues(status);
```

### material_returns
```sql
CREATE TABLE material_returns (
    return_id SERIAL PRIMARY KEY,
    return_reference VARCHAR(50) UNIQUE NOT NULL,
    
    material_issue_id INTEGER NOT NULL REFERENCES material_issues(issue_id),
    weaver_id INTEGER NOT NULL REFERENCES weavers(weaver_id),
    
    -- Returned quantities
    doli_returned DECIMAL(10, 2),
    bobbins_returned DECIMAL(10, 2),
    
    -- Damage/loss
    doli_damaged DECIMAL(10, 2) DEFAULT 0,
    bobbins_damaged DECIMAL(10, 2) DEFAULT 0,
    
    return_date DATE NOT NULL,
    notes TEXT,
    
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_material_returns_issue_id ON material_returns(material_issue_id);
```

### production_outputs
```sql
CREATE TABLE production_outputs (
    output_id SERIAL PRIMARY KEY,
    output_reference VARCHAR(50) UNIQUE NOT NULL,
    
    production_order_id INTEGER NOT NULL REFERENCES production_orders(production_id),
    weaver_id INTEGER NOT NULL REFERENCES weavers(weaver_id),
    design_id INTEGER NOT NULL REFERENCES designs(design_id),
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    
    -- Production details
    sarees_produced INTEGER NOT NULL,
    sarees_approved INTEGER NOT NULL,
    sarees_rejected INTEGER DEFAULT 0,
    quality_notes TEXT,
    
    completion_date DATE NOT NULL,
    
    -- Payment
    payment_amount DECIMAL(12, 2),  -- sarees_approved × rate
    
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_production_outputs_reference ON production_outputs(output_reference);
CREATE INDEX idx_production_outputs_production_order_id ON production_outputs(production_order_id);
CREATE INDEX idx_production_outputs_weaver_id ON production_outputs(weaver_id);
```

---

## 2. Saree Management (Phase 2)

### sarees
```sql
CREATE TABLE sarees (
    saree_id SERIAL PRIMARY KEY,
    saree_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'SR-2026-00001'
    
    design_id INTEGER NOT NULL REFERENCES designs(design_id),
    colour_id INTEGER NOT NULL REFERENCES colours(colour_id),
    saree_type VARCHAR(100),  -- e.g., 'Regular', 'Premium'
    
    weaver_id INTEGER REFERENCES weavers(weaver_id),
    production_order_id INTEGER REFERENCES production_orders(production_id),
    production_output_id INTEGER REFERENCES production_outputs(output_id),
    
    -- Costing
    production_cost DECIMAL(10, 2),
    selling_price DECIMAL(10, 2),
    
    production_date DATE,
    current_location_id INTEGER REFERENCES locations(location_id),
    
    -- Status
    status VARCHAR(20) CHECK (status IN ('In Stock', 'Reserved', 'Sold', 'Damaged', 'Returned', 'Rejected')),
    
    notes TEXT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sarees_reference ON sarees(saree_reference);
CREATE INDEX idx_sarees_design_id ON sarees(design_id);
CREATE INDEX idx_sarees_colour_id ON sarees(colour_id);
CREATE INDEX idx_sarees_weaver_id ON sarees(weaver_id);
CREATE INDEX idx_sarees_status ON sarees(status);
```

### saree_movements
```sql
CREATE TABLE saree_movements (
    movement_id SERIAL PRIMARY KEY,
    saree_id INTEGER NOT NULL REFERENCES sarees(saree_id),
    
    movement_type VARCHAR(50),  
    -- e.g., 'Produced', 'Sold', 'Damaged', 'Returned', 'Adjusted'
    
    from_location_id INTEGER REFERENCES locations(location_id),
    to_location_id INTEGER REFERENCES locations(location_id),
    
    reference_type VARCHAR(50),  -- 'Invoice', 'Damage Report', etc.
    reference_id INTEGER,
    
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_saree_movements_saree_id ON saree_movements(saree_id);
CREATE INDEX idx_saree_movements_type ON saree_movements(movement_type);
```

---

## 3. Sales & Billing (Phase 2)

### invoices
```sql
CREATE TABLE invoices (
    invoice_id SERIAL PRIMARY KEY,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'INV-2026-001'
    
    buyer_id INTEGER NOT NULL REFERENCES buyers(buyer_id),
    invoice_date DATE NOT NULL,
    sale_date DATE,
    
    -- Amounts
    subtotal DECIMAL(12, 2) NOT NULL,
    discount_amount DECIMAL(12, 2) DEFAULT 0,
    discount_percentage DECIMAL(5, 2) DEFAULT 0,
    tax_amount DECIMAL(12, 2) DEFAULT 0,
    total_amount DECIMAL(12, 2) NOT NULL,
    
    -- Payment tracking
    payment_status VARCHAR(20) CHECK (payment_status IN ('Unpaid', 'Partial', 'Paid')),
    amount_received DECIMAL(12, 2) DEFAULT 0,
    amount_pending DECIMAL(12, 2),
    
    notes TEXT,
    
    -- Audit
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_invoices_number ON invoices(invoice_number);
CREATE INDEX idx_invoices_buyer_id ON invoices(buyer_id);
CREATE INDEX idx_invoices_invoice_date ON invoices(invoice_date);
CREATE INDEX idx_invoices_payment_status ON invoices(payment_status);
```

### invoice_items
```sql
CREATE TABLE invoice_items (
    line_item_id SERIAL PRIMARY KEY,
    invoice_id INTEGER NOT NULL REFERENCES invoices(invoice_id),
    saree_id INTEGER NOT NULL REFERENCES sarees(saree_id),
    
    quantity INTEGER NOT NULL,
    rate DECIMAL(10, 2) NOT NULL,
    amount DECIMAL(12, 2) NOT NULL,  -- quantity × rate
    
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_invoice_items_invoice_id ON invoice_items(invoice_id);
CREATE INDEX idx_invoice_items_saree_id ON invoice_items(saree_id);
```

### sale_payments
```sql
CREATE TABLE sale_payments (
    payment_id SERIAL PRIMARY KEY,
    payment_reference VARCHAR(50) UNIQUE NOT NULL,
    
    invoice_id INTEGER NOT NULL REFERENCES invoices(invoice_id),
    buyer_id INTEGER NOT NULL REFERENCES buyers(buyer_id),
    
    payment_amount DECIMAL(12, 2) NOT NULL,
    payment_date DATE NOT NULL,
    payment_mode VARCHAR(50),  -- 'Cash', 'Check', 'Bank Transfer', 'Digital Wallet'
    reference_number VARCHAR(100),  -- Cheque number, transaction ID, etc.
    
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_sale_payments_invoice_id ON sale_payments(invoice_id);
CREATE INDEX idx_sale_payments_buyer_id ON sale_payments(buyer_id);
CREATE INDEX idx_sale_payments_payment_date ON sale_payments(payment_date);
```

---

## 4. Weaver Management & Ledger (Phase 2)

### weaver_transactions
```sql
CREATE TABLE weaver_transactions (
    transaction_id SERIAL PRIMARY KEY,
    weaver_id INTEGER NOT NULL REFERENCES weavers(weaver_id),
    
    transaction_type VARCHAR(50),  
    -- e.g., 'Opening Balance', 'Work Issued', 'Work Payment', 'Advance Paid', 'Adjustment'
    
    description VARCHAR(300),
    amount DECIMAL(12, 2) NOT NULL,
    
    reference_type VARCHAR(50),  -- e.g., 'Production Order', 'Material Issue', 'Payment'
    reference_id INTEGER,
    
    transaction_date DATE NOT NULL,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_weaver_transactions_weaver_id ON weaver_transactions(weaver_id);
CREATE INDEX idx_weaver_transactions_type ON weaver_transactions(transaction_type);
CREATE INDEX idx_weaver_transactions_date ON weaver_transactions(transaction_date);
```

### weaver_payments
```sql
CREATE TABLE weaver_payments (
    payment_id SERIAL PRIMARY KEY,
    payment_reference VARCHAR(50) UNIQUE NOT NULL,
    
    weaver_id INTEGER NOT NULL REFERENCES weavers(weaver_id),
    production_order_id INTEGER REFERENCES production_orders(production_id),
    
    payment_amount DECIMAL(12, 2) NOT NULL,
    payment_date DATE NOT NULL,
    payment_mode VARCHAR(50),  -- 'Cash', 'Check', 'Bank Transfer'
    reference_number VARCHAR(100),
    
    -- Adjustments
    advance_deducted DECIMAL(12, 2) DEFAULT 0,
    net_amount_paid DECIMAL(12, 2),
    
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_weaver_payments_weaver_id ON weaver_payments(weaver_id);
CREATE INDEX idx_weaver_payments_payment_date ON weaver_payments(payment_date);
```

---

## 5. Buyer Ledger (Phase 2)

### buyer_transactions
```sql
CREATE TABLE buyer_transactions (
    transaction_id SERIAL PRIMARY KEY,
    buyer_id INTEGER NOT NULL REFERENCES buyers(buyer_id),
    
    transaction_type VARCHAR(50),  
    -- e.g., 'Opening Balance', 'Invoice', 'Payment', 'Adjustment', 'Credit Note'
    
    description VARCHAR(300),
    amount DECIMAL(12, 2) NOT NULL,
    
    reference_type VARCHAR(50),  -- e.g., 'Invoice ID', 'Payment ID'
    reference_id INTEGER,
    
    transaction_date DATE NOT NULL,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_buyer_transactions_buyer_id ON buyer_transactions(buyer_id);
CREATE INDEX idx_buyer_transactions_type ON buyer_transactions(transaction_type);
CREATE INDEX idx_buyer_transactions_date ON buyer_transactions(transaction_date);
```

---

## 6. Income & Expenses (Phase 2)

### expense_categories
```sql
CREATE TABLE expense_categories (
    category_id SERIAL PRIMARY KEY,
    category_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) CHECK (status IN ('Active', 'Inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO expense_categories (category_code, name) VALUES
('E001', 'Raw Material Purchase'),
('E002', 'Colouring Charges'),
('E003', 'Warping Charges'),
('E004', 'Winding Charges'),
('E005', 'Weaver Payments'),
('E006', 'Transport'),
('E007', 'Electricity'),
('E008', 'Machine Maintenance'),
('E009', 'Labour/Salary'),
('E010', 'Rent'),
('E011', 'Packaging'),
('E012', 'Miscellaneous');
```

### expenses
```sql
CREATE TABLE expenses (
    expense_id SERIAL PRIMARY KEY,
    expense_reference VARCHAR(50) UNIQUE NOT NULL,  -- e.g., 'EXP-2026-001'
    
    expense_category_id INTEGER NOT NULL REFERENCES expense_categories(category_id),
    supplier_id INTEGER REFERENCES suppliers(supplier_id),  -- Optional, for vendor expenses
    
    description VARCHAR(300) NOT NULL,
    amount DECIMAL(12, 2) NOT NULL,
    expense_date DATE NOT NULL,
    
    -- Link to business activity
    related_entity_type VARCHAR(50),  -- e.g., 'Purchase', 'Colouring', 'Production'
    related_entity_id INTEGER,
    
    -- Payment tracking
    payment_status VARCHAR(20) CHECK (payment_status IN ('Pending', 'Paid')),
    payment_mode VARCHAR(50),
    reference_number VARCHAR(100),
    
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_expenses_reference ON expenses(expense_reference);
CREATE INDEX idx_expenses_category_id ON expenses(expense_category_id);
CREATE INDEX idx_expenses_supplier_id ON expenses(supplier_id);
CREATE INDEX idx_expenses_expense_date ON expenses(expense_date);
CREATE INDEX idx_expenses_payment_status ON expenses(payment_status);
```

### income_transactions
```sql
CREATE TABLE income_transactions (
    income_id SERIAL PRIMARY KEY,
    income_reference VARCHAR(50) UNIQUE NOT NULL,
    
    income_type VARCHAR(50),  -- e.g., 'Saree Sales', 'Other Income'
    description VARCHAR(300),
    amount DECIMAL(12, 2) NOT NULL,
    
    reference_type VARCHAR(50),  -- e.g., 'Invoice ID'
    reference_id INTEGER,
    
    income_date DATE NOT NULL,
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_income_transactions_type ON income_transactions(income_type);
CREATE INDEX idx_income_transactions_date ON income_transactions(income_date);
```

---

## 7. Supplier Payments (Phase 2 - Enhanced)

### supplier_payments
```sql
CREATE TABLE supplier_payments (
    payment_id SERIAL PRIMARY KEY,
    payment_reference VARCHAR(50) UNIQUE NOT NULL,
    
    supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id),
    
    payment_amount DECIMAL(12, 2) NOT NULL,
    payment_date DATE NOT NULL,
    payment_mode VARCHAR(50),
    reference_number VARCHAR(100),
    
    -- Which transactions being paid for
    against_reference_type VARCHAR(50),  -- 'Purchase', 'Colouring', etc.
    against_reference_id INTEGER,  -- Purchase ID, Colouring ID, etc.
    
    created_by INTEGER REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    notes TEXT
);

CREATE INDEX idx_supplier_payments_supplier_id ON supplier_payments(supplier_id);
CREATE INDEX idx_supplier_payments_payment_date ON supplier_payments(payment_date);
```

---

---

# IMPORTANT DATABASE RULES

---

## 1. Immutability Rule
```
inventory_movements table MUST NEVER be updated or deleted.
This ensures complete audit trail.
```

## 2. Stock Calculation Rule
```
Current Stock of Batch = SUM(quantity) from inventory_movements 
                       WHERE batch_id = X 
                       AND movement_type NOT IN ('Damage', 'Loss')
```

## 3. Location Tracking Rule
```
Every material must have a location.
When issued to weaver, location = Weaver location
When at colour factory, location = Colour Factory location
This answers: WHERE IS MY MATERIAL?
```

## 4. Transaction Integrity
```
Every financial transaction must:
1. Have a reference ID (links to business activity)
2. Have an audit log entry
3. Be immutable (no updates, corrections via new transaction)
```

## 5. Payment Tracking Rule
```
Payment Status:
- Pending: No payment received
- Partial: Some payment received (amount_received < amount_pending)
- Paid: Full payment received (amount_pending = 0)
```

---

# VIEWS (Helpful Queries)

---

## View: Current Inventory Stock
```sql
CREATE VIEW v_inventory_stock AS
SELECT 
    b.batch_id,
    b.batch_reference,
    p.name as product_name,
    c.name as colour,
    l.name as location,
    SUM(CASE WHEN im.movement_type NOT IN ('Damage', 'Loss') THEN im.quantity ELSE 0 END) as current_quantity,
    b.unit,
    b.cost_per_unit,
    SUM(CASE WHEN im.movement_type NOT IN ('Damage', 'Loss') THEN im.quantity ELSE 0 END) * b.cost_per_unit as total_value
FROM inventory_batches b
LEFT JOIN inventory_movements im ON b.batch_id = im.batch_id
LEFT JOIN products p ON b.product_id = p.product_id
LEFT JOIN colours c ON b.colour_id = c.colour_id
LEFT JOIN locations l ON b.current_location_id = l.location_id
GROUP BY b.batch_id, b.batch_reference, p.name, c.name, l.name, b.unit, b.cost_per_unit;
```

## View: Weaver Ledger
```sql
CREATE VIEW v_weaver_ledger AS
SELECT 
    w.weaver_id,
    w.name,
    w.opening_balance,
    wt.transaction_type,
    wt.amount,
    wt.transaction_date,
    SUM(CASE WHEN wt.transaction_type IN ('Work Payment', 'Advance Paid') THEN -wt.amount 
             WHEN wt.transaction_type IN ('Opening Balance', 'Work Issued', 'Adjustment') THEN wt.amount
             ELSE 0 END) OVER (PARTITION BY w.weaver_id ORDER BY wt.transaction_date) as running_balance
FROM weavers w
LEFT JOIN weaver_transactions wt ON w.weaver_id = wt.weaver_id
ORDER BY w.weaver_id, wt.transaction_date;
```

## View: Buyer Ledger
```sql
CREATE VIEW v_buyer_ledger AS
SELECT 
    b.buyer_id,
    b.name,
    b.opening_balance,
    bt.transaction_type,
    bt.amount,
    bt.transaction_date,
    SUM(CASE WHEN bt.transaction_type IN ('Invoice') THEN bt.amount 
             WHEN bt.transaction_type IN ('Payment', 'Adjustment') THEN -bt.amount
             ELSE 0 END) OVER (PARTITION BY b.buyer_id ORDER BY bt.transaction_date) as running_balance
FROM buyers b
LEFT JOIN buyer_transactions bt ON b.buyer_id = bt.buyer_id
ORDER BY b.buyer_id, bt.transaction_date;
```

---

**Database Version**: 1.0  
**Last Updated**: August 14, 2026
