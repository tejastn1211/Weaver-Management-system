# Weaving Business Management System - Project Documentation

A comprehensive digital management system for traditional weaving/silk businesses to replace manual notebooks, WhatsApp records, and memory-based tracking.

---

## 📋 Quick Navigation

- **[DEVELOPMENT_PHASES.md](./DEVELOPMENT_PHASES.md)** - Two-phase development plan (Phase 1 & 2)
- **[DATABASE_SCHEMA.md](./DATABASE_SCHEMA.md)** - Complete PostgreSQL schema design
- **[TECHNICAL_ARCHITECTURE.md](./TECHNICAL_ARCHITECTURE.md)** - API design and system architecture

---

## 🎯 Project Overview

### Business Objective
Build a digital management system that enables the business owner to answer critical business questions:

1. **How much raw silk do I currently have?**
2. **Where is every material currently located?**
3. **How much do I owe each weaver and how much are they owed?**
4. **How many sarees are in stock and what are they worth?**
5. **Who bought what and how much do they owe me?**
6. **How much money came in, went out, and what is my profit?**

---

## 📊 Two-Phase Development Plan

### Phase 1: Foundation & Inventory System (6-8 weeks)
**Goal**: Build a complete inventory tracking system with full visibility into material movements.

**Key Modules**:
- User Management & Authentication
- Master Data (Suppliers, Weavers, Buyers, Products, Colours, Locations)
- Raw Silk Purchase Management
- Colouring Process Tracking
- Coloured Silk Stock Management
- Doli (Warping) Management
- Bobbins/Pirns Management
- Inventory Movement Ledger
- Dashboard (Inventory KPIs)
- Basic Reports

**Success Criteria**:
✅ System can track all material movements  
✅ Material location is always known  
✅ Inventory is 100% accurate  
✅ Dashboard shows real-time stock levels  

---

### Phase 2: Production, Finance & Reports (6-8 weeks)
**Goal**: Complete the business workflow with production tracking, sales, and financial management.

**Key Modules**:
- Weaver Management & Material Tracking
- Production Management (Work Orders, Completion)
- Saree Stock Management
- Sales & Billing (Invoices, Payment)
- Income & Expense Tracking
- Complete Ledgers (Weaver, Buyer, Supplier)
- Financial Reports (P&L, Cash Flow, Cost Analysis)
- Enhanced Dashboard (Complete Business View)

**Success Criteria**:
✅ Complete business workflow is digitized  
✅ Profit/loss is calculated accurately  
✅ All ledgers are reconciled  
✅ Business owner has complete visibility  

---

## 🛠 Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Gin or Fiber
- **Database**: PostgreSQL 14+

### Frontend
- **Framework**: React 18+ / Next.js 14+
- **Styling**: Tailwind CSS
- **State Management**: React Query + Zustand

### DevOps
- **Containerization**: Docker
- **Deployment**: Docker on VPS
- **Web Server**: Nginx

---

## 📚 Documentation Overview

| Document | Purpose | Audience |
|----------|---------|----------|
| **DEVELOPMENT_PHASES.md** | Detailed phase-by-phase breakdown with modules, APIs, screens, deliverables | PMs, Developers, Managers |
| **DATABASE_SCHEMA.md** | Complete PostgreSQL schema with all tables and relationships | Database Engineers, Developers |
| **TECHNICAL_ARCHITECTURE.md** | System architecture, API design, project structure | Architects, Tech Leads, Developers |
| **README.md** (this file) | Quick reference and navigation | Everyone |

---

## 🚀 Quick Start by Role

### Project Manager
1. Read Phase 1 & 2 objectives in **DEVELOPMENT_PHASES.md**
2. Review success criteria and deliverables
3. Plan resource allocation and sprints

### Database Engineer
1. Study **DATABASE_SCHEMA.md** for complete schema
2. Review table relationships and constraints
3. Implement migration scripts

### Backend Developer (Go)
1. Review **TECHNICAL_ARCHITECTURE.md** → Backend architecture
2. Study **DATABASE_SCHEMA.md** → Data structure
3. Reference **DEVELOPMENT_PHASES.md** → Phase 1 requirements
4. Implement REST API endpoints from **TECHNICAL_ARCHITECTURE.md**

### Frontend Developer (React)
1. Review **TECHNICAL_ARCHITECTURE.md** → Frontend structure
2. Check Phase 1 screens in **DEVELOPMENT_PHASES.md**
3. Integrate with API endpoints from **TECHNICAL_ARCHITECTURE.md**

### DevOps Engineer
1. Review deployment architecture in **TECHNICAL_ARCHITECTURE.md**
2. Set up Docker, Nginx, and SSL configuration
3. Prepare development and production environments

---

## ✨ Key Features

### Inventory Management
✅ Immutable transaction ledger (complete audit trail)  
✅ Multi-location tracking (warehouse, factory, weaver, etc.)  
✅ Real-time stock levels by product, colour, location  
✅ Material journey visibility  
✅ Automatic movement tracking  

### Weaver Management
✅ Weaver master data and personal details  
✅ Material issue/return tracking  
✅ Production completion recording  
✅ Weaver ledger with balance tracking  
✅ Payment tracking  

