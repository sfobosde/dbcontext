package dbcontext

// Start migrations.
func Migrate() {
	model := getContextModel()

	if model.db != nil {
		entities := model.Entities

		if entities == nil || len(entities) == 0 {
			panic("No entities to migrate.")
		}
		model.db.AutoMigrate(entities...)
	} else {
		panic("No connection. Use Connect() before Migrate() call")
	}
}
