package database

import (
	"database/sql"
	"fmt"
)

func SeedDemoData(db *sql.DB) error {
	// Check if data already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM suppliers").Scan(&count)
	if err == nil && count > 0 {
		fmt.Println("✓ Demo data already exists")
		return nil
	}

	// Insert products
	productsSQL := `
		INSERT INTO products (product_code, name, category, unit, status) VALUES
		('P001', 'Kachha Silk', 'Raw Material', 'kg', 'Active'),
		('P002', 'Coloured Silk', 'Processed', 'kg', 'Active'),
		('P003', 'Doli', 'Processed', 'unit', 'Active'),
		('P004', 'Bobbins', 'Processed', 'unit', 'Active'),
		('P005', 'Pirns', 'Processed', 'unit', 'Active'),
		('P006', 'Saree', 'Finished', 'piece', 'Active')
		ON CONFLICT DO NOTHING;
	`
	if _, err := db.Exec(productsSQL); err != nil {
		return fmt.Errorf("failed to insert products: %v", err)
	}

	// Insert colours
	coloursSQL := `
		INSERT INTO colours (colour_code, name, hex_code, status) VALUES
		('C001', 'Red', '#FF0000', 'Active'),
		('C002', 'Blue', '#0000FF', 'Active'),
		('C003', 'Green', '#00AA00', 'Active'),
		('C004', 'Maroon', '#800000', 'Active'),
		('C005', 'Yellow', '#FFFF00', 'Active'),
		('C006', 'Black', '#000000', 'Active'),
		('C007', 'White', '#FFFFFF', 'Active')
		ON CONFLICT DO NOTHING;
	`
	if _, err := db.Exec(coloursSQL); err != nil {
		return fmt.Errorf("failed to insert colours: %v", err)
	}

	// Insert locations
	locationsSQL := `
		INSERT INTO locations (location_code, name, type, status) VALUES
		('L001', 'Main Warehouse', 'Warehouse', 'Active'),
		('L002', 'Colour Factory', 'Processing', 'Active'),
		('L003', 'Warping Unit', 'Processing', 'Active'),
		('L004', 'Winding Unit', 'Processing', 'Active'),
		('L005', 'Finished Goods Store', 'Finished Goods', 'Active')
		ON CONFLICT DO NOTHING;
	`
	if _, err := db.Exec(locationsSQL); err != nil {
		return fmt.Errorf("failed to insert locations: %v", err)
	}

	// Insert suppliers
	suppliersSQL := `
		INSERT INTO suppliers (supplier_code, name, phone, email, city, material_type, status) VALUES
		('SUP001', 'ABC Silk Supplier', '9876543210', 'abc@supplier.com', 'Bangalore', 'Raw Silk', 'Active'),
		('SUP002', 'XYZ Colour Factory', '9876543211', 'xyz@colour.com', 'Mumbai', 'Colour Factory', 'Active'),
		('SUP003', 'Premium Silk Co', '9876543212', 'premium@silk.com', 'Chennai', 'Raw Silk', 'Active'),
		('SUP004', 'Warping Services', '9876543213', 'warp@services.com', 'Bangalore', 'Warping', 'Active'),
		('SUP005', 'Winding Unit', '9876543214', 'wind@unit.com', 'Bangalore', 'Winding', 'Active')
		ON CONFLICT DO NOTHING;
	`
	if _, err := db.Exec(suppliersSQL); err != nil {
		return fmt.Errorf("failed to insert suppliers: %v", err)
	}

	// Insert weavers
	weaversSQL := `
		INSERT INTO weavers (weaver_code, name, phone, email, village, joining_date, status) VALUES
		('WV001', 'Ravi Kumar', '9876543220', 'ravi@weaver.com', 'Kanchipuram', '2024-01-15', 'Active'),
		('WV002', 'Lakshmi Devi', '9876543221', 'lakshmi@weaver.com', 'Kanchipuram', '2024-02-10', 'Active'),
		('WV003', 'Rajesh Singh', '9876543222', 'rajesh@weaver.com', 'Mysore', '2024-03-05', 'Active'),
		('WV004', 'Meena Sharma', '9876543223', 'meena@weaver.com', 'Kanchipuram', '2024-01-20', 'Active'),
		('WV005', 'Arjun Patel', '9876543224', 'arjun@weaver.com', 'Bangalore', '2024-04-01', 'Active')
		ON CONFLICT DO NOTHING;
	`
	if _, err := db.Exec(weaversSQL); err != nil {
		return fmt.Errorf("failed to insert weavers: %v", err)
	}

	// Insert buyers
	buyersSQL := `
		INSERT INTO buyers (buyer_code, name, business_name, phone, email, city, status) VALUES
		('BUY001', 'Rajesh Textiles', 'Rajesh & Co', '9876543230', 'rajesh@textiles.com', 'Delhi', 'Active'),
		('BUY002', 'Luxe Fabrics', 'Luxe Fabrics Ltd', '9876543231', 'luxe@fabrics.com', 'Mumbai', 'Active'),
		('BUY003', 'Premium Sarees', 'Premium Saree House', '9876543232', 'premium@sarees.com', 'Bangalore', 'Active'),
		('BUY004', 'Ethnic Wear Co', 'Ethnic Wear Company', '9876543233', 'ethnic@wear.com', 'Kolkata', 'Active'),
		('BUY005', 'Traditional Textiles', 'Traditional & Modern', '9876543234', 'trad@textiles.com', 'Chennai', 'Active')
		ON CONFLICT DO NOTHING;
	`
	if _, err := db.Exec(buyersSQL); err != nil {
		return fmt.Errorf("failed to insert buyers: %v", err)
	}

	// Insert designs
	designsSQL := `
		INSERT INTO designs (design_code, name, saree_type, status) VALUES
		('D001', 'Classic Kanchipuram', 'Regular', 'Active'),
		('D002', 'Modern Silk', 'Premium', 'Active'),
		('D003', 'Traditional Gold', 'Wedding', 'Active'),
		('D004', 'Contemporary Design', 'Regular', 'Active'),
		('D005', 'Festive Collection', 'Regular', 'Active')
		ON CONFLICT DO NOTHING;
	`
	if _, err := db.Exec(designsSQL); err != nil {
		return fmt.Errorf("failed to insert designs: %v", err)
	}

	fmt.Println("✓ Demo data seeded successfully")
	return nil
}
