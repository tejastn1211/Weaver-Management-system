# Weaving Business Management System - Technical Architecture

---

## Overview

This document describes the technical architecture, API design, and technology stack for the Weaving Business Management System.

---

---

# ARCHITECTURE OVERVIEW

---

## High-Level Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                     WEB APPLICATION                          │
│                   React / Next.js (Frontend)                 │
│                                                              │
│  ┌─────────────┬──────────────┬──────────────┬────────────┐ │
│  │ Dashboard   │ Inventory    │ Production   │ Finance    │ │
│  │ Screens     │ Screens      │ Screens      │ Screens    │ │
│  └──────┬──────┴──────┬───────┴──────┬───────┴────┬───────┘ │
│         │             │              │            │          │
│         └─────────────┴──────────────┴────────────┘          │
│                       │                                       │
│                REST API Calls                                 │
│                       │                                       │
└───────────────────────┼───────────────────────────────────────┘
                        │
                        ▼
        ┌───────────────────────────────────┐
        │      Go Backend API (Gin/Fiber)   │
        │                                   │
        │  ┌─────────────────────────────┐  │
        │  │  Authentication & Auth      │  │
        │  │  JWT/Session Management     │  │
        │  └─────────────────────────────┘  │
        │                                   │
        │  ┌─────────────────────────────┐  │
        │  │  Business Logic Layer       │  │
        │  │                             │  │
        │  │  ├─ Inventory Service       │  │
        │  │  ├─ Production Service      │  │
        │  │  ├─ Finance Service         │  │
        │  │  ├─ Weaver Service          │  │
        │  │  └─ Sales Service           │  │
        │  └─────────────────────────────┘  │
        │                                   │
        │  ┌─────────────────────────────┐  │
        │  │  Data Access Layer          │  │
        │  │  (Database Queries)         │  │
        │  └─────────────────────────────┘  │
        │                                   │
        └───────────────────────────────────┘
                        │
                        ▼
        ┌───────────────────────────────────┐
        │     PostgreSQL Database           │
        │                                   │
        │  ├─ Inventory Batches            │
        │  ├─ Inventory Movements          │
        │  ├─ Production Orders            │
        │  ├─ Sales Invoices               │
        │  ├─ Ledgers & Transactions       │
        │  └─ Audit Logs                   │
        └───────────────────────────────────┘
