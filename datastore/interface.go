package datastore

import (
	"gofr.dev/pkg/gofr"
	"Gofr_Zopsmart/model"
)

type user interface {
	// GetByID retrieves a book record based on its ID.
	GetByID(ctx *gofr.Context, id string) (*model.user, error)
	// Create inserts a new book record into the database.
	Create(ctx *gofr.Context, model *model.user) (*model.user, error)
	// Update updates an existing book with the provided information.
	Update(ctx *gofr.Context, model *model.user) (*model.user, error)
	// Delete removes a entry record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}
