package dbcontext_test

import (
	"github.com/sfobosde/dbcontext/dbcontext"
)

// Queue of methods intended to clean data.
var cleaninQueue []func()

// Cleaning database after test ends.
func cleanUp() {
	for _, clean := range cleaninQueue {
		clean()
	}
}

// Add tp clean up queue.
func subscribe(clean func()) {
	cleaninQueue = append(cleaninQueue, clean)
}

// Create model with autoremove.
func create[TEntity, TSearch any](model *dbcontext.DataModel[TEntity, TSearch]) *TEntity {
	entity := model.Create()

	subscribe(func() { model.Delete(entity) })
	return entity
}
