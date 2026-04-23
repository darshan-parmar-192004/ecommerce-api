package constants

import "time"

const (
	DefaultPort  = 8080
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100

	ResponseStatusOK   = "ok"
	ResponseStatusUp   = "up"
	ResponseStatusDown = "down"

	ResponseMessageDeleted = "Deleted"

	JSONFieldData       = "data"
	JSONFieldPagination = "pagination"
	JSONFieldPage       = "page"
	JSONFieldLimit      = "limit"
	JSONFieldTotalItems = "total_items"
	JSONFieldTotalPages = "total_pages"
	JSONFieldMessage    = "message"
	JSONFieldStatus     = "status"
	JSONFieldError      = "error"
	JSONFieldSuccess    = "success"
	JSONFieldDetails    = "details"
	JSONFieldCode       = "code"

	MsgPagePositive     = "page must be positive integer"
	MsgLimitPositive    = "limit must be positive integer"
	MsgMinPriceValid    = "min_price must be valid number"
	MsgMaxPriceValid    = "max_price must be valid number"
	MsgMinMaxPrice      = "min_price cannot be empty than max_price"
	MsgPageExceeds      = "page exceeds total page"
	MsgMalformedJSON    = "Malformed JSON request body"
	MsgInvalidJSON      = "Invalid JSON format/malformed JSON"
	MsgProductNotFound  = "Product with ID %s does not exists"
	MsgProductNotFound2 = "Product with ID %s not found"
	MsgValidationEmpty  = "any fields should not be empty in order to create product"
	MsgValidationUpdate = "all fields must be filled in order to update the product"
	MsgPersistFailed    = "Failed to persist product"
	MsgStorageFailed    = "Failed to update storage"

	HealthDBUp            = "It's healthy"
	HealthDBHeavyLoad     = "The database is experiencing heavy load."
	HealthDBHighWaits     = "The database has a high number of wait events, indicating potential bottlenecks."
	HealthDBIdleClosing   = "Many idle connections are being closed, consider revising the connection pool settings."
	HealthDBLifetimeClose = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."

	ProductIDPrefix = "PROD-"
	ProductIDFormat = "%08d"

	OrderIDPrefix = "ORD-"

	CSVDockerPath = "/app/datasets/ecommerce/"
	CSVProducts   = "products.csv"
	CSVCategories = "categories.csv"
	CSVCustomers  = "customers.csv"
	CSVInventory  = "inventory.csv"
	CSVOrders     = "orders.csv"
	CSVOrderItems = "order_items.csv"

	ErrInternalServer = "internal_server_error"
	MsgSomethingWrong = "Something went wrong"

	DBDriverPostgres = "postgres"
	DBDriverMysql    = "mysql"
	DBDriverSQLite   = "sqlite3"

	DBMaxOpenConns    = 20
	DBMaxIdleConns    = 5
	DBConnMaxLifetime = 30

	JWTExpiration = 24 * time.Hour
	JWTSecret     = "your-secret-key"

	ContextTimeoutShort  = 1 * time.Second
	ContextTimeoutMedium = 5 * time.Second
	ContextTimeoutLong   = 30 * time.Second

	MsgDeletedSuccessfully   = "Deleted successfully"
	MsgLoggedOutSuccessfully = "Logged out successfully"

	MsgInvalidRequestBody    = "Invalid request body"
	MsgMissingAuthHeader     = "Authorization token required"
	MsgInvalidAuthHeader     = "Invalid authorization header format"
	MsgInvalidOrExpiredToken = "Invalid or expired token"
	MsgSessionExpired        = "Session expired or invalid"

	CustomerIDPrefix     = "CUST-"
	CustomerStatusActive = "active"

	RouteProducts       = "/products"
	RouteProductsID     = "/products/:id"
	RouteAuthRegister   = "/auth/register"
	RouteAuthLogin      = "/auth/login"
	RouteAuthLogout     = "/auth/logout"
	RouteCustomersMe    = "/customers/me"
	RouteHealth         = "/health"
	RouteCategories     = "/categories"
	RouteCategoriesID   = "/categories/:id"
	RouteCategoriesProd = "/categories/:id/products"
	RouteCategoriesHier = "/categories/hierarchy"
	RouteCustomers      = "/customers"
	RouteCustomersID    = "/customers/:id"
	RouteCustomersOrd   = "/customers/:id/orders"
	RouteCustomersLTV   = "/customers/:id/lifetime-value"
	RouteOrders         = "/orders"
	RouteOrdersID       = "/orders/:id"
	RouteOrdersItems    = "/orders/:id/items"
	RouteInventory      = "/inventory"
	RouteInvStock       = "/inventory/stock"
	RouteInvCLV         = "/inventory/customer-lifetime-value"
	RouteInvHier        = "/inventory/hierarchy"
	RouteInvTopSell     = "/inventory/top-sellers"
	RouteStatsCache     = "/stats/cache"
)

var (
	CORSAllowOrigins     = []string{"*"}
	CORSAllowMethods     = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	CORSAllowHeaders     = []string{"Accept", "Authorization", "Content-Type"}
	CORSAllowCredentials = false
	CORSMaxAge           = 300
)

var (
	CacheProductsAllTTL   = 5 * time.Minute
	CacheProductByIDTTL   = 10 * time.Minute
	CacheCategoriesTTL    = 30 * time.Minute
	CacheKeyProductsAll   = "products:all"
	CacheKeyProductPrefix = "product:"
	CacheKeyCategoriesAll = "categories:all"
	CacheKeySessionPrefix = "session:"
	CacheSessionTTL       = 1 * time.Hour
)
