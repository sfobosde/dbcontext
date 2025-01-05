package dbcontext_test

// Запуск тестов.
func Execute() {
	defer cleanUp()

	properties := getConnectionPropertiesEnv()

	testConnect(properties)

	models := initModel()
	testMigrate()

	testCRUD(models)
	testGlobalOperands(models)
	testFieldOperands(models)
}