```

---

## Component Breakdown

### Frontend (React/Next.js)
- **Role**: User interface
- **Responsibilities**: Display data, collect user input, validation
- **Communication**: REST API calls to backend

### Backend (Go)
- **Role**: Business logic and data processing
- **Responsibilities**: Authentication, authorization, business rules, data validation
- **Communication**: REST API with frontend, database queries

### Database (PostgreSQL)
- **Role**: Data persistence
- **Responsibilities**: Store all data, maintain relationships, audit trail

---

---

# TECHNOLOGY STACK

---

## Backend

### Language & Framework
- **Language**: Go (Golang)
- **HTTP Framework**: Gin or Fiber
- **Package Management**: Go Modules
- **Version**: Go 1.21+

### Database
- **Primary**: PostgreSQL 14+
- **Migration Tool**: golang-migrate or Flyway
- **Connection Pool**: pgx with built-in pool

### Authentication
- **Method**: JWT (JSON Web Tokens)
- **Library**: github.com/golang-jwt/jwt
- **Refresh Token**: Short-lived access tokens + long-lived refresh tokens

### Validation
- **Input Validation**: github.com/go-playground/validator
- **Sanitation**: github.com/microcosm-cc/bluemonday

### Error Handling
- **Structured Errors**: Custom error types with codes
- **Logging**: Zerolog or Zap

### Testing
- **Unit Tests**: stdlib testing package + assertions
- **Integration Tests**: testcontainers for PostgreSQL
- **Mocking**: github.com/golang/mock

### Deployment
- **Containerization**: Docker
- **Container Orchestration**: Docker Compose (initially)

---

## Frontend

### Framework
- **Framework**: Next.js 14+ (React 18+)
- **Package Manager**: npm or yarn
- **TypeScript**: Recommended

### UI Components
- **Component Library**: Material-UI (MUI) or Ant Design or custom
- **Styling**: Tailwind CSS
- **Icons**: React Icons

### State Management
- **Option 1**: Redux Toolkit + React Query
- **Option 2**: Zustand + React Query
- **Option 3**: Context API + React Query

### Data Fetching
- **Library**: React Query (TanStack Query)
- **Caching**: Automatic with React Query

### Forms
- **Library**: React Hook Form + Zod (or Yup)
- **Validation**: Client + server-side

### Utilities
- **Date/Time**: Day.js or date-fns
- **Number Formatting**: numeral.js
- **Charting**: Chart.js or Recharts (for reports)
- **PDF Export**: jsPDF + html2pdf
- **Excel Export**: exceljs

### Development Tools
- **Build**: Next.js built-in
- **Linting**: ESLint
- **Formatting**: Prettier
- **Type Checking**: TypeScript

---

## DevOps & Deployment

### Local Development
- **Docker**: Docker Desktop
- **Docker Compose**: Multi-container setup
- **Environment Management**: .env files

### Production Deployment
- **VPS/Cloud**: AWS EC2 / DigitalOcean / Linode
- **Container Runtime**: Docker
- **Reverse Proxy**: Nginx
- **SSL Certificate**: Let's Encrypt
- **Monitoring**: Prometheus + Grafana (Phase 3+)
- **Logging**: ELK Stack (Phase 3+)

---

---

# REST API DESIGN

---

## API Principles

1. **RESTful**: Follow REST conventions
2. **Versioning**: API v1 (/api/v1/...)
3. **JSON**: All requests/responses in JSON
4. **Status Codes**: Proper HTTP status codes
5. **Error Handling**: Consistent error responses
6. **Rate Limiting**: Implement rate limiting
7. **Pagination**: Consistent pagination for lists
8. **Filtering & Sorting**: Support filtering and sorting

---

## API Response Format

### Success Response
```json
{
  "success": true,
  "statusCode": 200,
  "message": "Operation successful",
  "data": {
    // Actual response data
  },
  "timestamp": "2026-08-14T10:30:00Z"
}
```

### Error Response
```json
{
  "success": false,
  "statusCode": 400,
  "error": "INVALID_INPUT",
  "message": "Validation failed",
  "details": [
    {
      "field": "quantity",
      "message": "Quantity must be greater than 0"
    }
  ],
  "timestamp": "2026-08-14T10:30:00Z"
}
```

---

## Error Codes (Partial List)

```
2xx - Success
  200 OK
  201 Created
  204 No Content

4xx - Client Error
  400 Bad Request (Invalid input)
  401 Unauthorized (Not authenticated)
  403 Forbidden (Authenticated but not authorized)
  404 Not Found
  409 Conflict (Data conflict, e.g., duplicate entry)
  422 Unprocessable Entity (Validation failed)

5xx - Server Error
  500 Internal Server Error
  503 Service Unavailable
```

---

## API Pagination

For list endpoints, support pagination:

```
GET /api/v1/raw-silk/purchases?page=1&limit=20&sortBy=purchase_date&sortOrder=DESC

Response:
{
  "success": true,
  "data": [...],
  "pagination": {
    "page": 1,
    "limit": 20,
    "total": 150,
    "totalPages": 8,
    "hasNext": true,
    "hasPrev": false
  }
}
```

---

## Phase 1 API Endpoints

### Authentication
```
POST   /api/v1/auth/register
POST   /api/v1/auth/login
POST   /api/v1/auth/logout
POST   /api/v1/auth/refresh-token
GET    /api/v1/auth/me
POST   /api/v1/auth/change-password
```

### Suppliers
```
GET    /api/v1/suppliers
POST   /api/v1/suppliers
GET    /api/v1/suppliers/:id
PUT    /api/v1/suppliers/:id
DELETE /api/v1/suppliers/:id
GET    /api/v1/suppliers/:id/transactions
GET    /api/v1/suppliers/outstanding
```

### Weavers
```
GET    /api/v1/weavers
POST   /api/v1/weavers
GET    /api/v1/weavers/:id
PUT    /api/v1/weavers/:id
DELETE /api/v1/weavers/:id
```

### Buyers
```
GET    /api/v1/buyers
POST   /api/v1/buyers
GET    /api/v1/buyers/:id
PUT    /api/v1/buyers/:id
DELETE /api/v1/buyers/:id
```

### Products, Colours, Locations, Designs
```
GET    /api/v1/products
GET    /api/v1/colours
GET    /api/v1/locations
GET    /api/v1/designs
```

### Raw Silk Purchases
```
GET    /api/v1/raw-silk/purchases?page=1&limit=20
POST   /api/v1/raw-silk/purchases
GET    /api/v1/raw-silk/purchases/:id
PUT    /api/v1/raw-silk/purchases/:id
PUT    /api/v1/raw-silk/purchases/:id/status

