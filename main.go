package main

import (
	"github.com/dutdavid/cve-api/routes"
	"github.com/dutdavid/cve-api/scheduler"
)

func main() {
	go scheduler.StartScheduler() // Start scheduler in background

	router := routes.SetupRoutes()
	router.Run(":8080")
}
