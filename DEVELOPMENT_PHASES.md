# Weaving Business Management System - Two Phase Development Plan

---

## Executive Summary
The project is split into **2 major phases** to ensure a solid foundation before building production capabilities.

- **Phase 1**: Foundation + Complete Inventory System
- **Phase 2**: Production + Finance + Reporting

---

---

# PHASE 1: Foundation & Inventory System
**Duration**: 6-8 weeks | **Focus**: Stabilize core data and inventory tracking

## Phase 1 Objectives
1. ✅ Complete technical foundation
2. ✅ Master data management (Suppliers, Weavers, Buyers, Products)
3. ✅ Full inventory tracking with movement history
4. ✅ Material journey visibility (Raw Silk → Colouring → Coloured Silk → Doli → Bobbins)
5. ✅ Dashboard with key inventory metrics

---

## Phase 1: Core Components

### 1.1 Technical Foundation

#### Authentication & Users
```
- User registration & login
- JWT token-based authentication
- Role-based access control (RBAC)
  - Admin/Owner
  - Manager/Staff
  - Accountant
- Session management
- Password security
```

#### Audit Trail System
```
- Immutable audit logs
- Track: User, Action, Timestamp, Changes
- Examples:
  - User created supplier
  - Inventory increased by 50 kg
  - Material location changed
```

---

### 1.2 Master Data Management

#### A. Supplier Master
```
Table: suppliers
├── supplier_id (PK)
├── name
├── phone
├── email
├── address
├── city
├── state
├── material_type (Raw Silk, Colour Factory, etc.)
├── payment_terms
├── bank_details
├── created_date
├── status (Active/Inactive)
└── notes

UI Screens:
- Supplier list with search/filter
- Add/Edit supplier
- Supplier details view
```

#### B. Weaver Master
```
Table: weavers
├── weaver_id (PK)
├── name
├── phone
├── email
├── address
├── village
├── joining_date
├── bank_details
├── opening_balance (for ledger)
├── status (Active/Inactive)
├── notes
└── created_date

UI Screens:
- Weaver list with search/filter
- Add/Edit weaver
- Weaver details view
```

#### C. Buyer/Customer Master
```
Table: buyers
├── buyer_id (PK)
├── name
├── business_name
├── phone
├── email
├── address
├── city
├── gst_number
├── credit_limit
├── payment_terms
├── opening_balance
├── status (Active/Inactive)
├── notes
└── created_date

UI Screens:
- Buyer list
- Add/Edit buyer
- Buyer details view
```

#### D. Product Master
```
Table: products
├── product_id (PK)
├── name (Kachha Silk, Coloured Silk, etc.)
├── category (Raw Material, Processed, Finished)
├── unit (kg, piece, etc.)
├── description
└── status

Table: colours
├── colour_id (PK)
├── name (Red, Blue, Green, etc.)
├── hex_code (for UI display)
└── status

Table: locations
├── location_id (PK)
├── name (Main Warehouse, Colour Factory, Weaver-Ravi, etc.)
├── type (Warehouse, Factory, Weaver, Processing)
└── description
```

#### E. Design/Saree Types Master (Phase 1 Basic)
```
Table: designs
├── design_id (PK)
├── name (Design A, Design B, etc.)
├── description
└── status
```

---

### 1.3 Inventory Management System

#### A. Core Inventory Structure

```
Table: inventory_batches
├── batch_id (PK)
├── product_id (FK)
├── colour_id (FK)
├── batch_reference (KS-2026-001, etc.)
├── source_supplier_id (FK) [for raw materials]
├── purchase_date
├── quantity_received
├── unit
├── rate
├── total_cost
├── current_location_id (FK)
├── status (Available, Reserved, In-Process, Damaged)
├── created_date
└── notes

Key Principle: 
This table NEVER gets quantity updated directly.
Instead, inventory_movements create the history.
Current stock = SUM of all movements for this batch
```

#### B. Inventory Movements (Immutable Ledger)
```
Table: inventory_movements
├── movement_id (PK)
├── batch_id (FK)
├── movement_type (Purchase, Issued, Received, Adjustment, etc.)
├── quantity
├── from_location_id (FK)
├── to_location_id (FK)
├── reference_id (Purchase ID, Processing ID, etc.)
├── created_date
├── created_by (User ID)
├── notes
└── audit_id (Links to audit trail)

Example Flow:
Date       | Batch      | Type          | Qty  | From       | To         | Balance
-----------|------------|---------------|------|------------|------------|----------
14-Aug     | KS-001     | Purchase      | +50  | -          | Warehouse  | 50
15-Aug     | KS-001     | Sent Colour   | -25  | Warehouse  | Colour Fac | 25
17-Aug     | RS-RED-001 | Received      | +24  | Colour Fac | Warehouse  | 24
18-Aug     | RS-RED-001 | Warping       | -5   | Warehouse  | Warping    | 19
```

---

### 1.4 Raw Silk Purchase Process

#### Transaction Flow
```
Supplier receives purchase order
         ↓
Raw Silk Batch Created
         ↓
Purchase Invoice Created
         ↓
Material Received (Quantity + Batch)
         ↓
Inventory Movement: +Quantity to Warehouse
         ↓
Payment Tracking (Supplier Payable)
```

#### Tables
```
Table: raw_silk_purchases
├── purchase_id (PK)
├── supplier_id (FK)
├── batch_reference (KS-2026-001)
├── quantity
├── unit
├── rate
├── total_amount
├── purchase_date
├── expected_delivery_date
├── actual_delivery_date
├── invoice_number
├── payment_status (Pending, Partial, Paid)
├── amount_paid
├── amount_pending
├── created_date
└── notes

Automatic Action on "Confirm Receipt":
- Create inventory_batch
- Create inventory_movement (+quantity, Warehouse)
- Create supplier_transaction (Payable)
```

#### UI Screens
```
- Raw Silk Purchase List
- Create Purchase Order
- Mark as Received
- View Purchase Details
- Purchase Invoice Print
```

---

### 1.5 Colouring Process

#### Transaction Flow
```
Coloured Silk Batch Created (from Raw Silk)
         ↓
Material Sent to Colour Factory
         ↓
Processing Record Created
         ↓
Material Received from Factory
         ↓
Inventory Movement: +Coloured Silk, -Raw Silk
         ↓
Colouring Charges Added to Expenses
```

#### Tables
```
Table: colouring_batches
├── colouring_id (PK)
├── raw_silk_batch_id (FK)
├── quantity_sent
├── quantity_returned
├── colour_id (FK)
├── colour_factory_id (FK - Supplier)
├── date_sent
├── expected_return_date
├── date_received
├── wastage_quantity
├── colouring_charges
├── transportation_cost
├── status (Pending, Sent, Processing, Received, Partially Received)
├── created_date
└── notes

Automatic Actions on Status Changes:
- "Sent": Inventory movement -Raw Silk from Warehouse to Colour Factory location
- "Received": Create Coloured Silk Batch + Inventory movement +Coloured Silk to Warehouse
             Create Expense entry for colouring charges
```

