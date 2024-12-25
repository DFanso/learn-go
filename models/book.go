package models

import (
	"time"
)

// swagger:model
type Book struct {
	// ID is the primary key
	// required: true
	// swaggerignore:true
	ID uint `json:"-"`

	// swaggerignore:true
	// CreatedAt records the creation time
	// required: true
	CreatedAt time.Time `json:"-"`

	// UpdatedAt records the last update time
	// required: true
	// swaggerignore:true
	UpdatedAt time.Time `json:"-"`

	// DeletedAt is used for soft deletes
	// nullable: true
	// swaggerignore:true
	DeletedAt *time.Time `json:"-"`

	// Title of the book
	// required: true
	Title string `json:"title" binding:"required"`

	// Author of the book
	// required: true
	Author string `json:"author" binding:"required"`

	// Quantity of the book available
	// required: true
	Quantity int `json:"quantity" binding:"required"`

	// Availability status of the book
	Available bool `json:"available"`
}
