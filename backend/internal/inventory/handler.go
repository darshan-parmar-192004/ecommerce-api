package inventory

import (
	"backend/internal/database"
	"backend/internal/errors"
	"backend/internal/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	db database.Service
}

func NewHandler(db database.Service) *Handler {
	return &Handler{db: db}
}

// GetAll returns a simple list of all items in the inventory table.
func (h *Handler) GetAll(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	query := `SELECT product_id, warehouse_id, quantity, last_updated FROM inventory`

	rows, err := h.db.DB().QueryContext(ctx, query)
	if err != nil {
		return errors.SendError(c, 500, errors.ErrDatabase, "Failed to retrieve inventory list", nil)
	}
	defer rows.Close()

	var inventoryList []models.Inventory
	for rows.Next() {
		var i models.Inventory
		if err := rows.Scan(&i.ProductID, &i.WarehouseID, &i.Quantity, &i.LastUpdated); err != nil {
			return errors.SendError(c, 500, errors.ErrDatabase, "Data scanning error", nil)
		}
		inventoryList = append(inventoryList, i)
	}

	return c.JSON(inventoryList)
}

// Inventory across warehouses (JOIN inventory with products)
func (h *Handler) GetStockLevels(c fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	query := `
		SELECT p.name, i.product_id, i.warehouse_id, i.quantity, i.last_updated
		FROM inventory i
		JOIN products p ON i.product_id = p.product_id
		ORDER BY i.quantity ASC`

	rows, err := h.db.DB().QueryContext(ctx, query)
	if err != nil {
		return errors.SendError(c, 500, errors.ErrDatabase, "Failed to fetch stock", nil)
	}
	defer rows.Close()

	type StockInfo struct {
		ProductName string `json:"product_name"`
		models.Inventory
	}

	var results []StockInfo
	for rows.Next() {
		var s StockInfo
		if err := rows.Scan(&s.ProductName, &s.ProductID, &s.WarehouseID, &s.Quantity, &s.LastUpdated); err != nil {
			return err
		}
		results = append(results, s)
	}
	return c.JSON(results)
}

// Customer lifetime value (Aggregate orders per customer)
func (h *Handler) GetCustomerCLV(c fiber.Ctx) error {
	query := `
		SELECT customer_id, COUNT(order_id) as order_count, SUM(total_amount) as total_spent
		FROM orders
		GROUP BY customer_id
		ORDER BY total_spent DESC`

	rows, err := h.db.DB().Query(query)
	if err != nil {
		return errors.SendError(c, 500, errors.ErrDatabase, "Failed to calculate CLV", nil)
	}
	defer rows.Close()

	var stats []fiber.Map
	for rows.Next() {
		var id string
		var count int
		var total float64
		err := rows.Scan(&id, &count, &total)
		if err != nil {
			return err
		}
		stats = append(stats, fiber.Map{"customer_id": id, "order_count": count, "lifetime_value": total})
	}
	return c.JSON(stats)
}

// Category hierarchy (Recursive CTE for parent-child categories)
func (h *Handler) GetCategoryTree(c fiber.Ctx) error {

	query := `
	WITH RECURSIVE category_path AS (
    SELECT category_id, name, parent_category_id, name AS path
    FROM categories
    WHERE parent_category_id IS NULL

    UNION ALL

    SELECT c.category_id, c.name, c.parent_category_id,
           cp.path || ' > ' || c.name
    FROM categories c
    JOIN category_path cp
        ON cp.category_id = c.parent_category_id
        )
        SELECT category_id, name, parent_category_id, path
        FROM category_path
        ORDER BY path;
	`

	rows, err := h.db.DB().Query(query)
	if err != nil {
		return errors.SendError(
			c,
			fiber.StatusInternalServerError,
			errors.ErrDatabase,
			"Failed to build category tree",
			fiber.Map{
				"error": err.Error(),
			},
		)
	}
	defer rows.Close()

	var tree []fiber.Map

	for rows.Next() {

		var id, name, path string
		var parentID *string

		err := rows.Scan(&id, &name, &parentID, &path)
		if err != nil {
			return errors.SendError(
				c,
				fiber.StatusInternalServerError,
				errors.ErrDatabase,
				"Failed to scan category tree",
				fiber.Map{
					"error": err.Error(),
				},
			)
		}

		tree = append(tree, fiber.Map{
			"id":        id,
			"name":      name,
			"parent_id": parentID,
			"full_path": path,
		})
	}

	return c.JSON(fiber.Map{
		"data": tree,
	})
}

// Top selling products (Aggregate order_items)
func (h *Handler) GetTopSellers(c fiber.Ctx) error {
	query := `
		SELECT p.name, SUM(oi.quantity) as total_sold
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		GROUP BY p.product_id, p.name
		ORDER BY total_sold DESC
		LIMIT 10`

	rows, err := h.db.DB().Query(query)
	if err != nil {
		return errors.SendError(c, 500, errors.ErrDatabase, "Failed to fetch top sellers", nil)
	}
	defer rows.Close()

	var topProducts []fiber.Map
	for rows.Next() {
		var name string
		var total int
		err := rows.Scan(&name, &total)
		if err != nil {
			return err
		}
		topProducts = append(topProducts, fiber.Map{"product": name, "units_sold": total})
	}
	return c.JSON(topProducts)
}
