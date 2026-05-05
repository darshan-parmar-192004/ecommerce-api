package config

import (
	"backend/internal/constants"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Port          int    `envconfig:"BLUEPRINT_PORT" default:"8080"`
	DBHost        string `envconfig:"BLUEPRINT_DB_HOST"`
	DBPort        string `envconfig:"BLUEPRINT_DB_PORT" default:"5432"`
	DBName        string `envconfig:"BLUEPRINT_DB_DATABASE"`
	DBUser        string `envconfig:"BLUEPRINT_DB_USERNAME"`
	DBPass        string `envconfig:"BLUEPRINT_DB_PASSWORD"`
	DBSchema      string `envconfig:"BLUEPRINT_DB_SCHEMA" default:"public"`
	DBDialect     string `envconfig:"BLUEPRINT_DB_DIALECT" default:"pgx"`
	MigrationPath string `envconfig:"BLUEPRINT_MIGRATION_PATH" default:"internal/migrations"`
	DatasetPath   string `envconfig:"BLUEPRINT_DATASET_PATH" default:"internal/datasets/ecommerce"`
	CategoriesCSV string `envconfig:"BLUEPRINT_CSV_CATEGORIES" `
	CustomersCSV  string `envconfig:"BLUEPRINT_CSV_CUSTOMERS" `
	ProductsCSV   string `envconfig:"BLUEPRINT_CSV_PRODUCTS" `
	OrdersCSV     string `envconfig:"BLUEPRINT_CSV_ORDERS" `
	InventoryCSV  string `envconfig:"BLUEPRINT_CSV_INVENTORY" `
	OrderItemsCSV string `envconfig:"BLUEPRINT_CSV_ORDER_ITEMS" `
}

var cfg *AppConfig

func Load() (*AppConfig, error) {
	if cfg != nil {
		return cfg, nil
	}

	_ = godotenv.Load()

	cfg = &AppConfig{}

	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if cfg.Port == 0 {
		cfg.Port = constants.DefaultPort
	}

	if cfg.DBSchema == "" {
		cfg.DBSchema = "public"
	}
	if cfg.DBDialect == "" {
		cfg.DBDialect = "pgx"
	}

	return cfg, nil
}

func LoadTest(dbHost, dbPort, dbName, dbUser, dbPwd string) (*AppConfig, error) {
	return &AppConfig{
		Port:      constants.DefaultPort,
		DBHost:    dbHost,
		DBPort:    dbPort,
		DBName:    dbName,
		DBUser:    dbUser,
		DBPass:    dbPwd,
		DBSchema:  "public",
		DBDialect: "postgres",
	}, nil
}

func (c *AppConfig) GetDSN() string {
	return "postgres://" + c.DBUser + ":" + c.DBPass + "@" + c.DBHost + ":" + c.DBPort + "/" + c.DBName + "?sslmode=disable&search_path=" + c.DBSchema
}

func GetPort() int {
	cfg, err := Load()
	if err != nil {
		return constants.DefaultPort
	}
	if cfg.Port > 0 {
		return cfg.Port
	}
	return constants.DefaultPort
}

func GetDBDialect() string {
	if cfg != nil && cfg.DBDialect != "" {
		return cfg.DBDialect
	}
	return "pgx"
}

func (c *AppConfig) GetCSVPath(csvType string) string {
	switch csvType {
	case "categories":
		if c.CategoriesCSV != "" {
			return c.CategoriesCSV
		}
		return c.DatasetPath + "/categories.csv"
	case "customers":
		if c.CustomersCSV != "" {
			return c.CustomersCSV
		}
		return c.DatasetPath + "/customers.csv"
	case "products":
		if c.ProductsCSV != "" {
			return c.ProductsCSV
		}
		return c.DatasetPath + "/products.csv"
	case "orders":
		if c.OrdersCSV != "" {
			return c.OrdersCSV
		}
		return c.DatasetPath + "/orders.csv"
	case "inventory":
		if c.InventoryCSV != "" {
			return c.InventoryCSV
		}
		return c.DatasetPath + "/inventory.csv"
	case "order_items":
		if c.OrderItemsCSV != "" {
			return c.OrderItemsCSV
		}
		return c.DatasetPath + "/order_items.csv"
	default:
		return c.DatasetPath + "/" + csvType + ".csv"
	}
}
