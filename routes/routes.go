// routes/routes.go
package routes

import (
	"github.com/dutdavid/cve_api/controllers"
	"github.com/dutdavid/cve_api/mongodb"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Connect to MongoDB before starting the API
	mongodb.Connect()

	// API routes
	router.GET("/cve/:cveID", controllers.GetCVEByCVEID)
	router.GET("/cve/date/:dateAdded", controllers.GetCVEByDate)

	return router
}
