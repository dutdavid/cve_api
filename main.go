package main

import (
	"github.com/dutdavid/cve_api/routes"
	"github.com/dutdavid/cve_api/scheduler"
)

func main() {
	go scheduler.StartScheduler() // Start scheduler in background

	router := routes.SetupRoutes()
	router.Run(":8080")
}