#### UI Screens
```
- Colouring Batch List
- Create Colouring Request
- Send to Factory
- Receive from Factory
- Colouring Cost Tracking
```

---

### 1.6 Coloured Silk Stock

#### Tables
```
Table: coloured_silk_stock
├── coloured_silk_id (PK)
├── colour_id (FK)
├── source_raw_silk_batch_id (FK)
├── source_colouring_batch_id (FK)
├── quantity
├── unit
├── current_location_id (FK)
├── received_date
├── cost (from colouring batch)
├── status (Available, Reserved, In-Use)
├── created_date
└── notes

Note: This is essentially an inventory_batch with special properties
Tracked via inventory_movements just like raw silk
```

---

### 1.7 Warping / Doli Management

#### Transaction Flow
```
Coloured Silk Selected
         ↓
Warping Process (Convert to Doli)
         ↓
Doli Batch Created
         ↓
Inventory Movement: -Coloured Silk, +Doli
         ↓
Warping Charges Added
```

#### Tables
```
Table: doli_batches
├── doli_id (PK)
├── source_coloured_silk_batch_id (FK)
├── source_colour_id (FK)
├── quantity_created (number of doris)
├── saree_capacity (how many sarees per doli)
├── doli_reference (DOLI-00042)
├── warping_charges
├── warping_date
├── current_location_id (FK)
├── status (Available, Reserved, Issued, Returned)
├── created_date
└── notes

Automatic Action on "Create Doli":
- Calculate: Coloured Silk consumed / Saree capacity
- Create inventory_movement: -Coloured Silk, to Doli
- Create inventory_movement: +Doli to Warehouse, from Warping
```

#### UI Screens
```
- Doli Creation Form
- Doli List/Stock
- Doli Details (capacity, colour, weight)
- Doli Movement History
```

---

### 1.8 Bobbins / Pirns Management

#### Transaction Flow
```
Coloured Silk Selected
         ↓
Winding Process
         ↓
Bobbins/Pirns Produced
         ↓
Inventory Movement: -Coloured Silk, +Bobbins
         ↓
Winding Charges Added
```

#### Tables
```
Table: winding_batches
├── winding_id (PK)
├── source_coloured_silk_batch_id (FK)
├── source_colour_id (FK)
├── quantity_wound
├── bobbins_produced
├── pirns_produced
├── winding_charges
├── machine_operator
├── winding_date
├── status (Completed, In-Use)
├── created_date
└── notes

Table: bobbin_inventory (Track as simple inventory items)
├── bobbin_id (PK)
├── colour_id (FK)
├── source_winding_batch_id (FK)
├── quantity (number of bobbins)
├── current_location_id (FK)
├── status (Available, Issued, Returned)
└── created_date
```

#### UI Screens
```
- Winding Record Form
- Bobbin/Pirn Stock
- Bobbin Stock by Colour
- Winding Charge Tracking
```

---

### 1.9 Dashboard (Phase 1)

#### Key Metrics Display
```
┌────────────────────────────────────────────┐
│     WEAVING BUSINESS DASHBOARD - PHASE 1   │
├────────────────────────────────────────────┤
│                                            │
│  INVENTORY OVERVIEW                        │
│  ├─ Raw Silk: 182 kg (₹XX,XXX)             │
│  ├─ Coloured Silk: 95 kg (₹XX,XXX)         │
│  ├─ Doli: 42 units (₹XX,XXX)               │
│  ├─ Bobbins: 850 units (₹XX,XXX)           │
│  └─ Total Inventory Value: ₹XX,XXX         │
│                                            │
│  LOCATIONS                                 │
│  ├─ Main Warehouse: ₹XX,XXX (32% utilization) │
│  ├─ Colour Factory: ₹XX,XXX (pending)     │
│  ├─ Warping Unit: ₹XX,XXX (in-process)    │
│  └─ Winding Unit: ₹XX,XXX (in-process)    │
│                                            │
│  PROCESSING STATUS                         │
│  ├─ Colouring Batches In-Process: 5       │
│  ├─ Expected Receipts: 3 (within 2 days)  │
│  └─ Overdue Batches: 1                    │
│                                            │
│  SUPPLIER METRICS                          │
│  ├─ Active Suppliers: 12                   │
│  ├─ Recent Purchases: 4 (this month)       │
│  └─ Outstanding Payables: ₹XX,XXX          │
│                                            │
└────────────────────────────────────────────┘
```

#### Dashboard Screens
```
1. Inventory Overview
   - Real-time stock levels
   - Visual inventory distribution by location
   - Stock value by category

2. Material Journey
   - Material currently at Colour Factory
   - Material under Warping
   - Material under Winding
   - Expected receipts

3. Supplier Dashboard
   - Purchase history (last 30 days)
   - Outstanding payments
   - Delivery performance

4. System Health
   - Pending transactions
   - Overdue processes
   - Data audit status
```

---

### 1.10 Reports (Phase 1 - Basic)

```
1. Inventory Stock Report
   - Current stock by material type
   - Stock by colour
   - Stock by location
   - Stock valuation

2. Material Location Report
   - Where is every material right now?
   - Location-wise inventory

3. Inventory Movement Report
   - Historical movements
   - In/Out by date range
   - Loss/Wastage analysis

4. Supplier Report
   - Total purchases per supplier
   - Outstanding amount
   - Supply history

5. Colouring Status Report
   - Batches sent
   - Batches in-process
   - Batches received
   - Overdue batches
```

---

## Phase 1: Database Schema (Core Tables)

```sql
-- Users & Auth
users
supplier_transactions
supplier_payments

-- Masters
suppliers
weavers
buyers
products
colours
locations
designs

-- Inventory System (Core)
inventory_batches
inventory_movements
audit_logs

-- Raw Silk
raw_silk_purchases
supplier_invoices
supplier_payments

-- Colouring
colouring_batches
colouring_expenses

-- Coloured Silk
coloured_silk_stock

-- Warping
doli_batches

-- Winding
winding_batches
bobbin_inventory

-- Transaction Tracking
purchase_transactions
```

---

## Phase 1: API Endpoints (Core)

### Authentication
```
POST   /api/auth/register
POST   /api/auth/login
POST   /api/auth/logout
POST   /api/auth/refresh-token
GET    /api/auth/me
```

### Suppliers
```
GET    /api/suppliers
POST   /api/suppliers
GET    /api/suppliers/:id
PUT    /api/suppliers/:id
DELETE /api/suppliers/:id
GET    /api/suppliers/:id/transactions
```

