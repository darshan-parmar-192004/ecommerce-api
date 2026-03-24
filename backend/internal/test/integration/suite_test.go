package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	product   *ProductIntegrationTest
	category  *CategoryIntegrationTest
	customer  *CustomerIntegrationTest
	order     *OrderIntegrationTest
	inventory *InventoryIntegrationTest
}

func (s *IntegrationSuite) SetupSuite() {
	s.product = &ProductIntegrationTest{}
	s.category = &CategoryIntegrationTest{}
	s.customer = &CustomerIntegrationTest{}
	s.order = &OrderIntegrationTest{}
	s.inventory = &InventoryIntegrationTest{}
}

func (s *IntegrationSuite) TestProductIntegration() {
	s.product.Run(s.T())
}

func (s *IntegrationSuite) TestCategoryIntegration() {
	s.category.Run(s.T())
}

func (s *IntegrationSuite) TestCustomerIntegration() {
	s.customer.Run(s.T())
}

func (s *IntegrationSuite) TestOrderIntegration() {
	s.order.Run(s.T())
}

func (s *IntegrationSuite) TestInventoryIntegration() {
	s.inventory.Run(s.T())
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}
