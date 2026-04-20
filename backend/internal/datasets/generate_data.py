import csv
import random
import uuid
from datetime import datetime, timedelta
from collections import defaultdict

# Set seed for reproducibility
random.seed(42)

# Constants
NUM_CATEGORIES = 50
NUM_PRODUCTS = 10000
NUM_CUSTOMERS = 50000
NUM_ORDERS = 100000
NUM_ORDER_ITEMS = 300000
NUM_INVENTORY = 20000
NUM_WAREHOUSES = 10

# Sample data
FIRST_NAMES = ["John", "Jane", "Michael", "Emily", "David", "Sarah", "James", "Emma", "Robert", "Olivia", "William", "Sophia", "Richard", "Isabella", "Thomas", "Mia", "Charles", "Charlotte", "Daniel", "Amelia"]
LAST_NAMES = ["Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin"]
COUNTRIES = ["US", "UK", "CA", "DE", "FR", "AU", "JP", "BR", "IN", "MX", "ES", "IT", "NL", "SE", "NO"]
STATUSES = ["active", "inactive", "pending"]
ORDER_STATUSES = ["pending", "processing", "shipped", "delivered", "cancelled"]
PRODUCT_NAMES = ["Laptop", "Phone", "Tablet", "Monitor", "Keyboard", "Mouse", "Headphones", "Camera", "Speaker", "Watch", "Charger", "Cable", "Stand", "Case", "Cover", "Screen", "Battery", "Memory", "Storage", "Adapter"]
PRODUCT_ADJECTIVES = ["Premium", "Basic", "Pro", "Elite", "Ultra", "Mini", "Max", "Plus", "Wireless", "Portable", "Compact", "Standard", "Advanced", "Essential", "Smart"]
CATEGORIES = ["Electronics", "Computers", "Audio", "Accessories", "Gaming", "Home", "Office", "Photography", "Wearables", "Storage"]

# Generate warehouses
warehouses = [f"WH-{str(i).zfill(3)}" for i in range(1, NUM_WAREHOUSES + 1)]

def generate_id(prefix):
    """Generate a unique ID with prefix"""
    return f"{prefix}-{uuid.uuid4().hex[:12]}"

def generate_timestamp(start_date, end_date):
    """Generate a random timestamp between start and end date"""
    delta = end_date - start_date
    random_seconds = random.randint(0, int(delta.total_seconds()))
    return start_date + timedelta(seconds=random_seconds)

def generate_categories():
    """Generate categories.csv"""
    print("Generating categories.csv...")
    with open('ecommerce/categories.csv', 'w', newline='') as f:
        writer = csv.writer(f)
        writer.writerow(['category_id', 'name', 'parent_category_id'])
        
        # Generate main categories
        main_categories = []
        for i in range(len(CATEGORIES)):
            cat_id = f"CAT-{str(i+1).zfill(3)}-{uuid.uuid4().hex[:8]}"
            name = CATEGORIES[i]
            parent_id = None
            main_categories.append((cat_id, name, parent_id))
            writer.writerow([cat_id, name, parent_id])
        
        # Generate subcategories
        for i in range(len(CATEGORIES), NUM_CATEGORIES):
            cat_id = f"CAT-{str(i+1).zfill(3)}-{uuid.uuid4().hex[:8]}"
            parent = random.choice(main_categories)
            name = f"{parent[1]} {random.choice(['Pro', 'Basic', 'Mini', 'Max', 'Plus', 'Premium', 'Standard'])}"
            writer.writerow([cat_id, name, parent[0]])
        
    return main_categories

def generate_products(categories):
    """Generate products.csv"""
    print("Generating products.csv...")
    with open('ecommerce/products.csv', 'w', newline='') as f:
        writer = csv.writer(f)
        writer.writerow(['product_id', 'name', 'category_id', 'price', 'description', 'created_at'])
        
        start_date = datetime(2020, 1, 1)
        end_date = datetime(2024, 12, 31)
        
        for i in range(NUM_PRODUCTS):
            prod_id = f"PROD-{str(i+1).zfill(8)}"
            adj = random.choice(PRODUCT_ADJECTIVES)
            base_name = random.choice(PRODUCT_NAMES)
            name = f"{adj} {base_name} {random.choice(['Model', 'Series', 'Edition', 'Version', ''])} {random.randint(1, 999)}"
            category = random.choice(categories)
            price = round(random.uniform(9.99, 2999.99), 2)
            description = f"High-quality {base_name.lower()} for all your needs. {adj} edition with advanced features."
            created_at = generate_timestamp(start_date, end_date).isoformat()
            
            writer.writerow([prod_id, name, category[0], price, description, created_at])
            
            if (i + 1) % 1000 == 0:
                print(f"  {i + 1}/{NUM_PRODUCTS} products generated")
    
    return True

def generate_customers():
    """Generate customers.csv"""
    print("Generating customers.csv...")
    with open('ecommerce/customers.csv', 'w', newline='') as f:
        writer = csv.writer(f)
        writer.writerow(['customer_id', 'email', 'name', 'country', 'phone', 'created_at', 'status', 'password_hash'])
        
        start_date = datetime(2018, 1, 1)
        end_date = datetime(2024, 12, 31)
        used_emails = set()
        
        for i in range(NUM_CUSTOMERS):
            cust_id = f"CUST-{str(i+1).zfill(8)}"
            
            # Generate unique email
            while True:
                first = random.choice(FIRST_NAMES).lower()
                last = random.choice(LAST_NAMES).lower()
                num = random.randint(1, 999)
                email = f"{first}.{last}{num}@example.com"
                if email not in used_emails:
                    used_emails.add(email)
                    break
            
            name = f"{random.choice(FIRST_NAMES)} {random.choice(LAST_NAMES)}"
            country = random.choice(COUNTRIES)
            phone = f"+1-{random.randint(100,999)}-{random.randint(100,999)}-{random.randint(1000,9999)}"
            created_at = generate_timestamp(start_date, end_date).isoformat()
            status = random.choice(STATUSES)
            password_hash = f"$2a$10${uuid.uuid4().hex[:22]}${uuid.uuid4().hex[:31]}"
            
            writer.writerow([cust_id, email, name, country, phone, created_at, status, password_hash])
            
            if (i + 1) % 5000 == 0:
                print(f"  {i + 1}/{NUM_CUSTOMERS} customers generated")
    
    return True