-- Mark as received (creates inventory batch + movements)
POST   /api/v1/raw-silk/purchases/:id/receive
  Body: {
    "actual_quantity_received": 48.5,
    "batch_reference": "KS-2026-001",
    "notes": "2 kg wastage"
  }

GET    /api/v1/raw-silk/stock
GET    /api/v1/raw-silk/stock/summary
GET    /api/v1/raw-silk/purchase-summary?supplier_id=1&from_date=2026-08-01&to_date=2026-08-31
```

### Colouring
```
GET    /api/v1/colouring/batches?status=pending
POST   /api/v1/colouring/batches
GET    /api/v1/colouring/batches/:id
PUT    /api/v1/colouring/batches/:id/status

-- Send material to colour factory
POST   /api/v1/colouring/batches/:id/send-to-factory
  Body: {
    "date_sent": "2026-08-14",
    "expected_return_date": "2026-08-20",
    "charges": 5000
  }

-- Receive material from colour factory
POST   /api/v1/colouring/batches/:id/receive-from-factory
  Body: {
    "quantity_received": 9.5,
    "wastage": 0.5,
    "date_received": "2026-08-20",
    "coloured_silk_batch_reference": "RED-001"
  }

GET    /api/v1/colouring/status
GET    /api/v1/colouring/pending-batches
GET    /api/v1/colouring/overdue-batches
```

### Coloured Silk Stock
```
GET    /api/v1/coloured-silk/stock
GET    /api/v1/coloured-silk/by-colour/:colour_id
GET    /api/v1/coloured-silk/by-location/:location_id
GET    /api/v1/coloured-silk/:batch_id
```

### Doli Management
```
GET    /api/v1/doli/list?location_id=1&colour_id=2
POST   /api/v1/doli/create
  Body: {
    "source_coloured_silk_batch_id": 123,
    "colour_id": 5,
    "quantity_created": 5,
    "saree_capacity": 10,
    "warping_charges": 500
  }

GET    /api/v1/doli/:id
GET    /api/v1/doli/available
PUT    /api/v1/doli/:id/status
GET    /api/v1/doli/movement-history/:id
```

### Bobbins Management
```
GET    /api/v1/bobbins/stock
GET    /api/v1/bobbins/by-colour/:colour_id
POST   /api/v1/winding/batches
  Body: {
    "source_coloured_silk_batch_id": 123,
    "colour_id": 5,
    "quantity_wound": 10,
    "bobbins_produced": 500,
    "pirns_produced": 200,
    "winding_charges": 1000
  }

GET    /api/v1/winding/batches/:id
GET    /api/v1/winding/status
```

### Inventory Movements
```
GET    /api/v1/inventory/movements?batch_id=123&date_from=2026-08-01&date_to=2026-08-31
GET    /api/v1/inventory/movements/:id
GET    /api/v1/inventory/stock-summary
GET    /api/v1/inventory/by-location/:location_id

-- Manual adjustment (with audit trail)
POST   /api/v1/inventory/adjustment
  Body: {
    "batch_id": 123,
    "adjustment_type": "Damage",
    "quantity": 2.5,
    "notes": "Damaged in storage"
  }
```

### Dashboard (Phase 1)
```
GET    /api/v1/dashboard/inventory-overview
  Response: {
    "raw_silk": { "quantity": 182, "value": "₹XX,XXX" },
    "coloured_silk": { "quantity": 95, "value": "₹XX,XXX" },
    "doli": { "quantity": 42, "value": "₹XX,XXX" },
    "bobbins": { "quantity": 850, "value": "₹XX,XXX" },
    "total_inventory_value": "₹XX,XXX"
  }

