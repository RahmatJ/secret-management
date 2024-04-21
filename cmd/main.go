package main

import (
	"secret-management/app"
	"secret-management/migration"
)

func main() {
	migration.MigrateUp()
	app.Run()
}
