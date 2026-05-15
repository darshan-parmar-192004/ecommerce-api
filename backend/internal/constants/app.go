package constants

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
	JSONFieldDetails     = "details"
	JSONFieldCode        = "code"
	JSONFieldCustomerID  = "customer_id"
	JSONFieldTotalOrders = "total_orders"
	JSONFieldLifetimeValue = "lifetime_value"
	JSONFieldOrderID     = "order_id"
	JSONFieldTotalAmount = "total_amount"
	JSONFieldDebug       = "debug"
	JSONFieldProductName = "product_name"
	JSONFieldWarehouseID = "warehouse_id"
	JSONFieldQuantity    = "quantity"
	JSONFieldUpdatedAt   = "updated_at"
	JSONFieldProductID   = "product_id"
	JSONFieldProduct     = "product"
	JSONFieldOrderCount  = "order_count"
	JSONFieldUnitsSold   = "units_sold"

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
	MsgValidationFailed = "Validation failed"
	MsgProductIDExists  = "Product ID already exists"
	MsgFailedToFetch    = "Failed to fetch"
	MsgFailedToCreate   = "Failed to create"
	MsgFailedToUpdate   = "Failed to update"
	MsgFailedToDelete   = "Failed to delete"
	MsgCheckProduct     = "Failed to check product"
	MsgInternalError    = "Internal server error"

	HealthDBUp            = "It's healthy"
	HealthDBHeavyLoad     = "The database is experiencing heavy load."
	HealthDBHighWaits     = "The database has a high number of wait events, indicating potential bottlenecks."
	HealthDBIdleClosing   = "Many idle connections are being closed, consider revising the connection pool settings."
	HealthDBLifetimeClose = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."

	MaxProductNameLength = 200
	MaxProductDescLength = 500
	ErrCodeDuplicateKey  = "23505"

	ErrInternalServer = "internal_server_error"
	MsgSomethingWrong = "Something went wrong"

	DBDriverPostgres = "postgres"
	DBDriverMysql    = "mysql"
	DBDriverSQLite   = "sqlite3"

	DBMaxOpenConns    = 20
	DBMaxIdleConns    = 5
	DBConnMaxLifetime = 30

	DBDialectPgx      = "pgx"
	DBSchemaDefault   = "public"
	DBTimeoutSec      = 5
	ShutdownTimeoutSec = 5

	RouteProducts       = "/products"
	RouteProductsID     = "/products/:id"
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
	RouteInventory      = "/inventory"
	RouteInvStock       = "/inventory/stock"
	RouteInvCLV         = "/inventory/customer-lifetime-value"
	RouteInvHier        = "/inventory/hierarchy"
	RouteInvTopSell     = "/inventory/top-sellers"
	RouteCacheStats     = "/stats/cache"

	CacheKeyProductList   = "products:%s:%s:%s:%s:%d:%d"
	CacheKeyProductByID   = "product:%s"
	CacheKeyCategoriesAll = "categories:all"
	CacheKeyCategoriesHier = "categories:hierarchy"

	ServerHeader = "backend"
	AppName      = "backend"

	HeaderXRequestID = "X-Request-Id"
	LocalsRequestID  = "request_id"

	LoggerEncoding       = "json"
	LoggerOutputStdout   = "stdout"
	LoggerOutputStderr   = "stderr"
	LoggerKeyTime        = "time"
	LoggerKeyLevel       = "level"
	LoggerKeyLogger      = "logger"
	LoggerKeyCaller      = "caller"
	LoggerKeyMsg         = "msg"
	LoggerKeyStacktrace  = "stacktrace"

	LogDataTimestamp  = "timestamp"
	LogDataMethod     = "method"
	LogDataPath       = "path"
	LogDataDurationMs = "duration_ms"

	ParamID       = "id"
	QueryCategory = "category"
	QueryMinPrice = "min_price"
	QueryMaxPrice = "max_price"
	QuerySearch   = "search"
	QueryPage     = "page"
	QueryLimit    = "limit"
)

var (
	CORSAllowOrigins     = []string{"*"}
	CORSAllowMethods     = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	CORSAllowHeaders     = []string{"Accept", "Authorization", "Content-Type"}
	CORSAllowCredentials = false
	CORSMaxAge           = 300

	CacheProductsListTTL = 5 * 60
	CacheProductByIDTTL  = 10 * 60
	CacheCategoriesTTL   = 30 * 60
)
