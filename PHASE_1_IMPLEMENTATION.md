  # Phase 1 Implementation Guide

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Node.js 18+ (for local frontend development)
- Go 1.21+ (for local backend development)
- PostgreSQL 15+ (optional, if not using Docker)

### Option 1: Run with Docker Compose (Recommended)

```bash
# Start all services
docker-compose up --build

# Services will be available at:
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
# PostgreSQL: localhost:5432
```

### Option 2: Local Development

#### Backend
```bash
cd backend

# Install dependencies
go mod download

# Create .env file (if not exists)
cp .env.example .env

# Run migrations and start server
go run ./cmd/server/main.go

# Server will run on http://localhost:8080
```

#### Frontend
```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev

# App will run on http://localhost:3000
```

## 🔐 Demo Login Credentials

Use any of these accounts to login:

| Username | Password | Role |
|----------|----------|------|
| admin | admin123 | Admin |
| manager | manager123 | Manager |
| accountant | accountant123 | Accountant |

## 📋 Phase 1 Features Implemented

### Masters Management
- ✅ **Suppliers** - Add, view, edit, delete supplier information
- ✅ **Weavers** - Manage weaver details and information
- ✅ **Buyers** - Track buyer/customer information
- ✅ **Products** - Define products and material types
- ✅ **Colours** - Manage colour codes and hex values
- ✅ **Locations** - Warehouse and processing locations

### Inventory & Purchases
- ✅ **Raw Silk Purchases** - Purchase management from suppliers
- ✅ **Colouring Batches** - Track colouring process and batches
- ✅ **Inventory Movements** - Track stock movements and adjustments
- ✅ **Inventory Stock** - Current stock levels by location and colour

### Dashboard & Reporting
- ✅ **Dashboard** - KPI overview (suppliers, weavers, buyers, inventory value)
- ✅ **Inventory Report** - Stock levels and movements

## 🏗️ Architecture

### Backend (Go + Gin)
- **Port**: 8080
- **Database**: PostgreSQL
- **Authentication**: Demo hardcoded users (for demo purposes)
- **API Format**: REST JSON

Key Files:
- `backend/cmd/server/main.go` - Application entry point
- `backend/internal/handlers/` - API endpoint handlers
- `backend/internal/database/` - Database initialization and migrations
- `backend/internal/config/` - Configuration management

### Frontend (Next.js + React + Material-UI)
- **Port**: 3000
- **Framework**: Next.js 14
- **UI Library**: Material-UI (MUI)
- **State Management**: Zustand
- **API Client**: Axios

Key Files:
- `frontend/app/` - Pages and layouts
- `frontend/components/` - React components
- `frontend/lib/` - API client and utilities
- `frontend/hooks/` - Custom React hooks

### Database (PostgreSQL)
- **Connection**: `postgres://weaver_user:weaver_password@localhost:5432/weaver_db`
- **Tables**: 17 tables covering all Phase 1 entities
- **Migrations**: Automatic on application start

## 📊 API Endpoints

### Authentication
- `POST /api/auth/login` - User login
- `GET /api/auth/profile` - Get current user profile
- `POST /api/auth/logout` - User logout

### Masters
- `GET/POST /api/suppliers` - List and create suppliers
- `GET/PUT/DELETE /api/suppliers/:id` - Supplier operations
- `GET/POST /api/weavers` - List and create weavers
- `GET/POST /api/buyers` - List and create buyers

### Inventory
- `GET /api/raw-silk-purchases` - Raw silk purchases
- `POST /api/raw-silk-purchases` - Create purchase
- `GET /api/colouring` - Colouring batches
- `POST /api/colouring` - Create colouring batch
- `GET /api/inventory/stock` - Current inventory stock
- `GET /api/inventory/movements` - Inventory movements history

### Dashboard
- `GET /api/dashboard/stats` - Dashboard statistics

## 🔄 Sample Data

Demo data is automatically seeded on application startup:
- 5 Suppliers (Raw silk vendors, colour factories)
- 7 Colours (Red, Blue, Green, Maroon, Yellow, Black, White)
- 5 Locations (Warehouse, Colour Factory, Warping, Winding, Finished Goods)
- 5 Weavers
- 5 Buyers
- 6 Products
- 5 Designs

## 🛠️ Development

### Running Tests
```bash
# Backend tests
cd backend
go test ./...

# Frontend tests (if configured)
cd frontend
npm test
```

### Building for Production
```bash
# Build Docker images
docker-compose build

# Or individually
docker build -t weaver-api -f Dockerfile .
docker build -t weaver-frontend -f frontend/Dockerfile .
```

## 📝 Database Schema

Key tables:
- `users` - System users
- `suppliers` - Supplier/vendor information
- `weavers` - Weaver details
- `buyers` - Buyer/customer information
- `products` - Product definitions
- `colours` - Colour master data
- `locations` - Warehouse and processing locations
- `raw_silk_purchases` - Raw silk purchase orders
- `colouring_batches` - Colouring process batches
- `doli_batches` - Warping batches
- `winding_batches` - Winding batches
- `bobbin_inventory` - Bobbin stock tracking
- `inventory_batches` - Inventory batch tracking
- `inventory_movements` - Append-only inventory movement ledger
- `supplier_transactions` - Supplier financial transactions
- `audit_logs` - System audit trail

## 🚀 Deployment

### Docker Compose (Local)
```bash
docker-compose up -d
```

### Production Deployment
For production deployment, follow the TECHNICAL_ARCHITECTURE.md guide for:
- SSL/TLS certificate setup
- Environment variable configuration
- Database backups
- Monitoring and logging setup

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Change ports in docker-compose.yml or kill existing processes
lsof -i :8080  # Find process on port 8080
kill -9 <PID>  # Kill the process
```

### Database Connection Issues
```bash
# Verify PostgreSQL is running
docker-compose ps

# Check logs
docker-compose logs postgres
docker-compose logs backend
```

### Frontend Not Loading
```bash
# Clear node modules and reinstall
cd frontend
rm -rf node_modules
npm install
npm run dev
```

## 📞 Support

For issues or questions:
1. Check the logs: `docker-compose logs -f`
2. Review DATABASE_SCHEMA.md for database structure
3. Check TECHNICAL_ARCHITECTURE.md for API details

## 📄 Documentation

- `README.md` - Project overview
- `DEVELOPMENT_PHASES.md` - Phase 1 & 2 detailed requirements
- `DATABASE_SCHEMA.md` - Complete database schema
- `TECHNICAL_ARCHITECTURE.md` - API design and architecture
- `PHASE_1_IMPLEMENTATION.md` - This file

---

**Version**: 1.0.0 (Phase 1)  
**Status**: ✅ Implemented and Ready for Demo  
**Last Updated**: 2024