### Raw Silk Purchases
```
GET    /api/raw-silk/purchases
POST   /api/raw-silk/purchases
GET    /api/raw-silk/purchases/:id
PUT    /api/raw-silk/purchases/:id/status
POST   /api/raw-silk/purchases/:id/receive
GET    /api/raw-silk/stock
```

### Colouring
```
GET    /api/colouring/batches
POST   /api/colouring/batches
GET    /api/colouring/batches/:id
PUT    /api/colouring/batches/:id/status
POST   /api/colouring/batches/:id/send-to-factory
POST   /api/colouring/batches/:id/receive-from-factory
GET    /api/colouring/status
```

### Coloured Silk
```
GET    /api/coloured-silk/stock
GET    /api/coloured-silk/by-colour/:colour-id
GET    /api/coloured-silk/by-location/:location-id
```

### Doli
```
GET    /api/doli/list
POST   /api/doli/create
GET    /api/doli/:id
GET    /api/doli/available
```

### Bobbins
```
GET    /api/bobbins/stock
GET    /api/bobbins/by-colour/:colour-id
POST   /api/winding/batches
GET    /api/winding/batches/:id
```

### Inventory
```
GET    /api/inventory/movements
GET    /api/inventory/movements?batch_id=X&date_from=X&date_to=X
GET    /api/inventory/stock-summary
GET    /api/inventory/by-location/:location-id
POST   /api/inventory/adjustment (Manual adjustment with audit)
```

### Dashboard
```
GET    /api/dashboard/inventory-overview
GET    /api/dashboard/material-journey
GET    /api/dashboard/supplier-metrics
GET    /api/dashboard/processing-status
```

### Reports
```
GET    /api/reports/stock-report
GET    /api/reports/material-location
GET    /api/reports/inventory-movements
GET    /api/reports/supplier-report
GET    /api/reports/colouring-status
```

---

## Phase 1: Frontend Components (React/Next.js)

```
Components/
├── Auth/
│   ├── Login.jsx
│   ├── Register.jsx
│   └── AuthGuard.jsx
│
├── Layout/
│   ├── Sidebar.jsx
│   ├── Header.jsx
│   └── Dashboard.jsx
│
├── Suppliers/
│   ├── SupplierList.jsx
│   ├── SupplierForm.jsx
│   ├── SupplierDetails.jsx
│   └── SupplierTransactions.jsx
│
├── Weavers/
│   ├── WeaverList.jsx
│   ├── WeaverForm.jsx
│   └── WeaverDetails.jsx
│
├── Buyers/
│   ├── BuyerList.jsx
│   ├── BuyerForm.jsx
│   └── BuyerDetails.jsx
│
├── RawSilk/
│   ├── PurchaseList.jsx
│   ├── PurchaseForm.jsx
│   ├── ReceiveMaterial.jsx
│   ├── StockView.jsx
│   └── PurchaseInvoice.jsx
│
├── Colouring/
│   ├── ColouringBatchList.jsx
│   ├── ColouringBatchForm.jsx
│   ├── SendToFactory.jsx
│   ├── ReceiveFromFactory.jsx
│   ├── ColouringStatus.jsx
│   └── ColouringExpense.jsx
│
├── Coloured Silk/
│   ├── Stock.jsx
│   └── ByColour.jsx
│
├── Doli/
│   ├── DoliBatchList.jsx
│   ├── DoliBatchForm.jsx
│   └── DoliDetails.jsx
│
├── Bobbins/
│   ├── BobbiStock.jsx
│   ├── BobbiByColour.jsx
│   ├── WindingForm.jsx
│   └── WindingBatchList.jsx
│
├── Inventory/
│   ├── MovementHistory.jsx
│   ├── StockSummary.jsx
│   ├── ByLocation.jsx
│   └── AdjustmentForm.jsx
│
├── Dashboard/
│   ├── InventoryOverview.jsx
│   ├── MaterialJourney.jsx
│   ├── SupplierMetrics.jsx
│   └── ProcessingStatus.jsx
│
├── Reports/
│   ├── StockReport.jsx
│   ├── MaterialLocation.jsx
│   ├── InventoryMovements.jsx
│   ├── SupplierReport.jsx
│   └── ColouringStatus.jsx
│
└── Common/
    ├── Table.jsx
    ├── Form.jsx
    ├── Modal.jsx
    ├── DatePicker.jsx
    ├── Search.jsx
    └── Notifications.jsx
```

---

## Phase 1: Deliverables Checklist

- [ ] Backend API (Go) - All Phase 1 endpoints
- [ ] PostgreSQL database with all Phase 1 tables
- [ ] Frontend (React/Next.js) - All Phase 1 screens
- [ ] User authentication & role management
- [ ] Audit trail system
- [ ] Inventory movement ledger (immutable)
- [ ] Raw silk purchase workflow
- [ ] Colouring process workflow
- [ ] Doli and Bobbins management
- [ ] Dashboard with key metrics
- [ ] Phase 1 reports
- [ ] Unit tests (minimum 80% coverage)
- [ ] API documentation
- [ ] User manual (Phase 1)

---

## Phase 1: Success Criteria

The application can answer these questions perfectly:

1. ✅ **What raw materials do I currently have in stock?** (by type, colour, location)
2. ✅ **How much material is currently under processing?** (at colour factory, warping, winding)
3. ✅ **What is the current value of my inventory?**
4. ✅ **Which supplier provided which material and how much do I owe them?**
5. ✅ **Where is every piece of material currently located?**

---

---

# PHASE 2: Production, Finance & Reports
**Duration**: 6-8 weeks | **Focus**: Complete production workflow and financial tracking

---

## Phase 2 Objectives

1. ✅ Complete weaver management and work tracking
2. ✅ Material issue and return workflow
3. ✅ Production tracking and saree management
4. ✅ Complete sales and billing system
5. ✅ Income and expense tracking
6. ✅ Supplier and weaver payment management
7. ✅ Comprehensive reporting and analytics
8. ✅ Profit/Loss calculation and business intelligence

---

## Phase 2: Core Components

### 2.1 Weaver Management (Detailed)

#### A. Weaver Master (Enhanced from Phase 1)
```
Table: weavers [Already in Phase 1]
├── weaver_id
├── name
├── phone
├── address
├── village
├── joining_date
├── bank_details
├── opening_balance
├── status
└── notes

New Fields for Phase 2:
├── total_work_assigned
├── total_sarees_completed
├── total_payment_earned
├── total_advance_paid
├── total_payment_done
├── outstanding_balance
└── last_work_date
```

