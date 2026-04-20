#!/usr/bin/env python3
"""
E-commerce Dataset Generator

Generates realistic e-commerce data with configurable edge cases and referential integrity.

Usage:
    python generate_data.py
    python generate_data.py --customers 50000 --products 10000 --orders 100000
    python generate_data.py --clean --strict
    python generate_data.py --start-date 2020-01-01 --end-date 2025-12-31

Requirements:
    From datasets/ directory: pip install -r requirements.txt
"""

import argparse
import csv
import random
import uuid
from datetime import datetime, timedelta
from pathlib import Path

try:
    from faker import Faker
    from tqdm import tqdm
except ImportError as e:
    print(f"Error: Missing required library. Please install dependencies:")
    print(f"    cd datasets/")
    print(f"    pip install -r requirements.txt")
    print(f"\nSee datasets/README.md for detailed installation instructions.")
    exit(1)


# Default configuration
DEFAULT_CUSTOMERS = 50000
DEFAULT_PRODUCTS = 10000
DEFAULT_ORDERS = 100000
DEFAULT_CATEGORIES = 50
DEFAULT_WAREHOUSES = 10

# Edge case percentages (when not in --clean mode)
EDGE_CASE_CONFIG = {
    'minor_issues_pct': 0.15,  # 15% minor issues
    'major_issues_pct': 0.05,  # 5% major issues
    'orphaned_records_pct': 0.02,  # 2% orphaned foreign keys (when not --strict)
}


