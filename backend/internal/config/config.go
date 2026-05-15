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
	CategoriesCSV string `envconfig:"BLUEPRINT_CSV_CATEGORIES" default:"internal/datasets/ecommerce/categories.csv"`
	CustomersCSV  string `envconfig:"BLUEPRINT_CSV_CUSTOMERS" default:"internal/datasets/ecommerce/customers.csv"`
	ProductsCSV   string `envconfig:"BLUEPRINT_CSV_PRODUCTS" default:"internal/datasets/ecommerce/products.csv"`
	OrdersCSV     string `envconfig:"BLUEPRINT_CSV_ORDERS" default:"internal/datasets/ecommerce/orders.csv"`
	InventoryCSV  string `envconfig:"BLUEPRINT_CSV_INVENTORY" default:"internal/datasets/ecommerce/inventory.csv"`
	OrderItemsCSV string `envconfig:"BLUEPRINT_CSV_ORDER_ITEMS" default:"internal/datasets/ecommerce/order_items.csv"`

	RedisHost     string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort     string `envconfig:"REDIS_PORT" default:"6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:""`
	RedisDB       int    `envconfig:"REDIS_DB" default:"0"`
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
		cfg.DBSchema = constants.DBSchemaDefault
	}
	if cfg.DBDialect == "" {
		cfg.DBDialect = constants.DBDialectPgx
	}

	return cfg, nil
}

func LoadTest(dbHost, dbPort, dbName, dbUser, dbPwd string) (*AppConfig, error) {
	return &AppConfig{
		Port:      constants.DefaultPort,
		DBHost:    dbHost,
		DBPort:    dbPort,
		DBName:    dbName,
		DBUser:    dbPwd,
		DBSchema:  constants.DBSchemaDefault,
		DBDialect: constants.DBDriverPostgres,
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
	return constants.DBDialectPgx
}