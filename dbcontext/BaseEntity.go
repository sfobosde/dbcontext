package dbcontext

import (
	"time"
)

// Base Entity describe with common fields: Id, CreatedAt, UpdatedAt.
type BaseEntity struct {
	// Entity Id.
	Id string `gorm:"primaryKey" json:"id"`

	// Create date/time.
	CreatedAt time.Time `json:"createdAt"`

	// Last update date/time.
	UpdatedAt time.Time `json:"updatedAt"`
}

// Base entity search fields.
type BaseEntitySearch struct {
	Id *StringFieldOperands `dbcontext:"Id"`
}
