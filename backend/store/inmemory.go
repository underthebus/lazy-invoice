// In memory store implementation
package store

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/underthebus/lazy-invoice/backend/models"
)

var customerID uint64
var invoiceID uint64

type InMemoryStore struct {
	customers     []*models.Customer
	invoices      []*models.Invoice
	organizations []*models.Organization
	roles         []*models.Role
	users         []*models.User
}

var _ Storer = (*InMemoryStore)(nil)

func (s *InMemoryStore) GetInvoices(context.Context) ([]*models.Invoice, error) {
	return s.invoices, nil
}

func (s *InMemoryStore) CreateInvoice(ctx context.Context, invoiceInput models.NewInvoice) (*models.Invoice, error) {
	var customer *models.Customer
	for _, c := range s.customers {
		if c.ID == invoiceInput.CustomerID {
			customer = c
			break
		}
	}
	if customer == nil {
		return nil, errors.New("models.Customer not found.")
	}
	var organization *models.Organization
	for _, o := range s.organizations {
		if o.ID == invoiceInput.OrganizationID {
			organization = o
			break
		}
	}
	if organization == nil {
		return nil, errors.New("models.Organization not found.")
	}
	invoice := &models.Invoice{
		Date:       invoiceInput.Date,
		From:       organization,
		Identifier: invoiceInput.Identifier,
		To:         customer,
	}
	for _, input := range invoiceInput.Items {
		item := &models.InvoiceItem{
			ID:          fmt.Sprintf("%d", rand.Int()),
			Description: input.Description,
			UnitPrice:   input.UnitPrice,
			Quantity:    input.Quantity,
		}
		invoice.Items = append(invoice.Items, item)
	}
	invoiceID += 1
	invoice.ID = fmt.Sprintf("%d", invoiceID)
	s.invoices = append(s.invoices, invoice)
	return invoice, nil
}

func (s *InMemoryStore) EditInvoice(ctx context.Context, invoiceInput models.EditInvoice) (*models.Invoice, error) {
	var invoice *models.Invoice
	for _, i := range s.invoices {
		if i.ID == invoiceInput.ID {
			invoice = i
			break
		}
	}
	if invoice == nil {
		return nil, errors.New("models.Invoice not found.")
	}
	if invoiceInput.Date != nil {
		invoice.Date = *invoiceInput.Date
	}
	if invoiceInput.OrganizationID != nil {
		var organization *models.Organization
		for _, o := range s.organizations {
			if o.ID == *invoiceInput.OrganizationID {
				organization = o
				break
			}
		}
		if organization == nil {
			return nil, errors.New("models.Organization not found.")
		}
		invoice.From = organization
	}
	if invoiceInput.CustomerID != nil {
		var customer *models.Customer
		for _, c := range s.customers {
			if c.ID == *invoiceInput.CustomerID {
				customer = c
				break
			}
		}
		if customer == nil {
			return nil, errors.New("models.Customer not found.")
		}
		invoice.To = customer
	}
	if invoiceInput.Items != nil {
		invoice.Items = nil
		for _, input := range invoiceInput.Items {
			item := &models.InvoiceItem{
				ID:          fmt.Sprintf("%d", rand.Int()),
				Description: input.Description,
				UnitPrice:   input.UnitPrice,
				Quantity:    input.Quantity,
			}
			invoice.Items = append(invoice.Items, item)
		}
	}
	return invoice, nil
}

func (s *InMemoryStore) DeleteInvoice(ctx context.Context, id string) error {
	var idx *int
	for n, i := range s.invoices {
		if i.ID == id {
			idx = &n
			break
		}
	}
	if idx == nil {
		return errors.New("models.Invoice not found.")
	}
	s.invoices = append(s.invoices[:*idx], s.invoices[*idx+1:]...)
	return nil
}

func (s *InMemoryStore) GetCustomers(context.Context) ([]*models.Customer, error) {
	return s.customers, nil
}