GET    /api/v1/dashboard/material-journey
  Response: {
    "at_colour_factory": { "quantity": 50, "batches": [] },
    "at_warping": { "quantity": 30, "batches": [] },
    "at_winding": { "quantity": 20, "batches": [] }
  }

GET    /api/v1/dashboard/supplier-metrics
GET    /api/v1/dashboard/processing-status
```

### Reports (Phase 1)
```
GET    /api/v1/reports/stock-report?format=json
GET    /api/v1/reports/material-location?format=json
GET    /api/v1/reports/inventory-movements?format=json&from=2026-08-01&to=2026-08-31
GET    /api/v1/reports/supplier-report?format=json
GET    /api/v1/reports/colouring-status?format=json

-- Export to Excel/PDF
GET    /api/v1/reports/stock-report?format=excel
GET    /api/v1/reports/stock-report?format=pdf
```

---

## Phase 2 API Endpoints (Additional)

### Production Orders
```
GET    /api/v1/production-orders?status=pending&weaver_id=1
POST   /api/v1/production-orders
GET    /api/v1/production-orders/:id
PUT    /api/v1/production-orders/:id/status
DELETE /api/v1/production-orders/:id
GET    /api/v1/production-orders/weaver/:weaver_id
```

### Material Issue & Return
```
GET    /api/v1/material-issues
POST   /api/v1/material-issues
  Body: {
    "production_order_id": 42,
    "weaver_id": 5,
    "doli_to_issue": 2,
    "bobbins_to_issue": 120
  }

GET    /api/v1/material-issues/:id
POST   /api/v1/material-issues/:id/receive-return
  Body: {
    "doli_returned": 2,
    "bobbins_returned": 120,
    "doli_damaged": 0,
    "bobbins_damaged": 0
  }
```

### Production Outputs
```
GET    /api/v1/production-outputs
POST   /api/v1/production-outputs
  Body: {
    "production_order_id": 42,
    "sarees_produced": 10,
    "sarees_approved": 10,
    "sarees_rejected": 0,
    "completion_date": "2026-08-20"
  }

GET    /api/v1/production-outputs/:id
```

### Sarees
```
GET    /api/v1/sarees?status=in_stock&design_id=1&colour_id=2
POST   /api/v1/sarees
GET    /api/v1/sarees/:id
PUT    /api/v1/sarees/:id/status
GET    /api/v1/sarees/stock-summary
GET    /api/v1/sarees/by-design/:design_id
GET    /api/v1/sarees/by-colour/:colour_id
GET    /api/v1/sarees/valuation
```

### Invoices
```
GET    /api/v1/invoices?buyer_id=1&status=unpaid&from_date=2026-08-01&to_date=2026-08-31
POST   /api/v1/invoices
  Body: {
    "buyer_id": 1,
    "invoice_date": "2026-08-20",
    "items": [
      { "saree_id": 1, "quantity": 5, "rate": 6000 }
    ],
    "discount_amount": 0,
    "notes": ""
  }

GET    /api/v1/invoices/:id
PUT    /api/v1/invoices/:id
DELETE /api/v1/invoices/:id
GET    /api/v1/invoices/:id/print  -- Returns HTML/PDF

-- Add/remove items
POST   /api/v1/invoices/:id/items
DELETE /api/v1/invoices/:id/items/:line_item_id
```

### Sale Payments
```
GET    /api/v1/sale-payments?invoice_id=1&from_date=2026-08-01
POST   /api/v1/sale-payments
  Body: {
    "invoice_id": 1,
    "payment_amount": 50000,
    "payment_date": "2026-08-20",
    "payment_mode": "Bank Transfer",
    "reference_number": "TXN12345"
  }

GET    /api/v1/sale-payments/:id
```

### Weaver Management
```
GET    /api/v1/weavers/:id/ledger?from_date=2026-08-01&to_date=2026-08-31
  Response: [
    { "transaction_type": "Opening Balance", "amount": 2000, "date": "2026-07-01", "balance": 2000 },
    { "transaction_type": "Work Issued", "amount": 5000, "date": "2026-08-14", "balance": 7000 },
    ...
  ]

