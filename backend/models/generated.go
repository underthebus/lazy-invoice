// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Customer struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	TaxID   string `json:"taxId"`
}

type EditCustomer struct {
	ID      string  `json:"id"`
	Address *string `json:"address"`
	Email   *string `json:"email"`
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	TaxID   *string `json:"taxId"`
}

type EditInvoice struct {
	ID             string            `json:"id"`
	Identifier     *string           `json:"identifier"`
	CustomerID     *string           `json:"customerId"`
	Date           *time.Time        `json:"date"`
	OrganizationID *string           `json:"organizationId"`
	Items          []*NewInvoiceItem `json:"items"`
}

type Invoice struct {
	ID         string         `json:"id"`
	Identifier string         `json:"identifier"`
	Date       time.Time      `json:"date"`
	From       *Organization  `json:"from"`
	To         *Customer      `json:"to"`
	Items      []*InvoiceItem `json:"items"`
}

type InvoiceItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	UnitPrice   string `json:"unitPrice"`
	Quantity    string `json:"quantity"`
}

type NewCustomer struct {
	Address string `json:"address"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	TaxID   string `json:"taxId"`
}

type NewInvoice struct {
	Identifier     string            `json:"identifier"`
	CustomerID     string            `json:"customerId"`
	Date           time.Time         `json:"date"`
	OrganizationID string            `json:"organizationId"`
	Items          []*NewInvoiceItem `json:"items"`
}

type NewInvoiceItem struct {
	Description string `json:"description"`
	UnitPrice   string `json:"unitPrice"`
	Quantity    string `json:"quantity"`
}

type Organization struct {
	ID      string  `json:"id"`
	Address string  `json:"address"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Roles   []*Role `json:"roles"`
	TaxID   string  `json:"taxId"`
}

type Role struct {
	User   *User      `json:"user"`
	Access RoleAccess `json:"access"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type RoleAccess string

const (
	RoleAccessAdmin RoleAccess = "ADMIN"
	RoleAccessRead  RoleAccess = "READ"
	RoleAccessWrite RoleAccess = "WRITE"
)

var AllRoleAccess = []RoleAccess{
	RoleAccessAdmin,
	RoleAccessRead,
	RoleAccessWrite,
}

func (e RoleAccess) IsValid() bool {
	switch e {
	case RoleAccessAdmin, RoleAccessRead, RoleAccessWrite:
		return true
	}
	return false
}

func (e RoleAccess) String() string {
	return string(e)
}

func (e *RoleAccess) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleAccess(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleAccess", str)
	}
	return nil
}

func (e RoleAccess) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