class EcommerceDataGenerator:
    """Generates realistic e-commerce dataset with configurable edge cases."""
    
    def __init__(self, clean_mode=False, strict_mode=False, start_date=None, end_date=None, seed=None):
        self.fake = Faker()
        if seed:
            Faker.seed(seed)
            random.seed(seed)
        
        self.clean_mode = clean_mode
        self.strict_mode = strict_mode
        self.start_date = start_date or datetime(2023, 1, 1)
        self.end_date = end_date or datetime.now()
        
        # Storage for cross-references
        self.customer_ids = []
        self.product_ids = []
        self.order_ids = []
        self.category_ids = []
        self.warehouse_ids = []
    
    def random_date(self, start=None, end=None):
        """Generate random datetime between start and end dates."""
        start = start or self.start_date
        end = end or self.end_date
        delta = end - start
        
        # Handle edge case where start >= end
        if delta.days < 0:
            return start
        elif delta.days == 0:
            # Same day, just vary the time
            random_seconds = random.randint(0, min(86400, delta.seconds))
            return start + timedelta(seconds=random_seconds)
        
        random_days = random.randint(0, delta.days)
        random_seconds = random.randint(0, 86400)
        return start + timedelta(days=random_days, seconds=random_seconds)
    
    def should_add_edge_case(self, severity='minor'):
        """Determine if edge case should be added based on configuration."""
        if self.clean_mode:
            return False
        
        if severity == 'minor':
            return random.random() < EDGE_CASE_CONFIG['minor_issues_pct']
        elif severity == 'major':
            return random.random() < EDGE_CASE_CONFIG['major_issues_pct']
        return False
    
    def should_orphan_record(self):
        """Determine if record should have invalid foreign key."""
        if self.strict_mode or self.clean_mode:
            return False
        return random.random() < EDGE_CASE_CONFIG['orphaned_records_pct']
    
    def generate_categories(self, count):
        """Generate product categories with hierarchy."""
        print(f"\nGenerating categories...")
        categories = []
        
        # Root categories (no parent)
        root_categories = [
            'Electronics', 'Clothing', 'Home & Garden', 'Sports & Outdoors',
            'Books & Media', 'Toys & Games', 'Health & Beauty', 'Automotive',
            'Food & Beverage', 'Office Supplies'
        ]
        
        for i in tqdm(range(count)):
            category_id = f"CAT-{str(uuid.uuid4())[:8]}"
            self.category_ids.append(category_id)
            
            # First 10 are root categories
            if i < len(root_categories):
                name = root_categories[i]
                parent_id = None
            else:
                # Subcategories
                name = f"{self.fake.word().capitalize()} {random.choice(['Pro', 'Plus', 'Premium', 'Basic', 'Standard'])}"
                
                # Edge case: orphaned parent reference
                if not self.strict_mode and self.should_add_edge_case('major'):
                    parent_id = f"CAT-{str(uuid.uuid4())[:8]}"  # Non-existent parent
                else:
                    # Reference an existing category (only from already created ones)
                    parent_id = random.choice(self.category_ids[:i]) if self.category_ids[:i] else None
            
            categories.append({
                'category_id': category_id,
                'name': name,
                'parent_category_id': parent_id or ''
            })
        
        return categories
    
    def generate_customers(self, count):
        """Generate customer records with edge cases."""
        print(f"\nGenerating {count} customers...")
        customers = []
        
        for _ in tqdm(range(count)):
            customer_id = f"CUST-{str(uuid.uuid4())[:8]}"
            self.customer_ids.append(customer_id)
            
            # Base data
            email = self.fake.email()
            name = self.fake.name()
            country = self.fake.country_code()
            phone = self.fake.phone_number()
            created_at = self.random_date()
            status = random.choice(['active', 'active', 'active', 'inactive', 'suspended'])
            
            # Apply edge cases
            if self.should_add_edge_case('minor'):
                # Minor issues: missing phone, empty fields
                if random.random() < 0.5:
                    phone = ''
                else:
                    name = ''
            
            if self.should_add_edge_case('major'):
                # Major issues: invalid email, malformed data
                edge_case_type = random.choice(['invalid_email', 'special_chars', 'duplicate'])
                
                if edge_case_type == 'invalid_email':
                    email = random.choice([
                        'not-an-email',
                        'missing@domain',
                        '@nodomain.com',
                        'spaces in email@test.com',
                        ''
                    ])
                elif edge_case_type == 'special_chars':
                    name = f"{name} 🎉 <script>alert('xss')</script>"
                # duplicate will naturally occur with faker
            
            customers.append({
                'customer_id': customer_id,
                'email': email,
                'name': name,
                'country': country,
                'phone': phone,
                'created_at': created_at.isoformat(),
                'status': status
            })
        
        return customers
    
    def generate_products(self, count):
        """Generate product catalog."""
        print(f"\nGenerating {count} products...")
        products = []
        
        for _ in tqdm(range(count)):
            product_id = f"PROD-{str(uuid.uuid4())[:8]}"
            self.product_ids.append(product_id)
            
            name = f"{self.fake.word().capitalize()} {random.choice(['Pro', 'Max', 'Ultra', 'Mini', 'Standard'])}"
            # Generate description and remove newlines to prevent CSV issues
            if random.random() > 0.1:
                description = self.fake.text(max_nb_chars=200).replace('\n', ' ').replace('\r', ' ')
            else:
                description = ''
            price = round(random.uniform(5.99, 999.99), 2)
            created_at = self.random_date()
            
            # Category assignment
            if self.category_ids and not self.should_orphan_record():
                category_id = random.choice(self.category_ids)
            else:
                category_id = f"CAT-{str(uuid.uuid4())[:8]}"  # Orphaned or if no categories yet
            
            # Edge cases
            if self.should_add_edge_case('minor'):
                if random.random() < 0.3:
                    description = ''
                elif random.random() < 0.3:
                    name = name[:50] + "..." # Truncated name
            
            if self.should_add_edge_case('major'):
                edge_type = random.choice(['negative_price', 'zero_price', 'extreme_price'])
                if edge_type == 'negative_price':
                    price = -abs(price)
                elif edge_type == 'zero_price':
                    price = 0.0
                elif edge_type == 'extreme_price':
                    price = 999999.99
            
            products.append({
                'product_id': product_id,
                'name': name,
                'category_id': category_id,
                'price': f"{price:.2f}",
                'description': description,
                'created_at': created_at.isoformat()
            })
        
        return products
    
    def generate_orders(self, count):
        """Generate order transactions."""
        print(f"\nGenerating {count} orders...")
        orders = []
        
        for _ in tqdm(range(count)):
            order_id = f"ORD-{str(uuid.uuid4())[:8]}"
            self.order_ids.append(order_id)
            
            # Customer reference
            if self.customer_ids and not self.should_orphan_record():
                customer_id = random.choice(self.customer_ids)
            else:
                customer_id = f"CUST-{str(uuid.uuid4())[:8]}"  # Orphaned
            
            order_date = self.random_date()
            # Status distribution: ~60% completed, ~30% in-progress, ~10% cancelled/refunded
            status = random.choices(
                ['completed', 'pending', 'processing', 'shipped', 'cancelled', 'refunded'],
                weights=[60, 15, 10, 5, 7, 3]
            )[0]
            
            total_amount = round(random.uniform(10.0, 5000.0), 2)
            shipping_address = self.fake.address().replace('\n', ', ')
            
            # Edge cases
            if self.should_add_edge_case('minor'):
                if random.random() < 0.3:
                    shipping_address = ''
            
            if self.should_add_edge_case('major'):
                edge_type = random.choice(['zero_total', 'negative_total', 'no_address'])
                if edge_type == 'zero_total':
                    total_amount = 0.0
                elif edge_type == 'negative_total':
                    total_amount = -abs(total_amount)
                elif edge_type == 'no_address':
                    shipping_address = ''
            
            orders.append({
                'order_id': order_id,
                'customer_id': customer_id,
                'order_date': order_date.isoformat(),
                'status': status,
                'total_amount': f"{total_amount:.2f}",
                'shipping_address': shipping_address
            })
        
        return orders
    
    def generate_order_items(self, orders_count):
        """Generate order line items (1-5 items per order)."""
        # Calculate approximate count
        avg_items_per_order = 3
        estimated_count = orders_count * avg_items_per_order
        
        print(f"\nGenerating order items (~{estimated_count} items)...")
        order_items = []
        
        for order_id in tqdm(self.order_ids):
            # Random items per order (1-5, weighted towards 2-3)
            num_items = random.choices([1, 2, 3, 4, 5], weights=[10, 30, 35, 20, 5])[0]
            
            for _ in range(num_items):
                order_item_id = f"ITEM-{str(uuid.uuid4())[:8]}"
                
                # Product reference
                if self.product_ids and not self.should_orphan_record():
                    product_id = random.choice(self.product_ids)
                else:
                    product_id = f"PROD-{str(uuid.uuid4())[:8]}"  # Orphaned
                
                quantity = random.randint(1, 10)
                unit_price = round(random.uniform(5.99, 999.99), 2)
                
                # Edge cases
                if self.should_add_edge_case('major'):
                    edge_type = random.choice(['zero_quantity', 'negative_price', 'extreme_quantity'])
                    if edge_type == 'zero_quantity':
                        quantity = 0
                    elif edge_type == 'negative_price':
                        unit_price = -abs(unit_price)
                    elif edge_type == 'extreme_quantity':
                        quantity = 99999
                
                order_items.append({
                    'order_item_id': order_item_id,
                    'order_id': order_id,
                    'product_id': product_id,
                    'quantity': quantity,
                    'unit_price': f"{unit_price:.2f}"
                })
        
        return order_items
    
    def generate_inventory(self, products_count, warehouses_count):
        """Generate inventory levels across warehouses."""
        # Create warehouse IDs
        self.warehouse_ids = [f"WH-{i+1:03d}" for i in range(warehouses_count)]
        
        # Each product appears in 1-3 warehouses
        estimated_count = products_count * 2
        print(f"\nGenerating inventory records (~{estimated_count} records)...")
        
        inventory = []
        
        for product_id in tqdm(self.product_ids):
            # Random warehouses per product
            num_warehouses = random.choices([1, 2, 3], weights=[50, 35, 15])[0]
            selected_warehouses = random.sample(self.warehouse_ids, num_warehouses)
            
            for warehouse_id in selected_warehouses:
                quantity = random.randint(0, 1000)
                last_updated = self.random_date()
                
                # Edge cases
                if self.should_add_edge_case('minor'):
                    if random.random() < 0.2:
                        quantity = 0  # Out of stock
                
                if self.should_add_edge_case('major'):
                    if random.random() < 0.5:
                        quantity = -random.randint(1, 100)  # Negative stock (data error)
                
                inventory.append({
                    'product_id': product_id,
                    'warehouse_id': warehouse_id,
                    'quantity': quantity,
                    'last_updated': last_updated.isoformat()
                })
        
        return inventory
    
    def write_csv(self, filename, data, fieldnames):
        """Write data to CSV file with validation."""
        filepath = Path(__file__).parent / filename
        
        # Write CSV
        with open(filepath, 'w', newline='', encoding='utf-8') as f:
            writer = csv.DictWriter(f, fieldnames=fieldnames, quoting=csv.QUOTE_MINIMAL)
            writer.writeheader()
            writer.writerows(data)
        
        # Validate: count lines vs rows
        with open(filepath, 'r') as f:
            file_lines = sum(1 for _ in f)
        expected_lines = len(data) + 1  # data rows + header
        
        if file_lines != expected_lines:
            print(f"⚠ WARNING: {filename} has {file_lines} lines but expected {expected_lines}")
            print(f"  This may indicate unescaped newlines in data fields")
        else:
            print(f"✓ Wrote {len(data)} rows to {filename}")