GET    /api/v1/weavers/:id/transactions
GET    /api/v1/weavers/:id/pending-work
GET    /api/v1/weavers/:id/production-history
```

### Weaver Payments
```
GET    /api/v1/weaver-payments?weaver_id=1&from_date=2026-08-01
POST   /api/v1/weaver-payments
  Body: {
    "weaver_id": 1,
    "production_order_id": 42,
    "payment_amount": 10000,
    "payment_date": "2026-08-20",
    "payment_mode": "Cash"
  }

GET    /api/v1/weaver-payments/:id
```

### Buyer Management
```
GET    /api/v1/buyers/:id/ledger
GET    /api/v1/buyers/:id/invoices
GET    /api/v1/buyers/:id/outstanding
GET    /api/v1/buyers/:id/payment-history
```

### Supplier Management
```
GET    /api/v1/suppliers/:id/ledger
GET    /api/v1/suppliers/:id/purchases
GET    /api/v1/suppliers/:id/payments
GET    /api/v1/suppliers/:id/outstanding
```

### Finance & Accounting
```
GET    /api/v1/expenses?category_id=1&from_date=2026-08-01&to_date=2026-08-31
POST   /api/v1/expenses
GET    /api/v1/expenses/:id
PUT    /api/v1/expenses/:id
PUT    /api/v1/expenses/:id/payment-status

GET    /api/v1/expense-categories
GET    /api/v1/income?from_date=2026-08-01&to_date=2026-08-31

GET    /api/v1/supplier-payments?supplier_id=1&from_date=2026-08-01
POST   /api/v1/supplier-payments
GET    /api/v1/supplier-payments/:id
```

### Reports (Phase 2)
```
GET    /api/v1/reports/profit-loss?from_date=2026-08-01&to_date=2026-08-31&format=json
GET    /api/v1/reports/cash-flow?from_date=2026-08-01&to_date=2026-08-31
GET    /api/v1/reports/cost-per-saree?from_date=2026-08-01&to_date=2026-08-31
GET    /api/v1/reports/inventory-valuation
GET    /api/v1/reports/weaver-production?from_date=2026-08-01&to_date=2026-08-31
GET    /api/v1/reports/design-production?from_date=2026-08-01&to_date=2026-08-31
GET    /api/v1/reports/buyer-outstanding
GET    /api/v1/reports/supplier-outstanding
GET    /api/v1/reports/weaver-outstanding

-- Export options
GET    /api/v1/reports/profit-loss?format=pdf&from_date=...
GET    /api/v1/reports/profit-loss?format=excel&from_date=...
```

---

---

# BACKEND PROJECT STRUCTURE (Go)

---

## Directory Organization

```
weaver-api/
├── cmd/
│   └── server/
│       └── main.go              -- Application entry point
│
├── internal/
│   ├── config/
│   │   └── config.go            -- Configuration management
│   │
│   ├── database/
│   │   ├── postgres.go          -- Database connection
│   │   ├── migrations/          -- SQL migration files
│   │   │   ├── 001_init_schema.sql
│   │   │   ├── 002_add_indices.sql
│   │   │   └── ...
│   │   └── query/               -- Database queries
│   │       ├── suppliers.go
│   │       ├── inventory.go
│   │       └── ...
│   │
│   ├── models/
│   │   ├── supplier.go
│   │   ├── weaver.go
│   │   ├── inventory.go
│   │   └── ...
│   │
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── inventory_service.go
│   │   ├── production_service.go
│   │   ├── sales_service.go
│   │   └── finance_service.go
│   │
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── supplier_handler.go
│   │   ├── inventory_handler.go
│   │   ├── production_handler.go
│   │   ├── sales_handler.go
│   │   └── ...
│   │
│   ├── middleware/
│   │   ├── auth.go              -- Authentication middleware
│   │   ├── error_handler.go     -- Error handling
│   │   ├── request_logger.go    -- Request logging
│   │   └── rate_limit.go        -- Rate limiting
│   │
│   ├── validators/
│   │   ├── supplier_validator.go
│   │   ├── inventory_validator.go
│   │   └── ...
│   │
│   └── utils/
│       ├── jwt.go               -- JWT utilities
│       ├── response.go          -- Response formatting
│       ├── errors.go            -- Custom errors
│       └── ...
│
├── pkg/
│   └── logger/
│       └── logger.go            -- Logging utilities
│
├── tests/
│   ├── unit/
│   │   ├── services/
│   │   └── handlers/
│   ├── integration/
│   └── fixtures/
│       └── test_data.sql
│
├── docker/
│   ├── Dockerfile              -- Docker image for API
│   └── docker-compose.yml      -- Docker compose for local dev
│
├── .env.example               -- Environment variables template
├── go.mod                      -- Go module file
├── go.sum                      -- Go module checksums
└── README.md
```

---

## Go Module Structure

### Main Packages

```go
// cmd/server/main.go
package main