#### B. Weaver Transactions (New)
```
Table: weaver_transactions
├── transaction_id (PK)
├── weaver_id (FK)
├── transaction_type (Work Assigned, Advance Paid, Work Payment, Adjustment, etc.)
├── description
├── amount
├── date
├── reference_id (Material Issue ID, Production ID, etc.)
├── created_by
├── created_date
└── notes

This creates an immutable ledger for each weaver.
```

#### C. Weaver Ledger (View)
```
Query: Generate ledger for any weaver

Example Output:
WEAVER: Ravi

Opening Balance              ₹2,000
14-Aug  Work Assigned        ₹5,000  Balance: ₹7,000
14-Aug  Advance Paid        -₹2,000  Balance: ₹5,000
17-Aug  Work Payment        -₹4,000  Balance: ₹1,000
Material Adjustment         -₹500   Balance: ₹500
─────────────────────────────────────
Closing Balance              ₹500
```

#### UI Screens (Phase 2)
```
1. Weaver Master List (enhanced)
   - Total work assigned
   - Total production
   - Outstanding balance
   - Status

2. Weaver Details (new detailed view)
   - Personal info
   - Bank details
   - Opening balance

3. Weaver Ledger (new)
   - Complete transaction history
   - Running balance
   - Export to PDF/Excel

4. Weaver Statistics
   - Sarees produced by design
   - Payment history
   - Productivity metrics
```

---

### 2.2 Production Management

#### A. Work Order / Production Order (New)
```
Table: production_orders
├── production_id (PK)
├── po_reference (PO-00042)
├── weaver_id (FK)
├── design_id (FK)
├── colour_id (FK)
├── doli_required
├── bobbins_required
├── expected_sarees
├── rate_per_saree
├── total_expected_payment
├── issued_date
├── expected_completion_date
├── status (Pending, Issued, In-Progress, Completed, Cancelled)
├── notes
└── created_date

Automatic Actions:
- Create Production Order
- Auto-reserve required Doli and Bobbins
- Material Issue creates separate record
```

#### B. Material Issue to Weaver (New)
```
Table: material_issues
├── issue_id (PK)
├── production_order_id (FK)
├── weaver_id (FK)
├── issued_date
├── doli_issued
├── bobbins_issued
├── other_material
├── status (Pending Return, Returned)
├── return_date
├── created_date
└── notes

Automatic Actions on "Issue Material":
- Create inventory movement: -Doli from Warehouse to Weaver location
- Create inventory movement: -Bobbins from Warehouse to Weaver location
- Mark Doli as "Issued"
- Mark Bobbins as "Issued"
- Calculate: Warehouse Stock - Weaver Stock = Available in Warehouse

This answers: WHERE IS MY MATERIAL?
Warehouse: 10 Doli
Weaver Ravi: 2 Doli
Weaver Kumar: 1 Doli
────────────────
Total: 13 Doli
```

#### C. Material Return from Weaver (New)
```
Table: material_returns
├── return_id (PK)
├── material_issue_id (FK)
├── weaver_id (FK)
├── doli_returned
├── bobbins_returned
├── doli_damaged
├── bobbins_damaged
├── return_date
├── created_date
└── notes

Automatic Actions on "Receive Return":
- Create inventory movement: +Doli to Warehouse from Weaver location
- Create inventory movement: -Doli to Damaged location (if any)
- Update material_issue status to "Returned"
```

#### D. Production Output (New)
```
Table: production_outputs
├── output_id (PK)
├── production_order_id (FK)
├── weaver_id (FK)
├── design_id (FK)
├── colour_id (FK)
├── sarees_produced
├── sarees_approved
├── sarees_rejected
├── quality_notes
├── completion_date
├── payment_amount (sarees_approved × rate)
├── created_date
└── notes

Automatic Actions on "Confirm Production":
- Create inventory movement: +Sarees to Saree Stock
- Create weaver_transaction: Work Payment +₹XXXX
- Update production_order status to "Completed"
```

#### UI Screens (Phase 2)
```
1. Production Order List
   - Status-wise view (Pending, In-Progress, Completed)
   - Weaver filter
   - Expected vs Actual completion

2. Production Order Form
   - Select weaver
   - Select design
   - Select colour
   - Input material requirements
   - Set rate and expected sarees

3. Issue Material
   - Auto-select materials from Production Order
   - Confirm issue
   - Generate material issue receipt

4. Receive Material Return
   - Record returned quantity
   - Record damaged quantity
   - Update status

5. Production Completion
   - Record sarees produced
   - Record rejections
   - Approve quality
   - Auto-create payment transaction

6. Weaver Assignment Dashboard
   - Current work assigned to each weaver
   - Status of each assignment
   - Overdue productions
```

---

### 2.3 Saree Stock Management

#### A. Saree Inventory (New)
```
Table: sarees
├── saree_id (PK)
├── saree_reference (SR-2026-00001)
├── design_id (FK)
├── colour_id (FK)
├── saree_type (Regular, Special, etc.)
├── weaver_id (FK) [Who produced it]
├── production_order_id (FK)
├── production_date
├── cost_per_saree
├── selling_price
├── current_location_id (FK)
├── status (In Stock, Reserved, Sold, Damaged, Returned)
├── created_date
└── notes

Stock Calculation:
Sarees in Stock = COUNT WHERE status = "In Stock"
Sarees Reserved = COUNT WHERE status = "Reserved"
Sarees Sold = COUNT WHERE status = "Sold"
```

#### B. Saree Movement Ledger (Similar to Inventory)
```
Table: saree_movements
├── movement_id (PK)
├── saree_id (FK)
├── movement_type (Produced, Sold, Damaged, Adjusted)
├── from_location_id (FK)
├── to_location_id (FK)
├── reference_id (Sale Invoice, Damage Report, etc.)
├── created_date
└── notes
```

#### UI Screens (Phase 2)
```
1. Saree Stock List
   - All sarees with filter (design, colour, status)
   - Cost and selling price
   - Current location

2. Saree Details
   - Complete history
   - Weaver info
   - Production date
   - Buyer info (if sold)

3. Stock Summary by Design/Colour
   - Quantity available
   - Total value
   - Slow-moving items
```

---

### 2.4 Sales & Billing

#### A. Sales Order (New)
```
Table: sales_orders
├── order_id (PK)
├── buyer_id (FK)
├── order_date
├── expected_delivery_date
├── status (Pending, Fulfilled, Cancelled)
├── notes
└── created_date

This is optional - can directly create invoices
```

#### B. Sales Invoice (New - CORE)
```
Table: invoices
├── invoice_id (PK)
├── invoice_number (INV-2026-001)
├── buyer_id (FK)
├── invoice_date
├── sale_date
├── subtotal
├── discount_amount
├── tax_amount (SGST/CGST if applicable)
├── total_amount
├── payment_status (Unpaid, Partial, Paid)
├── amount_received
├── amount_pending
├── notes
├── created_date
└── created_by

Automatic Actions:
- Generate unique invoice number
- Calculate totals
- Create buyer_transaction: +amount_pending
```