def parse_date(date_string):
    """Parse date string in YYYY-MM-DD format."""
    try:
        return datetime.strptime(date_string, '%Y-%m-%d')
    except ValueError:
        raise argparse.ArgumentTypeError(f"Invalid date format: {date_string}. Use YYYY-MM-DD")


def main():
    parser = argparse.ArgumentParser(
        description='Generate realistic e-commerce dataset with configurable edge cases',
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Generate with default sizes
  python generate_data.py
  
  # Generate custom sizes
  python generate_data.py --customers 100000 --products 20000 --orders 500000
  
  # Generate clean data with perfect referential integrity
  python generate_data.py --clean --strict
  
  # Generate data for specific date range
  python generate_data.py --start-date 2020-01-01 --end-date 2023-12-31
  
  # Small dataset for testing
  python generate_data.py --customers 1000 --products 500 --orders 5000
        """
    )
    
    # Row count arguments
    parser.add_argument('--customers', type=int, default=DEFAULT_CUSTOMERS,
                        help=f'Number of customers (default: {DEFAULT_CUSTOMERS})')
    parser.add_argument('--products', type=int, default=DEFAULT_PRODUCTS,
                        help=f'Number of products (default: {DEFAULT_PRODUCTS})')
    parser.add_argument('--orders', type=int, default=DEFAULT_ORDERS,
                        help=f'Number of orders (default: {DEFAULT_ORDERS})')
    parser.add_argument('--categories', type=int, default=DEFAULT_CATEGORIES,
                        help=f'Number of categories (default: {DEFAULT_CATEGORIES})')
    parser.add_argument('--warehouses', type=int, default=DEFAULT_WAREHOUSES,
                        help=f'Number of warehouses (default: {DEFAULT_WAREHOUSES})')
    
    # Data quality flags
    parser.add_argument('--clean', action='store_true',
                        help='Generate 100%% clean data without edge cases')
    parser.add_argument('--strict', action='store_true',
                        help='Maintain perfect referential integrity (no orphaned records)')
    
    # Date range
    parser.add_argument('--start-date', type=parse_date,
                        help='Start date for timestamps (YYYY-MM-DD, default: 2023-01-01)')
    parser.add_argument('--end-date', type=parse_date,
                        help='End date for timestamps (YYYY-MM-DD, default: today)')
    
    # Other options
    parser.add_argument('--seed', type=int,
                        help='Random seed for reproducible generation')
    
    args = parser.parse_args()
    
    # Initialize generator
    generator = EcommerceDataGenerator(
        clean_mode=args.clean,
        strict_mode=args.strict,
        start_date=args.start_date,
        end_date=args.end_date,
        seed=args.seed
    )
    
    print("=" * 60)
    print("E-COMMERCE DATA GENERATOR")
    print("=" * 60)
    print(f"Configuration:")
    print(f"  Customers: {args.customers:,}")
    print(f"  Products: {args.products:,}")
    print(f"  Orders: {args.orders:,}")
    print(f"  Categories: {args.categories:,}")
    print(f"  Warehouses: {args.warehouses:,}")
    print(f"  Clean mode: {args.clean}")
    print(f"  Strict mode: {args.strict}")
    print(f"  Date range: {generator.start_date.date()} to {generator.end_date.date()}")
    if args.seed:
        print(f"  Random seed: {args.seed}")
    print("=" * 60)
    
    # Generate all tables
    categories = generator.generate_categories(args.categories)
    customers = generator.generate_customers(args.customers)
    products = generator.generate_products(args.products)
    orders = generator.generate_orders(args.orders)
    order_items = generator.generate_order_items(args.orders)
    inventory = generator.generate_inventory(args.products, args.warehouses)
    
    # Write to CSV files
    print("\nWriting CSV files...")
    print("-" * 60)
    
    generator.write_csv('categories.csv', categories, 
                        ['category_id', 'name', 'parent_category_id'])
    
    generator.write_csv('customers.csv', customers,
                        ['customer_id', 'email', 'name', 'country', 'phone', 'created_at', 'status'])
    
    generator.write_csv('products.csv', products,
                        ['product_id', 'name', 'category_id', 'price', 'description', 'created_at'])
    
    generator.write_csv('orders.csv', orders,
                        ['order_id', 'customer_id', 'order_date', 'status', 'total_amount', 'shipping_address'])
    
    generator.write_csv('order_items.csv', order_items,
                        ['order_item_id', 'order_id', 'product_id', 'quantity', 'unit_price'])
    
    generator.write_csv('inventory.csv', inventory,
                        ['product_id', 'warehouse_id', 'quantity', 'last_updated'])
    
    # Print summary statistics
    print("=" * 60)
    print("GENERATION SUMMARY")
    print("=" * 60)
    print(f"✓ Categories: {len(categories):,} rows")
    print(f"✓ Customers: {len(customers):,} rows")
    print(f"✓ Products: {len(products):,} rows")
    print(f"✓ Orders: {len(orders):,} rows")
    print(f"✓ Order Items: {len(order_items):,} rows ({len(order_items)/len(orders):.2f} avg per order)")
    print(f"✓ Inventory: {len(inventory):,} rows ({len(inventory)/len(products):.2f} avg per product)")
    
    if not args.clean:
        print("\n" + "=" * 60)
        print("EDGE CASES INCLUDED")
        print("=" * 60)
        
        # Count edge cases
        invalid_emails = sum(1 for c in customers if '@' not in c['email'] or not c['email'])
        missing_desc = sum(1 for p in products if not p['description'])
        negative_prices = sum(1 for p in products if float(p['price']) < 0)
        cancelled_orders = sum(1 for o in orders if o['status'] in ['cancelled', 'refunded'])
        
        print(f"Customers with invalid emails: {invalid_emails} ({invalid_emails/len(customers)*100:.1f}%)")
        print(f"Products with missing descriptions: {missing_desc} ({missing_desc/len(products)*100:.1f}%)")
        print(f"Products with negative/zero prices: {negative_prices} ({negative_prices/len(products)*100:.1f}%)")
        print(f"Cancelled/refunded orders: {cancelled_orders} ({cancelled_orders/len(orders)*100:.1f}%)")
        
        if not args.strict:
            # Check referential integrity
            cat_ids = set(c['category_id'] for c in categories)
            orphaned_products = sum(1 for p in products if p['category_id'] not in cat_ids)
            
            cust_ids = set(c['customer_id'] for c in customers)
            orphaned_orders = sum(1 for o in orders if o['customer_id'] not in cust_ids)
            
            print(f"\nReferential integrity issues:")
            print(f"  Products with invalid category_id: {orphaned_products} ({orphaned_products/len(products)*100:.1f}%)")
            print(f"  Orders with invalid customer_id: {orphaned_orders} ({orphaned_orders/len(orders)*100:.1f}%)")
    
    print("\n" + "=" * 60)
    print("✓ Dataset generation complete!")
    print("=" * 60)


if __name__ == '__main__':
    main()