import (
    "github.com/yourcompany/weaver-api/internal/config"
    "github.com/yourcompany/weaver-api/internal/database"
    "github.com/yourcompany/weaver-api/internal/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    // 1. Load configuration
    cfg := config.LoadConfig()
    
    // 2. Connect to database
    db := database.InitDB(cfg.DatabaseURL)
    defer db.Close()
    
    // 3. Run migrations
    database.RunMigrations(db)
    
    // 4. Initialize router
    router := gin.Default()
    
    // 5. Setup routes
    handlers.SetupRoutes(router, db)
    
    // 6. Start server
    router.Run(cfg.ServerPort)
}
```

### Services Layer Example

```go
// internal/services/inventory_service.go
package services

type InventoryService struct {
    db *sql.DB
}

func NewInventoryService(db *sql.DB) *InventoryService {
    return &InventoryService{db: db}
}

// Get current stock for a batch
func (s *InventoryService) GetCurrentStock(batchID int) (float64, error) {
    var quantity float64
    err := s.db.QueryRow(`
        SELECT COALESCE(SUM(quantity), 0)
        FROM inventory_movements
        WHERE batch_id = $1
        AND movement_type NOT IN ('Damage', 'Loss')
    `, batchID).Scan(&quantity)
    
    return quantity, err
}

// Create inventory movement (immutable)
func (s *InventoryService) CreateMovement(movement *InventoryMovement) error {
    result, err := s.db.Exec(`
        INSERT INTO inventory_movements 
        (batch_id, movement_type, quantity, from_location_id, 
         to_location_id, reference_type, reference_id, created_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `, movement.BatchID, movement.MovementType, movement.Quantity,
       movement.FromLocationID, movement.ToLocationID,
       movement.ReferenceType, movement.ReferenceID, movement.CreatedBy)
    
    return err
}
```

---

---

# FRONTEND PROJECT STRUCTURE (React/Next.js)

---

## Directory Organization

```
weaver-frontend/
├── app/                         -- Next.js 14 app directory
│   ├── layout.tsx              -- Root layout
│   ├── page.tsx                -- Home page
│   ├── dashboard/
│   │   ├── page.tsx
│   │   └── layout.tsx
│   ├── inventory/
│   │   ├── raw-silk/
│   │   ├── coloured-silk/
│   │   ├── doli/
│   │   └── bobbins/
│   ├── production/
│   ├── sales/
│   ├── finance/
│   ├── reports/
│   └── admin/
│
├── components/                  -- Reusable React components
│   ├── Auth/
│   │   ├── LoginForm.tsx
│   │   ├── PrivateRoute.tsx
│   │   └── useAuth.ts
│   │
│   ├── Common/
│   │   ├── Header.tsx
│   │   ├── Sidebar.tsx
│   │   ├── Table.tsx
│   │   ├── Form.tsx
│   │   ├── Modal.tsx
│   │   ├── Button.tsx
│   │   ├── Input.tsx
│   │   ├── DatePicker.tsx
│   │   └── Select.tsx
│   │
│   ├── Inventory/
│   │   ├── StockList.tsx
│   │   ├── PurchaseForm.tsx
│   │   ├── ReceiveMaterial.tsx
│   │   └── MovementHistory.tsx
│   │
│   ├── Production/
│   │   ├── ProductionOrderForm.tsx
│   │   ├── IssueToWeaver.tsx
│   │   └── ProductionCompletion.tsx
│   │
│   ├── Sales/
│   │   ├── InvoiceForm.tsx
│   │   ├── InvoiceList.tsx
│   │   └── PaymentRecord.tsx
│   │
│   └── Reports/
│       ├── ReportViewer.tsx
│       ├── DataExport.tsx
│       └── ChartComponent.tsx
│
├── hooks/                       -- Custom React hooks
│   ├── useApi.ts               -- API call hook
│   ├── useForm.ts              -- Form handling
│   ├── usePagination.ts        -- Pagination
│   └── useNotification.ts       -- Notifications
│
├── services/                    -- API service functions
│   ├── api.ts                  -- Axios/fetch configuration
│   ├── authService.ts
│   ├── inventoryService.ts
│   ├── productionService.ts
│   ├── salesService.ts
│   └── financeService.ts
│
├── store/                       -- State management (Zustand/Redux)
│   ├── authStore.ts
│   ├── uiStore.ts
│   └── ...
│
├── types/                       -- TypeScript type definitions
│   ├── index.ts
│   ├── supplier.ts
│   ├── inventory.ts
│   ├── production.ts
│   └── ...
│
├── utils/                       -- Utility functions
│   ├── formatters.ts           -- Number, date formatting
│   ├── validators.ts           -- Input validation
│   ├── constants.ts            -- Constants
│   └── helpers.ts              -- Helper functions
│
├── styles/                      -- Global styles
│   ├── globals.css
│   └── variables.css
│
├── public/                      -- Static assets
│   ├── images/
│   └── icons/
│
├── .env.example               -- Environment variables
├── tsconfig.json              -- TypeScript config
├── tailwind.config.js         -- Tailwind config
├── next.config.js             -- Next.js config
├── package.json
└── README.md
```

---

## API Service Layer

```typescript
// services/api.ts
import axios, { AxiosInstance } from 'axios';

