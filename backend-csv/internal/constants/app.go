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

	CSVProductsPath = "./internal/datasets/ecommerce/products.csv"

	ErrInternalServer = "internal_server_error"
	MsgSomethingWrong = "Something went wrong"

	RouteProducts   = "/products"
	RouteProductsID = "/products/:id"
	RouteHealth     = "/health"
)

var (
	CORSAllowOrigins     = []string{"*"}
	CORSAllowMethods     = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	CORSAllowHeaders     = []string{"Accept", "Authorization", "Content-Type"}
	CORSAllowCredentials = false
	CORSMaxAge           = 300
)
