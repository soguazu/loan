package types

import "fmt"

// Messages creates types of response messages
type Messages string

const (
	// CREATED creates types of response messages for post endpoint
	CREATED Messages = "created successfully"
	// OKAY creates types of response messages for get endpoint
	OKAY = "retrieved successfully"
	// DELETED creates types of response messages for delete endpoint
	DELETED = "deleted successfully"
	// UPDATED creates types of response messages for patch endpoint
	UPDATED = "updated successfully"
)

// GetResponseMessage generates dynamic messages
func (m *Messages) GetResponseMessage(entity string, message Messages) string {
	return fmt.Sprintf("%v %v", entity, message)
}