#### C. Invoice Line Items (New)
```
Table: invoice_items
├── line_item_id (PK)
├── invoice_id (FK)
├── saree_id (FK)
├── quantity
├── rate
├── amount
└── notes

Automatic Action on "Confirm Sale":
- Update saree status: "In Stock" → "Sold"
- Create saree_movement: to Buyer location
- Create buyer_transaction: Sale entry
```

#### D. Sale Payments (New)
```
Table: sale_payments
├── payment_id (PK)
├── invoice_id (FK)
├── buyer_id (FK)
├── payment_amount
├── payment_date
├── payment_mode (Cash, Check, Bank Transfer, etc.)
├── reference_number (Cheque no, Transaction ID, etc.)
├── created_date
└── notes

Automatic Action:
- Update invoice: amount_received, amount_pending
- Create buyer_transaction: Payment entry
```

#### UI Screens (Phase 2)
```
1. Create Invoice
   - Select buyer
   - Add sarees to invoice
   - Calculate totals
   - Auto-generate invoice

2. Invoice List
   - All invoices
   - Status filter (Unpaid, Partial, Paid)
   - Outstanding amount

3. Invoice Details
   - Items
   - Payment history
   - Print invoice

4. Record Payment
   - Select invoice
   - Record payment amount
   - Payment mode
   - Reference number

5. Buyer Invoices
   - All invoices for a buyer
   - Total outstanding
```

---

### 2.5 Buyer Ledger (New)

#### Buyer Transactions
```
Table: buyer_transactions
├── transaction_id (PK)
├── buyer_id (FK)
├── transaction_type (Invoice, Payment, Adjustment, Credit Note, etc.)
├── description
├── amount
├── date
├── reference_id (Invoice ID, Payment ID, etc.)
├── created_date
└── notes
```

#### Buyer Ledger View
```
Query: Generate ledger for any buyer

Example:
BUYER: XYZ Textiles

Opening Balance         ₹0
14-Aug  Invoice 001     ₹1,00,000  Outstanding: ₹1,00,000
15-Aug  Payment         -₹50,000   Outstanding: ₹50,000
18-Aug  Invoice 002     ₹80,000    Outstanding: ₹1,30,000
20-Aug  Payment         -₹30,000   Outstanding: ₹1,00,000
─────────────────────────────────────────────
Outstanding Amount      ₹1,00,000
```

#### UI Screens
```
1. Buyer Ledger
   - Transaction history
   - Running balance
   - Outstanding amount
   - Export

2. Payment Reminder
   - Overdue invoices
   - Days pending
   - Amount pending
```

---

### 2.6 Income & Expenses

#### A. Income (Already tracked via Sales)
```
Table: income_transactions
├── income_id (PK)
├── income_type (Saree Sales, Other Income, etc.)
├── description
├── amount
├── date
├── reference_id (Invoice ID)
├── created_date
└── notes

Automatically created from:
- Sales invoices
- Buyer payments (if explicitly recorded)
```

#### B. Expenses (New - CORE)
```
Table: expenses
├── expense_id (PK)
├── expense_category_id (FK)
├── vendor/supplier_id (FK) [Optional - for tracking]
├── description
├── amount
├── expense_date
├── payment_status (Pending, Paid)
├── related_batch_id (FK) [Links to material/production]
├── payment_mode (Cash, Check, Bank Transfer, etc.)
├── reference_number
├── created_date
└── notes

Expense Categories:
- Raw Material Purchase
- Colouring Charges
- Warping Charges
- Winding Charges
- Weaver Payment
- Transport
- Electricity
- Machine Maintenance
- Labour/Salary
- Rent
- Packaging
- Miscellaneous
```

#### C. Expense Categories (Master)
```
Table: expense_categories
├── category_id (PK)
├── name
├── code
└── description
```

#### Automatic Expense Creation
```
Expenses are auto-created for:
- Raw silk purchase (Raw Material Purchase)
- Colouring charges (Colouring Charges)
- Warping charges (Warping Charges)
- Winding charges (Winding Charges)
- Weaver payments (Weaver Payment)
```

#### UI Screens (Phase 2)
```
1. Expense Entry
   - Category selection
   - Amount and date
   - Link to batch/order
   - Payment tracking

2. Expense List
   - By category
   - By date range
   - Status filter (Paid/Pending)
   - Supplier filter

3. Expense Summary
   - Monthly breakdown
   - Category-wise
   - vs. Budget (if available)
```

---

### 2.7 Supplier Payments

#### Tables (Already have from Phase 1)
```
Table: supplier_transactions [Enhanced]
├── transaction_id (PK)
├── supplier_id (FK)
├── transaction_type (Purchase, Payment, Adjustment, etc.)
├── description
├── amount
├── date
├── reference_id (Purchase ID, Colouring ID, etc.)
├── created_date
└── notes

Auto-created for:
- Raw silk purchases
- Colouring batches
- Other vendor services
```

#### Supplier Ledger View
```
Supplier: ABC Silk Supplier

Purchase 001     +₹2,00,000
Purchase 002     +₹1,50,000
Payment          -₹1,50,000
─────────────────────────────
Outstanding      ₹2,00,000
```

#### UI Screens (Phase 2)
```
1. Supplier Ledger
   - Transaction history
   - Outstanding amount
   - Payment due dates

2. Payment Tracking
   - Overdue payments
   - Upcoming due dates
   - Amount to pay

3. Payment Record
   - Record supplier payment
   - Cheque details
   - Bank details
```

---

### 2.8 Weaver Payments

#### Payment Processing
```
Workflow:
Production Completed
        ↓
Calculate Payment = Sarees Approved × Rate per Saree
        ↓
Create Weaver Payment Record
        ↓
Process Payment (Cash/Check/Bank)
        ↓
Update Weaver Ledger
```

#### Tables
```
Table: weaver_payments
├── payment_id (PK)
├── weaver_id (FK)
├── production_order_id (FK) [What work is being paid for]
├── payment_amount
├── payment_date
├── payment_mode (Cash, Check, Bank Transfer)
├── reference_number
├── advance_deducted (if any)
├── net_amount_paid
├── created_date
└── notes

Automatic Action:
- Create weaver_transaction: Payment -₹XXXXX
- Update weaver: total_payment_done
```

#### UI Screens (Phase 2)
```
1. Generate Payment
   - Select weaver
   - Show pending work (auto-calculated)
   - Confirm payment amount

2. Record Payment
   - Payment mode
   - Reference number
   - Advance adjustments

3. Weaver Payment History
   - All payments to weaver
   - Date-wise
   - Amount-wise
```

