package graphql

import (
	"context"

	"github.com/underthebus/lazy-invoice/backend/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateInvoice(ctx context.Context, invoice models.NewInvoice) (*models.Invoice, error) {
	panic("not implemented")
}
func (r *mutationResolver) EditInvoice(ctx context.Context, invoice models.EditInvoice) (*models.Invoice, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteInvoice(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateCustomer(ctx context.Context, customer models.NewCustomer) (*models.Customer, error) {
	panic("not implemented")
}
func (r *mutationResolver) EditCustomer(ctx context.Context, customer models.EditCustomer) (*models.Customer, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Invoices(ctx context.Context) ([]*models.Invoice, error) {
	panic("not implemented")
}
func (r *queryResolver) Customers(ctx context.Context) ([]*models.Customer, error) {
	panic("not implemented")
}