def generate_orders():
    """Generate orders.csv"""
    print("Generating orders.csv...")
    with open('ecommerce/orders.csv', 'w', newline='') as f:
        writer = csv.writer(f)
        writer.writerow(['order_id', 'customer_id', 'order_date', 'status', 'total_amount', 'shipping_address'])
        
        start_date = datetime(2019, 1, 1)
        end_date = datetime(2024, 12, 31)
        
        customer_ids = [f"CUST-{str(i+1).zfill(8)}" for i in range(NUM_CUSTOMERS)]
        addresses = ["123 Main St", "456 Oak Ave", "789 Pine Rd", "321 Elm Blvd", "654 Maple Dr", "987 Cedar Ln", "147 Birch Way", "258 Spruce Ct", "369 Willow Pl", "741 Cherry Rd"]
        
        for i in range(NUM_ORDERS):
            order_id = f"ORD-{str(i+1).zfill(8)}"
            customer_id = random.choice(customer_ids)
            order_date = generate_timestamp(start_date, end_date).isoformat()
            status = random.choice(ORDER_STATUSES)
            total_amount = round(random.uniform(10.0, 5000.0), 2)
            address = f"{random.randint(1, 999)} {random.choice(addresses).split()[1]}, {random.choice(COUNTRIES)}"
            
            writer.writerow([order_id, customer_id, order_date, status, total_amount, address])
            
            if (i + 1) % 10000 == 0:
                print(f"  {i + 1}/{NUM_ORDERS} orders generated")
    
    return True

def generate_order_items():
    """Generate order_items.csv"""
    print("Generating order_items.csv...")
    with open('ecommerce/order_items.csv', 'w', newline='') as f:
        writer = csv.writer(f)
        writer.writerow(['order_item_id', 'order_id', 'product_id', 'quantity', 'unit_price'])
        
        order_ids = [f"ORD-{str(i+1).zfill(8)}" for i in range(NUM_ORDERS)]
        product_ids = [f"PROD-{str(i+1).zfill(8)}" for i in range(NUM_PRODUCTS)]
        
        items_per_order = NUM_ORDER_ITEMS // NUM_ORDERS
        
        for i in range(NUM_ORDER_ITEMS):
            item_id = f"OI-{str(i+1).zfill(10)}"
            order_id = order_ids[i // items_per_order] if i // items_per_order < NUM_ORDERS else order_ids[-1]
            product_id = random.choice(product_ids)
            quantity = random.randint(1, 5)
            unit_price = round(random.uniform(5.0, 2000.0), 2)
            
            writer.writerow([item_id, order_id, product_id, quantity, unit_price])
            
            if (i + 1) % 30000 == 0:
                print(f"  {i + 1}/{NUM_ORDER_ITEMS} order items generated")
    
    return True

def generate_inventory():
    """Generate inventory.csv"""
    print("Generating inventory.csv...")
    with open('ecommerce/inventory.csv', 'w', newline='') as f:
        writer = csv.writer(f)
        writer.writerow(['product_id', 'warehouse_id', 'quantity', 'last_updated'])
        
        product_ids = [f"PROD-{str(i+1).zfill(8)}" for i in range(NUM_PRODUCTS)]
        used_combinations = set()
        
        for i in range(NUM_INVENTORY):
            while True:
                product_id = random.choice(product_ids)
                warehouse_id = random.choice(warehouses)
                key = (product_id, warehouse_id)
                if key not in used_combinations:
                    used_combinations.add(key)
                    break
            
            quantity = random.randint(0, 500)
            last_updated = datetime.now().isoformat()
            
            writer.writerow([product_id, warehouse_id, quantity, last_updated])
            
            if (i + 1) % 2000 == 0:
                print(f"  {i + 1}/{NUM_INVENTORY} inventory records generated")
    
    return True

def main():
    print("Starting data generation...")
    print(f"Categories: {NUM_CATEGORIES}")
    print(f"Products: {NUM_PRODUCTS}")
    print(f"Customers: {NUM_CUSTOMERS}")
    print(f"Orders: {NUM_ORDERS}")
    print(f"Order Items: {NUM_ORDER_ITEMS}")
    print(f"Inventory: {NUM_INVENTORY}")
    print("-" * 50)
    
    categories = generate_categories()
    generate_products(categories)
    generate_customers()
    generate_orders()
    generate_order_items()
    generate_inventory()
    
    print("-" * 50)
    print("Data generation complete!")
    
    # Print file sizes
    import os
    for filename in ['categories.csv', 'products.csv', 'customers.csv', 'orders.csv', 'order_items.csv', 'inventory.csv']:
        path = f'ecommerce/{filename}'
        if os.path.exists(path):
            size = os.path.getsize(path) / (1024 * 1024)  # MB
            print(f"  {filename}: {size:.2f} MB")

if __name__ == "__main__":
    main()
