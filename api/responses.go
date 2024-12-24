package api

// ErrorResponse represents the structure of error responses
// swagger:model
type ErrorResponse struct {
	// Error message
	// required: true
	Error string `json:"error" example:"Internal Server Error"`
}

// MessageResponse represents the structure of success messages
// swagger:model
type MessageResponse struct {
	// Success message
	// required: true
	Message string `json:"message" example:"Book deleted successfully"`
}
