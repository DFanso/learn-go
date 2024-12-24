package models

import (
	"time"
)

// Book represents the book model in the database
// swagger:model
type Book struct {
	// ID is the primary key
	// required: true
	// swaggerignore:true
	ID uint `json:"id" gorm:"primaryKey"`

	// CreatedAt records the creation time
	// required: true
	// swaggerignore:true
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt records the last update time
	// required: true
	// swaggerignore:true
	UpdatedAt time.Time `json:"updated_at"`

	// DeletedAt is used for soft deletes
	// nullable: true
	// swaggerignore:true
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`

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
