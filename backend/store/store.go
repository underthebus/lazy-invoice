// Interfaces store implementations must follow
package store

import (
	"context"

	"github.com/underthebus/lazy-invoice/backend/models"
)

// InvoiceStorer defines interfaces for interaction with invoice objects.
type InvoiceStorer interface {
	// GetInvoices returns a list of invoices matching query.
	GetInvoices(ctx context.Context) ([]*models.Invoice, error)
	// CreateInvoice creates a new Invoice in the store.
	CreateInvoice(ctx context.Context, invoiceInput models.NewInvoice) (*models.Invoice, error)
	// EditInvoice edits an Invoice in the store.
	EditInvoice(ctx context.Context, invoiceInput models.EditInvoice) (*models.Invoice, error)
	// DeleteInvoice deletes an invoice from the store.
	DeleteInvoice(ctx context.Context, ID string) error
}

// CustomerStorer defines interfaces for interaction with customer objects.
type CustomerStorer interface {
	// GetCustomers returns a list of customers matching query.
	GetCustomers(ctx context.Context) ([]*models.Customer, error)
	// CreateCustomer creates a new customer in the store.
	CreateCustomer(ctx context.Context, customerInput models.NewCustomer) (*models.Customer, error)
	// EditCustomer edits a customer in the store.
	EditCustomer(ctx context.Context, customerInput models.EditCustomer) (*models.Customer, error)
	// DeleteCustomer deletes an customer from the store.
	DeleteCustomer(ctx context.Context, ID string) error
}

// Storer defines the interface to handle lazy-invoice data.
type Storer interface {
	InvoiceStorer
	CustomerStorer
}