func (s *InMemoryStore) CreateCustomer(ctx context.Context, customerInput models.NewCustomer) (*models.Customer, error) {
	customer := &models.Customer{
		Address: customerInput.Address,
		Email:   customerInput.Email,
		Name:    customerInput.Name,
		Phone:   customerInput.Phone,
		TaxID:   customerInput.TaxID,
	}
	customerID += 1
	customer.ID = fmt.Sprintf("%d", customerID)
	s.customers = append(s.customers, customer)
	return customer, nil
}

func (s *InMemoryStore) EditCustomer(ctx context.Context, customerInput models.EditCustomer) (*models.Customer, error) {
	var customer *models.Customer
	for _, c := range s.customers {
		if c.ID == customerInput.ID {
			customer = c
			break
		}
	}
	if customer == nil {
		return nil, errors.New("models.Customer not found.")
	}
	if customerInput.Address != nil {
		customer.Address = *customerInput.Address
	}
	if customerInput.Email != nil {
		customer.Email = *customerInput.Email
	}
	if customerInput.Name != nil {
		customer.Name = *customerInput.Name
	}
	if customerInput.Phone != nil {
		customer.Phone = *customerInput.Phone
	}
	if customerInput.TaxID != nil {
		customer.TaxID = *customerInput.TaxID
	}
	return customer, nil
}

func (s *InMemoryStore) DeleteCustomer(ctx context.Context, id string) error {
	var idx *int
	for n, i := range s.customers {
		if i.ID == id {
			idx = &n
			break
		}
	}
	if idx == nil {
		return errors.New("models.Customer not found.")
	}
	s.customers = append(s.customers[:*idx], s.customers[*idx+1:]...)
	return nil
}

func NewInMemoryStore() *InMemoryStore {
	firstUser := &models.User{
		ID:    "1",
		Email: "first@email.com",
		Name:  "First",
	}
	secondUser := &models.User{
		ID:    "2",
		Email: "second@email.com",
		Name:  "Second",
	}
	thirdUser := &models.User{
		ID:    "3",
		Email: "third@email.com",
		Name:  "Third",
	}
	adminRole := &models.Role{
		User:   firstUser,
		Access: models.RoleAccessAdmin,
	}
	readRole := &models.Role{
		User:   secondUser,
		Access: models.RoleAccessRead,
	}
	writeRole := &models.Role{
		User:   thirdUser,
		Access: models.RoleAccessWrite,
	}
	org1 := &models.Organization{
		ID:      "1",
		Address: "Some street 1234\nSome floor, ZIPCODE\nCity, Country.",
		Email:   "info@org.com",
		Name:    "models.Organization 1",
		Phone:   "555 5555 5555",
		Roles:   []*models.Role{adminRole, readRole, writeRole},
		TaxID:   "AAA-AAAA-AAA",
	}
	store := &InMemoryStore{
		users:         []*models.User{firstUser, secondUser, thirdUser},
		roles:         []*models.Role{adminRole, readRole, writeRole},
		organizations: []*models.Organization{org1},
	}
	customer1, _ := store.CreateCustomer(
		context.TODO(),
		models.NewCustomer{
			Address: "Some street 1234\nSome floor, ZIPCODE\nCity, Country.",
			Email:   "info@customer.com",
			Name:    "models.Customer 1",
			Phone:   "555 5555 5555",
			TaxID:   "NNN-NNNNN-NNN",
		},
	)
	store.CreateInvoice(
		context.TODO(),
		models.NewInvoice{
			Identifier:     "#0001",
			Date:           time.Now(),
			OrganizationID: org1.ID,
			CustomerID:     customer1.ID,
			Items: []*models.NewInvoiceItem{
				&models.NewInvoiceItem{
					Description: "Thing per kg",
					UnitPrice:   "20.0",
					Quantity:    "1.0",
				},
				&models.NewInvoiceItem{
					Description: "Other Thing per kg",
					UnitPrice:   "10.0",
					Quantity:    "1.5",
				},
			},
		},
	)
	return store
}