### Sales & Billing
✅ Invoice generation  
✅ Payment recording  
✅ Buyer ledger with outstanding balance  
✅ Multiple payment modes  

### Financial Reporting
✅ Profit & Loss statement  
✅ Cash flow report  
✅ Cost per saree analysis  
✅ Outstanding receivables/payables  

---

## 💾 Database Architecture

### Key Principles
1. **Immutable Transactions**: No updates to inventory_movements (only inserts)
2. **Location Tracking**: Every material knows its location
3. **Audit Trail**: All critical operations are logged
4. **Referential Integrity**: Foreign keys and constraints ensure data consistency

### Core Tables (Phase 1)
```
- users, roles, permissions
- suppliers, weavers, buyers, products, colours, locations
- inventory_batches, inventory_movements
- raw_silk_purchases, colouring_batches, doli_batches, winding_batches
- audit_logs
```

### Additional Tables (Phase 2)
```
- production_orders, material_issues, material_returns, production_outputs
- sarees, saree_movements
- invoices, invoice_items, sale_payments
- weaver_transactions, weaver_payments
- buyer_transactions, expenses, income_transactions
```

---

## 🔌 API Design

### Key Endpoints (Phase 1)

**Authentication**:
```
POST   /api/v1/auth/login
POST   /api/v1/auth/logout
GET    /api/v1/auth/me
```

**Raw Silk Purchases**:
```
GET    /api/v1/raw-silk/purchases
POST   /api/v1/raw-silk/purchases
POST   /api/v1/raw-silk/purchases/:id/receive
```

**Inventory**:
```
GET    /api/v1/inventory/stock-summary
GET    /api/v1/inventory/by-location/:location_id
GET    /api/v1/inventory/movements
```

**Dashboard**:
```
GET    /api/v1/dashboard/inventory-overview
GET    /api/v1/dashboard/material-journey
```

See **TECHNICAL_ARCHITECTURE.md** for complete API reference.

---

## 📂 Project Structure

```
Weaver-Management-system/
├── README.md (this file)
├── DEVELOPMENT_PHASES.md
├── DATABASE_SCHEMA.md
├── TECHNICAL_ARCHITECTURE.md
│
├── backend/
│   ├── cmd/server/
│   ├── internal/
│   │   ├── database/
│   │   ├── handlers/
│   │   ├── services/
│   │   └── models/
│   ├── docker/
│   └── tests/
│
├── frontend/
│   ├── app/
│   ├── components/
│   ├── services/
│   ├── types/
│   └── public/
│
└── docker-compose.yml
```

---

## 🏁 Getting Started

### Prerequisites
```bash
# Backend
Go 1.21+
PostgreSQL 14+

# Frontend
Node.js 18+
npm or yarn

# DevOps
Docker
Docker Compose
```

### Local Development Setup
```bash
# 1. Clone repository
git clone https://github.com/yourcompany/weaver-management.git
cd weaver-management

# 2. Start database
docker-compose up -d postgres

# 3. Backend setup
cd backend
go mod download
go run cmd/server/main.go

# 4. Frontend setup (new terminal)
cd frontend
npm install
npm run dev

# 5. Access application
# Frontend: http://localhost:3000
# API: http://localhost:8080
```

---

## 📊 Success Criteria

### Phase 1 Success
The system can answer these 5 questions perfectly:
1. What raw materials do I currently have in stock?
2. How much material is under processing?
3. What is the current value of my inventory?
4. Which supplier provided which material?
5. Where is every piece of material located?

### Phase 2 Success
The system can answer these 10 questions perfectly:
1. What has each weaver produced and how much are they owed?
2. How many sarees are in stock?
3. Who bought what and how much do they owe?
4. What is my total income and expenses?
5. What is my current profit/loss?
6. What is my cost per saree?
7. How much money is outstanding?
8. What is inventory value?
9. Which weavers are most productive?
10. Which designs/colours are best sellers?

---

## 🔒 Security

### Authentication
- JWT-based with refresh tokens
- Role-based access control (Admin, Manager, Accountant)
- Session management

### Data Protection
- Encrypted database connections
- Password hashing (bcrypt)
- Input validation and sanitization
- Audit trail for all operations

### Compliance
- Daily automated backups
- 30-day retention
- GDPR and data protection compliance

---

## 📞 Team & Contact

**Project Owner**: [Name]  
**Technical Lead**: [Name]  
**Project Manager**: [Name]  

---

## 📅 Timeline

```
Phase 1: Week 1-8 (Foundation & Inventory)
Phase 2: Week 9-16 (Production, Finance & Reports)

Total: 16 weeks to complete system
```

---

## 📝 Document References

- **Phase Planning**: See [DEVELOPMENT_PHASES.md](./DEVELOPMENT_PHASES.md)
- **Database Design**: See [DATABASE_SCHEMA.md](./DATABASE_SCHEMA.md)
- **API & Architecture**: See [TECHNICAL_ARCHITECTURE.md](./TECHNICAL_ARCHITECTURE.md)

---

**Last Updated**: August 14, 2026  
**Status**: Planning & Design Phase  
**Next Review**: After Phase 1 Completion