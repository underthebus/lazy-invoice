//go:generate rm -f graphql/generated.go
//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"

	"github.com/underthebus/lazy-invoice/backend/models"
	"github.com/underthebus/lazy-invoice/backend/store"
)

type Resolver struct {
	store store.Storer
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateInvoice(ctx context.Context, invoice models.NewInvoice) (*models.Invoice, error) {
	return r.store.CreateInvoice(ctx, invoice)
}
func (r *mutationResolver) EditInvoice(ctx context.Context, invoice models.EditInvoice) (*models.Invoice, error) {
	return r.store.EditInvoice(ctx, invoice)
}
func (r *mutationResolver) DeleteInvoice(ctx context.Context, id string) (string, error) {
	return id, r.store.DeleteInvoice(ctx, id)
}
func (r *mutationResolver) CreateCustomer(ctx context.Context, customer models.NewCustomer) (*models.Customer, error) {
	return r.store.CreateCustomer(ctx, customer)
}
func (r *mutationResolver) EditCustomer(ctx context.Context, customer models.EditCustomer) (*models.Customer, error) {
	return r.store.EditCustomer(ctx, customer)
}
func (r *mutationResolver) DeleteCustomer(ctx context.Context, id string) (string, error) {
	return id, r.store.DeleteCustomer(ctx, id)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Invoices(ctx context.Context) ([]*models.Invoice, error) {
	return r.store.GetInvoices(ctx)
}
func (r *queryResolver) Customers(ctx context.Context) ([]*models.Customer, error) {
	return r.store.GetCustomers(ctx)
}

func NewResolver(store store.Storer) *Resolver {
	return &Resolver{store: store}
}
