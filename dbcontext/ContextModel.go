package dbcontext

import (
	"sync"

	"gorm.io/gorm"
)

var (
	// Instance of ContextModel (Singletone),
	context *ContextModel
	once    sync.Once
)

// Get single instance of context model.
func getContextModel() *ContextModel {
	once.Do(func() {
		context = &ContextModel{}
	})

	return context
}

// Context model structure.
type ContextModel struct {
	db       *gorm.DB
	Entities []interface{}
}