---

### 2.9 Financial Reports (Phase 2)

#### A. Profit & Loss Statement
```
Report: Profit & Loss (Customizable Date Range)

                         Month      Year-to-Date
INCOME
Saree Sales           ₹XX,XXX      ₹XX,XXX,XXX
Other Income          ₹X,XXX       ₹XX,XXX
────────────────────────────────────────────
Total Income          ₹XX,XXX      ₹XX,XXX,XXX

EXPENSES
Raw Material          ₹XX,XXX      ₹XX,XXX,XXX
Colouring             ₹X,XXX       ₹XX,XXX
Warping               ₹X,XXX       ₹XX,XXX
Winding               ₹X,XXX       ₹XX,XXX
Weaver Payments       ₹XX,XXX      ₹XX,XXX,XXX
Transport             ₹X,XXX       ₹XX,XXX
Utilities             ₹X,XXX       ₹XX,XXX
Maintenance           ₹X,XXX       ₹XX,XXX
Other                 ₹X,XXX       ₹XX,XXX
────────────────────────────────────────────
Total Expenses        ₹XX,XXX      ₹XX,XXX,XXX

PROFIT/LOSS           ₹XX,XXX      ₹XX,XXX,XXX
```

#### B. Cash Flow Report
```
Cash Inflows          ₹XX,XXX
Cash Outflows         ₹XX,XXX
─────────────────────────────
Net Cash Flow         ₹XX,XXX

Receivables Outstanding    ₹XX,XXX
Payables Outstanding       ₹XX,XXX
```

#### C. Cost per Saree Analysis
```
Report: Cost Breakdown per Saree

                        Value        % of Cost
Raw Material Cost      ₹2,500         58%
Colouring              ₹300           7%
Warping                ₹150           3%
Winding                ₹100           2%
Weaver Payment         ₹1,000         23%
Transport              ₹150           3%
Other Overhead         ₹100           2%
───────────────────────────────────
Total Cost             ₹4,300         100%

Selling Price          ₹6,000
Gross Profit           ₹1,700         28%

Profit Margin          28%
```

#### D. Inventory Valuation Report
```
Report: Current Inventory Value

Material Type          Quantity    Unit Cost    Total Value
────────────────────────────────────────────────────────────
Raw Silk               182 kg      ₹X,XXX       ₹XX,XX,XXX
Coloured Silk          95 kg       ₹Y,YYY       ₹XX,XX,XXX
Doli                   42 units    ₹Z,ZZZ       ₹X,XX,XXX
Bobbins                850 units   ₹50          ₹42,500
Sarees in Stock        127 units   ₹4,300       ₹5,46,100
────────────────────────────────────────────────────────────
Total Inventory Value                           ₹XX,XX,XXX
```

---

### 2.10 Production Reports (Phase 2)

#### A. Weaver-wise Production Report
```
Report: Weaver Production Summary

Weaver          Sarees       Approved      Rejected      Rate        Total Payment
                Produced     Sarees        Sarees        (per Saree) Earned
─────────────────────────────────────────────────────────────────────────────────
Ravi            25           24            1             ₹1,000      ₹24,000
Kumar           18           18            0             ₹1,000      ₹18,000
Lakshmi         22           20            2             ₹1,000      ₹20,000
─────────────────────────────────────────────────────────────────────────────────
Total           65           62            3                         ₹62,000
```

#### B. Design-wise Production Report
```
Report: Design Production Summary

Design          Quantity Produced    Colour      Selling Price    Total Value
─────────────────────────────────────────────────────────────────────────────
Design A        45                   Red         ₹6,000           ₹2,70,000
Design B        12                   Blue        ₹5,500           ₹66,000
Design C        8                    Green       ₹7,000           ₹56,000
─────────────────────────────────────────────────────────────────────────────
Total           65                                                 ₹3,92,000
```

#### C. Production vs. Target (Future Enhancement)
```
If targets are set, compare actual vs. target
```

---

### 2.11 Outstanding Reports (Phase 2)

#### A. Buyer Outstanding Report
```
Report: Buyer Outstanding Amount

Buyer           Total Sales    Amount Paid    Outstanding    Days Pending
────────────────────────────────────────────────────────────────────────
XYZ Textiles    ₹3,50,000      ₹2,00,000      ₹1,50,000      45 days
ABC Enterprises ₹1,20,000      ₹1,20,000      ₹0             -
PQR Fabrics     ₹2,00,000      ₹50,000        ₹1,50,000      60 days
────────────────────────────────────────────────────────────────────────
Total           ₹6,70,000      ₹3,70,000      ₹3,00,000
```

#### B. Supplier Outstanding Report
```
Report: Supplier Outstanding Amount

Supplier           Total Purchases    Paid        Outstanding    Days Pending
──────────────────────────────────────────────────────────────────────────
Silk Vendor A      ₹4,50,000          ₹3,00,000   ₹1,50,000      30 days
Colour Factory B   ₹1,80,000          ₹1,80,000   ₹0             -
Warping Unit C     ₹90,000            ₹45,000     ₹45,000        15 days
──────────────────────────────────────────────────────────────────────────
Total              ₹7,20,000          ₹5,25,000   ₹1,95,000
```

#### C. Weaver Outstanding Report
```
Report: Weaver Outstanding Amount

Weaver          Total Earned    Paid        Outstanding    Notes
──────────────────────────────────────────────────────────
Ravi            ₹2,40,000       ₹2,00,000   ₹40,000        Advance: ₹5,000
Kumar           ₹1,80,000       ₹1,80,000   ₹0             
Lakshmi         ₹2,00,000       ₹1,50,000   ₹50,000        
──────────────────────────────────────────────────────────
Total           ₹6,20,000       ₹5,30,000   ₹90,000
```

---

### 2.12 Dashboard (Phase 2 - Enhanced)

