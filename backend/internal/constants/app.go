package constants

const (
	DefaultPort  = 8080
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100

	// Database pool health thresholds
	DBPoolMaxConnectionsThreshold = 40
	DBPoolWaitCountThreshold      = 1000

	// File permissions
	FilePermissionReadWrite = 0644

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

	CSVProductsPath   = "./datasets/ecommerce/products.csv"
	CSVCategoriesPath = "./datasets/ecommerce/categories.csv"
	CSVCustomersPath  = "./datasets/ecommerce/customers.csv"
	CSVInventoryPath  = "./datasets/ecommerce/inventory.csv"
	CSVOrdersPath     = "./datasets/ecommerce/orders.csv"
	CSVOrderItemsPath = "./datasets/ecommerce/order_items.csv"

	ErrInternalServer = "internal_server_error"
	MsgSomethingWrong = "Something went wrong"

	DBDriverPostgres = "postgres"
	DBDriverMysql    = "mysql"
	DBDriverSQLite   = "sqlite3"

	DBMaxOpenConns    = 20
	DBMaxIdleConns    = 5
	DBConnMaxLifetime = 30

	RouteParamID = ":id"

	RouteProducts       = "/products"
	RouteProductsID     = "/products/" + RouteParamID
	RouteHealth         = "/health"
	RouteCategories     = "/categories"
	RouteCategoriesID   = "/categories/" + RouteParamID
	RouteCategoriesProd = "/categories/" + RouteParamID + "/products"
	RouteCategoriesHier = "/categories/hierarchy"
	RouteCustomers      = "/customers"
	RouteCustomersID    = "/customers/" + RouteParamID
	RouteCustomersOrd   = "/customers/" + RouteParamID + "/orders"
	RouteCustomersLTV   = "/customers/" + RouteParamID + "/lifetime-value"
	RouteOrders         = "/orders"
	RouteOrdersID       = "/orders/" + RouteParamID
	RouteInventory      = "/inventory"
	RouteInvStock       = "/inventory/stock"
	RouteInvCLV         = "/inventory/customer-lifetime-value"
	RouteInvHier        = "/inventory/hierarchy"
	RouteInvTopSell     = "/inventory/top-sellers"
)

var (
	CORSAllowOrigins     = []string{"*"}
	CORSAllowMethods     = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	CORSAllowHeaders     = []string{"Accept", "Authorization", "Content-Type"}
	CORSAllowCredentials = false
	CORSMaxAge           = 300
)
