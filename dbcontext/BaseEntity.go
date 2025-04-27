package dbcontext

import "time"

// Base Entity describe with common fields: ID, CreatedAt, UpdatedAt.
type BaseEntity struct {
	// Entity ID.
	ID string `gorm:"primaryKey" json:"id"`

	// Create date/time.
	CreatedAt time.Time `json:"createdAt"`

	// Last update date/time.
	UpdatedAt time.Time `json:"updatedAt"`
}

// Base entity search fields.
type BaseEntitySearch struct {
	ID *StringFieldOperands `dbcontext:"ID"`
}