#### Enhanced KPIs
```
┌─────────────────────────────────────────────────────────┐
│   WEAVING BUSINESS DASHBOARD - PHASE 2 (COMPLETE)       │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  INVENTORY (From Phase 1)                               │
│  └─ Total Value: ₹XX,XX,XXX                             │
│                                                         │
│  PRODUCTION                                             │
│  ├─ Sarees in Stock: 127 units (Value: ₹5,46,100)     │
│  ├─ Sarees Sold (Month): 34 units (₹2,04,000)         │
│  ├─ Sarees in Production: 12 units                     │
│  ├─ Active Weavers: 18                                 │
│  └─ Production Rate: 65/month                          │
│                                                         │
│  FINANCE                                                │
│  ├─ Total Income (Month): ₹XX,XX,XXX                   │
│  ├─ Total Expenses (Month): ₹XX,XX,XXX                 │
│  ├─ Profit (Month): ₹XX,XX,XXX                         │
│  ├─ Profit Margin: 28%                                 │
│  ├─ Cost per Saree: ₹4,300                             │
│  └─ Selling Price (Avg): ₹6,000                        │
│                                                         │
│  RECEIVABLES & PAYABLES                                │
│  ├─ Buyer Outstanding: ₹3,00,000 (5 days avg pending) │
│  ├─ Weaver Outstanding: ₹90,000                        │
│  ├─ Supplier Outstanding: ₹1,95,000                    │
│  └─ Net Position: -₹1,95,000 (We owe)                  │
│                                                         │
│  CASH POSITION                                          │
│  ├─ Current Month Cash In: ₹XX,XX,XXX                  │
│  ├─ Current Month Cash Out: ₹XX,XX,XXX                 │
│  └─ Net Cash Flow: ₹XX,XX,XXX                          │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

#### Dashboard Sections
```
1. Financial Summary
   - YTD Income
   - YTD Expenses
   - YTD Profit/Loss
   - Current month comparison

2. Production Metrics
   - Sarees produced (month)
   - Sarees sold (month)
   - Current stock
   - Production by design

3. Weaver Dashboard
   - Active weavers
   - Pending work (by weaver)
   - Pending payments

4. Buyer Dashboard
   - Recent sales
   - Outstanding amounts
   - Best customers

5. Financial Health
   - Gross margin
   - Outstanding receivables
   - Outstanding payables
   - Cash position

6. Alerts & Notifications
   - Overdue weaver work
   - Overdue buyer payments
   - Overdue supplier payments
   - Low inventory items
```

---

## Phase 2: Database Schema (Complete)

```sql
-- Phase 1 Tables +

-- Production
production_orders
material_issues
material_returns
production_outputs

-- Saree Management
sarees
saree_movements

-- Sales & Billing
invoices
invoice_items
sale_payments
buyers [Enhanced from Phase 1]

-- Ledgers & Transactions
weaver_transactions
buyer_transactions
income_transactions
expenses
expense_categories

-- Payments
weaver_payments
supplier_payments [Enhanced]
```

---

## Phase 2: API Endpoints (Additional)

### Production Management
```
GET    /api/production-orders
POST   /api/production-orders
GET    /api/production-orders/:id
PUT    /api/production-orders/:id/status

GET    /api/material-issues
POST   /api/material-issues
GET    /api/material-issues/:id

GET    /api/material-returns
POST   /api/material-returns/:issue-id

GET    /api/production-outputs
POST   /api/production-outputs
GET    /api/production-outputs/:id
```

### Saree Management
```
GET    /api/sarees
POST   /api/sarees
GET    /api/sarees/:id
PUT    /api/sarees/:id/status
GET    /api/sarees/stock-summary
GET    /api/sarees/by-design/:design-id
GET    /api/sarees/by-colour/:colour-id
```

### Sales & Billing
```
GET    /api/invoices
POST   /api/invoices
GET    /api/invoices/:id
GET    /api/invoices/:id/print

POST   /api/invoices/:id/items
GET    /api/invoices/:id/items
DELETE /api/invoices/:id/items/:item-id

GET    /api/sale-payments
POST   /api/sale-payments
GET    /api/sale-payments/:id
```

### Weaver Management
```
GET    /api/weavers/:id/ledger
GET    /api/weavers/:id/transactions
GET    /api/weavers/:id/payments

GET    /api/weaver-payments
POST   /api/weaver-payments
GET    /api/weaver-payments/:id
```

### Buyer Management
```
GET    /api/buyers/:id/ledger
GET    /api/buyers/:id/invoices
GET    /api/buyers/:id/payments
GET    /api/buyers/:id/outstanding
```

### Supplier Management (Enhanced)
```
GET    /api/suppliers/:id/ledger
GET    /api/suppliers/:id/transactions
GET    /api/suppliers/:id/payments
GET    /api/suppliers/:id/outstanding

GET    /api/supplier-payments
POST   /api/supplier-payments
GET    /api/supplier-payments/:id
```

### Finance & Accounting
```
GET    /api/expenses
POST   /api/expenses
GET    /api/expenses/:id
PUT    /api/expenses/:id/payment-status

GET    /api/expense-categories
POST   /api/expense-categories

GET    /api/income
GET    /api/income/:id
```

### Reports (Phase 2)
```
GET    /api/reports/profit-loss?from=X&to=X
GET    /api/reports/cash-flow?from=X&to=X
GET    /api/reports/cost-per-saree
GET    /api/reports/inventory-valuation
GET    /api/reports/weaver-production?from=X&to=X
GET    /api/reports/design-production?from=X&to=X
GET    /api/reports/buyer-outstanding
GET    /api/reports/supplier-outstanding
GET    /api/reports/weaver-outstanding
```

---

## Phase 2: Frontend Components (React/Next.js - Additional)

```
Components/

├── Production/
│   ├── ProductionOrderList.jsx
│   ├── ProductionOrderForm.jsx
│   ├── IssueToWeaver.jsx
│   ├── ReceiveFromWeaver.jsx
│   ├── RecordProduction.jsx
│   └── ProductionDashboard.jsx
│
├── Sarees/
│   ├── SareeList.jsx
│   ├── SareeForm.jsx
│   ├── SareeDetails.jsx
│   ├── SareeByDesign.jsx
│   └── SareeByColour.jsx
│
├── Sales/
│   ├── InvoiceList.jsx
│   ├── CreateInvoice.jsx
│   ├── InvoiceDetails.jsx
│   ├── InvoicePrint.jsx
│   ├── RecordPayment.jsx
│   └── PaymentHistory.jsx
│
├── Ledgers/
│   ├── WeaverLedger.jsx
│   ├── BuyerLedger.jsx
│   ├── SupplierLedger.jsx
│   └── LedgerComparison.jsx
│
├── Finance/
│   ├── ExpenseEntry.jsx
│   ├── ExpenseList.jsx
│   ├── IncomeList.jsx
│   ├── PaymentRecords.jsx
│   └── CashManagement.jsx
│
├── Reports/
│   ├── ProfitLossReport.jsx
│   ├── CashFlowReport.jsx
│   ├── CostPerSaree.jsx
│   ├── InventoryValuation.jsx
│   ├── WeaverProduction.jsx
│   ├── DesignProduction.jsx
│   ├── BuyerOutstanding.jsx
│   ├── SupplierOutstanding.jsx
│   ├── WeaverOutstanding.jsx
│   └── ReportGenerator.jsx
│
└── Dashboard/
    └── FinanceDashboard.jsx [Enhanced]
