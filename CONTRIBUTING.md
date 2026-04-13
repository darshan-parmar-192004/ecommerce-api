# Contributing to Ecommerce API

Thank you for your interest in contributing! This document outlines the process for contributing to this project.

## Code of Conduct

Please be respectful and professional. We expect all contributors to follow our [Code of Conduct](CODE_OF_CONDUCT.md) (if present).

## How to Contribute

### 1. Fork the Repository

```bash
# Click "Fork" on GitHub, then clone your fork
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
```

### 2. Create a Feature Branch

```bash
# Create a new branch for your feature or fix
git checkout -b feature/your-feature-name

# Or for bug fixes
git checkout -b fix/issue-description
```

### 3. Make Your Changes

Follow the code style guides below and make your changes.

### 4. Write Tests

- Write unit tests for new functionality
- Ensure existing tests still pass
- Run tests locally before submitting

### 5. Commit Your Changes

```bash
# Stage your changes
git add .

# Commit with a descriptive message
git commit -m "Add feature: brief description of changes"
```

Follow the commit message conventions below.

### 6. Push and Create PR

```bash
# Push your branch
git push origin feature/your-feature-name

# Create Pull Request on GitHub
```

## Commit Message Conventions

Use clear, descriptive commit messages:

```
type(scope): description

[optional body]

[optional footer]
```

**Types:**
- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation
- `style` - Code style (formatting, no logic)
- `refactor` - Code refactoring
- `test` - Testing
- `chore` - Maintenance

**Examples:**
```
feat(products): add product search by name

Added search functionality to filter products by name
using case-insensitive matching.

Closes #123
```

```
fix(auth): handle expired JWT tokens

Return 401 when JWT token is expired or invalid
instead of 500 error.
```

## Code Style Guide

### Go (Backend)

Follow standard Go conventions:

```go
// Package names: short, lowercase
package product

// Group imports: stdlib, external, internal
import (
    "fmt"
    "time"

    "github.com/gofiber/fiber/v3"

    "backend/internal/models"
)

// Type names: PascalCase
type ProductService struct{}

// Functions: camelCase
func (s *ProductService) GetAll() {}

// Interfaces: single method first, then alphabetically
type Reader interface {
    Read(ctx context.Context, id string) (*Model, error)
}

// Error handling: return error as last return value
func Get(id string) (*Model, error) {
    if id == "" {
        return nil, errors.New("id is required")
    }
    return &Model{}, nil
}

// Naming: be descriptive, avoid abbreviations
customerID, not custID
```

**Rules:**
- Run `gofmt` before commits
- Run `go vet` to catch issues
- Maximum line length: 120 characters
- Add comments for exported functions
- Use `const` for magic numbers
- Handle errors, don't ignore them

### React/TypeScript (Frontend)

```tsx
// Component naming: PascalCase
function ProductCard({ product }) {
  // Hooks first
  const [isLoading, setIsLoading] = useState(false);
  
  // Helper functions after hooks
  const handleClick = () => {};
  
  // Render last
  return <div>{product.name}</div>;
}

// Props: TypeScript interfaces
interface ProductCardProps {
  product: Product;
  onSelect: (id: string) => void;
}

// Naming: camelCase for variables, PascalCase for components
const productList = [];
function ProductList() {}
```

## Testing Requirements

### Backend Tests

```bash
# Run all tests
cd backend
make test

# Run unit tests
go test ./internal/test/unit/... -v

# Run specific test
go test -run TestProduct ./... -v

# Run with coverage
go test -cover ./...
```

### Test Coverage

- New features should have >80% test coverage
- Bug fixes should include regression tests
- Integration tests for database operations

### Test Structure

```go
// File: product_service_test.go

func TestProductService_GetAll(t *testing.T) {
    // Arrange
    repo := mocks.NewProductRepository(t)
    service := NewProductService(repo)
    
    // Act
    products, err := service.GetAll()
    
    // Assert
    assert.NoError(t, err)
    assert.Len(t, products, 2)
}
```

## Documentation Requirements

- Update README.md if adding new features
- Add docstrings to new public functions
- Update OpenAPI spec for new endpoints
- Add ADR if making architectural decisions

## Pull Request Guidelines

### PR Title

Use conventional PR titles:
- `feat: add product search`
- `fix: resolve login redirect`
- `docs: update API docs`

### PR Description

Include:
1. **Summary**: What does this PR do?
2. **Type**: Feature, Bug Fix, Refactor, etc.
3. **Testing**: How did you test?
4. **Screenshots**: If UI changes (optional)

### Review Process

1. All PRs require review
2. Address feedback promptly
3. Request re-review after changes

## Development Workflow

```bash
# 1. Set up development environment
make setup  # or follow README instructions

# 2. Start services
podman-compose up -d

# 3. Run tests
make test

# 4. Run linter
go vet ./...

# 5. Format code
gofmt -w ./...
```

## Getting Help

- Open an issue for questions
- Join discussions in GitHub Discussions
- Check existing issues and PRs

## Recognition

Contributors will be recognized in:
- README.md contributors section
- Release notes

Thank you for contributing!