class ApiClient {
  private instance: AxiosInstance;
  
  constructor() {
    this.instance = axios.create({
      baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1',
      timeout: 10000,
    });
    
    // Add request interceptor for auth token
    this.instance.interceptors.request.use((config) => {
      const token = localStorage.getItem('token');
      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    });
    
    // Add response interceptor for error handling
    this.instance.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          // Handle unauthorized
          localStorage.removeItem('token');
          window.location.href = '/login';
        }
        return Promise.reject(error);
      }
    );
  }
  
  get<T>(url: string, config?: any) {
    return this.instance.get<T>(url, config);
  }
  
  post<T>(url: string, data?: any, config?: any) {
    return this.instance.post<T>(url, data, config);
  }
  
  put<T>(url: string, data?: any, config?: any) {
    return this.instance.put<T>(url, data, config);
  }
  
  delete<T>(url: string, config?: any) {
    return this.instance.delete<T>(url, config);
  }
}

export const apiClient = new ApiClient();

// services/inventoryService.ts
export const inventoryService = {
  getRawSilkPurchases: (params?: any) => 
    apiClient.get('/raw-silk/purchases', { params }),
    
  createPurchase: (data: CreatePurchaseRequest) =>
    apiClient.post('/raw-silk/purchases', data),
    
  receiveMaterial: (purchaseId: number, data: ReceiveRequest) =>
    apiClient.post(`/raw-silk/purchases/${purchaseId}/receive`, data),
    
  // ... more methods
};
```

---

---

# DEPLOYMENT ARCHITECTURE

---

## Local Development Setup

### Docker Compose
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:14-alpine
    environment:
      POSTGRES_DB: weaver_db
      POSTGRES_USER: weaver_user
      POSTGRES_PASSWORD: weaver_password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  api:
    build:
      context: .
      dockerfile: docker/Dockerfile
    environment:
      DATABASE_URL: postgres://weaver_user:weaver_password@postgres:5432/weaver_db
      JWT_SECRET: your-secret-key
      ENV: development
    ports:
      - "8080:8080"
    depends_on:
      - postgres

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    environment:
      NEXT_PUBLIC_API_URL: http://localhost:8080/api/v1
    ports:
      - "3000:3000"
    depends_on:
      - api

volumes:
  postgres_data:
```

---

## Production Deployment

