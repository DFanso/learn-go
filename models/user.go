package models

import (
	"time"
)

// swagger:model
type User struct {
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

	// User Name
	// required: true
	UserName string `json:"userName" binding:"required"`

	// First Name
	// required: true
	FirstName string `json:"firstName" binding:"required"`

	// Last Name
	// required: true
	LastName string `json:"lastName" binding:"required"`

	// password of the user
	// required: true
	Password string `json:"password" binding:"required"`

	// Email of the user
	// required: true
	Email string `json:"email" binding:"required"`
}