```

---

## Phase 2: Deliverables Checklist

- [ ] Production order management
- [ ] Material issue and return workflow
- [ ] Production completion and tracking
- [ ] Saree inventory management
- [ ] Sales and billing system
- [ ] Invoice generation and management
- [ ] Payment tracking (Sales, Supplier, Weaver)
- [ ] Complete ledgers (Buyer, Weaver, Supplier)
- [ ] Income and expense tracking
- [ ] Profit & loss calculation
- [ ] Cost per saree analysis
- [ ] Inventory valuation
- [ ] Production reports
- [ ] Outstanding reports
- [ ] Financial reports
- [ ] Enhanced dashboard
- [ ] Unit tests (minimum 80% coverage)
- [ ] Integration tests
- [ ] API documentation (updated)
- [ ] User manual (complete)
- [ ] Data migration utilities (if needed)

---

## Phase 2: Success Criteria

The application can answer these questions perfectly:

1. ✅ **What has each weaver produced and how much are they owed?**
2. ✅ **How many sarees do I have in stock and what are they worth?**
3. ✅ **Who bought what sarees and how much do they owe me?**
4. ✅ **What is my total income and expenses?**
5. ✅ **What is my current profit/loss?**
6. ✅ **What is my cost per saree and am I profitable?**
7. ✅ **How much money is receivable from buyers and payable to suppliers/weavers?**
8. ✅ **What is the current value of all my inventory?**
9. ✅ **Which weavers are most productive?**
10. ✅ **Which designs/colours are best sellers?**

---

---

# PHASE COMPARISON & TIMELINE

## Timeline Overview

```
PHASE 1: Weeks 1-8 (Foundation & Inventory)
├─ Week 1-2: Backend setup, Database design
├─ Week 3-4: Core API endpoints
├─ Week 5-6: Frontend development
├─ Week 7-8: Testing, bug fixes, deployment
└─ Result: Fully working inventory system

PHASE 2: Weeks 9-16 (Production, Finance & Reports)
├─ Week 9-10: Production management
├─ Week 11-12: Sales and billing
├─ Week 13-14: Finance and reporting
├─ Week 15-16: Testing, optimization, deployment
└─ Result: Complete business management system
```

---

## Feature Comparison

| Feature | Phase 1 | Phase 2 |
|---------|---------|---------|
| **Masters** | Suppliers, Weavers, Buyers, Products | Enhanced |
| **Inventory** | Raw Silk, Coloured Silk, Doli, Bobbins | ✅ Complete |
| **Production** | - | ✅ Work Orders, Material Issue/Return |
| **Sarees** | - | ✅ Stock & Management |
| **Sales** | - | ✅ Invoicing & Payment |
| **Finance** | Basic | ✅ Income, Expenses, P&L |
| **Reports** | Inventory | ✅ All Reports |
| **Dashboard** | Inventory KPIs | ✅ Complete Business View |

---

## Data Flow by Phase

### Phase 1: Inventory Flow
```
SUPPLIER → Raw Silk Purchase → Warehouse Inventory
                                      ↓
                          Colour Factory Process
                                      ↓
                         Coloured Silk Inventory
                                      ↓
                            Warping → Doli Inventory
                                      ↓
                            Winding → Bobbin Inventory
```

### Phase 2: Complete Business Flow
```
SUPPLIER → Purchase → Inventory
                         ↓
                    Colour Factory
                         ↓
                    Coloured Silk
                         ↓
                    WEAVER (Material Issue)
                         ↓
                    Production (Sarees)
                         ↓
                    Saree Stock
                         ↓
                    BUYER (Invoice)
                         ↓
                    Payment
                         ↓
                    ACCOUNTS
                         ↓
                    Profit/Loss
```

---

## Testing Strategy

### Phase 1
- Unit tests for inventory movement logic
- API endpoint tests
- UI component tests
- Integration tests for purchase workflow

### Phase 2
- Unit tests for production logic
- Unit tests for financial calculations
- API endpoint tests
- End-to-end workflow tests
- Report accuracy tests
- Performance tests (large datasets)

---

## Risk Mitigation

### Phase 1 Risks
- Inventory ledger immutability (SOLUTION: Database constraints + audit logs)
- Material location tracking complexity (SOLUTION: Centralized location system)
- Performance with large inventory (SOLUTION: Proper indexing)

### Phase 2 Risks
- Complex P&L calculations (SOLUTION: Transaction-based approach, audit trail)
- Data integrity across modules (SOLUTION: Foreign keys, transactions)
- Report performance (SOLUTION: Pre-calculated summaries)

---

## Deployment Strategy

### Phase 1 Deployment
```
1. Setup PostgreSQL database
2. Deploy Go backend (Dockerized)
3. Deploy React frontend
4. Test with real data
5. Train staff on Phase 1 features
6. Go live with inventory system
```

### Phase 2 Deployment
```
1. Database migrations
2. Deploy new backend modules
3. Deploy new frontend screens
4. Historical data import (if needed)
5. Test complete workflows
6. Staff training on new features
7. Go live with production & finance
```

---

## Future Enhancements (Post-Phase 2)

1. **QR/Barcode Scanning** - For faster material tracking
2. **WhatsApp Integration** - Send invoices, payment reminders
3. **Mobile App** - Weaver updates, material tracking
4. **Advanced Analytics** - Trends, predictions
5. **Bulk Operations** - Batch processing
6. **Multi-location** - Support multiple warehouses
7. **API for Partners** - For external integrations
8. **Machine Learning** - Demand forecasting, price optimization

---

## Success Metrics

### Phase 1 Success
- ✅ System can track all material movements
- ✅ Material location is always known
- ✅ Inventory is 100% accurate
- ✅ Dashboard shows real-time stock levels
- ✅ Zero data loss or corruption

### Phase 2 Success
- ✅ Complete business workflow is digitized
- ✅ Profit/loss is calculated accurately
- ✅ All ledgers are reconciled
- ✅ Reporting is comprehensive
- ✅ Business owner has complete visibility
- ✅ Decision-making is faster and data-driven

---

## Important Notes

1. **DO NOT skip Phase 1** - A solid inventory foundation is critical
2. **Test thoroughly** - Money and materials are involved
3. **Keep audit logs** - For compliance and debugging
4. **Backup frequently** - This is business-critical data
5. **Train well** - Good tool usage depends on training
6. **Iterate with real data** - Use actual business scenarios

---

# Next Steps

1. ✅ Review this two-phase plan
2. ✅ Confirm technology stack (Go, React, PostgreSQL)
3. ✅ Prepare development environment
4. ✅ Create detailed database ERD
5. ✅ Start Phase 1 backend development

---

**Document Version**: 1.0  
**Created**: August 14, 2026  
**Purpose**: Development roadmap for Weaving Business Management System