### Architecture
```
┌─────────────────────────────────────────────────────┐
│               Internet                              │
└─────────────────────┬───────────────────────────────┘
                      │
                      ▼
         ┌────────────────────────┐
         │   SSL/TLS Certificate  │
         │   (Let's Encrypt)      │
         └────────────┬───────────┘
                      │
                      ▼
         ┌────────────────────────┐
         │   Nginx Reverse Proxy  │
         │   (Load Balancing)     │
         └────────────┬───────────┘
                      │
            ┌─────────┴──────────┐
            ▼                    ▼
   ┌──────────────────┐  ┌──────────────────┐
   │  Go API Instance │  │  Go API Instance │
   │  (Docker)        │  │  (Docker)        │
   │  Port 8080       │  │  Port 8081       │
   └────────┬─────────┘  └────────┬─────────┘
            │                    │
            └─────────┬──────────┘
                      │
                      ▼
         ┌────────────────────────┐
         │  PostgreSQL Database   │
         │  (Managed Service)     │
         │  Automated Backups     │
         └────────────────────────┘

   Separate Static Hosting:
   Frontend (Next.js) → Vercel/Netlify
   CDN for static assets → CloudFlare/AWS CloudFront
```

### Deployment Steps

1. **Prepare VPS**
   ```bash
   # Update system
   sudo apt update && sudo apt upgrade -y
   
   # Install Docker
   curl -fsSL https://get.docker.com -o get-docker.sh
   sudo sh get-docker.sh
   
   # Install Docker Compose
   sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
   sudo chmod +x /usr/local/bin/docker-compose
   
   # Install Nginx
   sudo apt install -y nginx
   ```

2. **Deploy Application**
   ```bash
   # Clone repository
   git clone https://github.com/yourcompany/weaver-management.git
   cd weaver-management
   
   # Create .env file with production secrets
   cp .env.example .env.production
   # Edit .env.production with real values
   
   # Build and run
   docker-compose -f docker-compose.production.yml up -d
   ```

3. **Setup SSL Certificate**
   ```bash
   # Install certbot
   sudo apt install -y certbot python3-certbot-nginx
   
   # Generate certificate
   sudo certbot certonly --standalone -d yourdomain.com
   ```

4. **Configure Nginx**
   ```nginx
   server {
       listen 443 ssl http2;
       server_name yourdomain.com;
       
       ssl_certificate /etc/letsencrypt/live/yourdomain.com/fullchain.pem;
       ssl_certificate_key /etc/letsencrypt/live/yourdomain.com/privkey.pem;
       
       location /api/ {
           proxy_pass http://localhost:8080;
       }
       
       location / {
           proxy_pass http://frontend:3000;
       }
   }
   ```

---

---

# SECURITY CONSIDERATIONS

---

## Authentication & Authorization

1. **JWT Tokens**
   - Short-lived access tokens (15 minutes)
   - Long-lived refresh tokens (7 days)
   - Stored in httpOnly cookies

2. **Password Security**
   - Bcrypt hashing with salt rounds 10+
   - Password requirements: 8+ chars, uppercase, lowercase, number

3. **CORS**
   - Whitelist trusted origins
   - Allow specific methods and headers

---

## Data Protection

1. **Database**
   - Use HTTPS for all connections
   - Implement backups (daily, automated)
   - Encryption at rest (optional, depends on hosting)

2. **Sensitive Data**
   - Bank details stored encrypted
   - PCI compliance for future payment processing
   - GDPR compliance for user data

3. **Audit Trail**
   - All critical actions logged
   - Immutable audit logs
   - Retention: Minimum 7 years

---

## API Security

1. **Rate Limiting**
   - 100 requests/minute per IP
   - 1000 requests/hour per user

2. **Input Validation**
   - All inputs validated server-side
   - SQL injection prevention
   - XSS prevention

3. **Error Handling**
   - No sensitive info in error messages
   - Generic error responses
   - Detailed logs for debugging

---

---

# MONITORING & LOGGING

---

## Logging Strategy

```
Level 1: Application Logs
- Request/response logging
- Business event logging
- Error logging

Level 2: System Logs
- API performance
- Database queries
- Resource usage

Level 3: Security Logs
- Login attempts
- Permission denials
- Data modifications
```

## Health Checks

```
GET /health
Response: {
  "status": "healthy",
  "database": "connected",
  "timestamp": "2026-08-14T10:30:00Z"
}
```

---

**Architecture Version**: 1.0  
**Last Updated**: August 14, 2026
