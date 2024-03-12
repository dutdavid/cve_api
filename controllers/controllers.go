package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/dutdavid/cve_api/models"
	"github.com/dutdavid/cve_api/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetCVEByCVEID godoc
// @Summary Get CVE details by CVE ID
// @Description Get a CVE by its CVE ID
// @ID get-cve-by-id
// @Param cveID path string true "CVE ID"
// @Produce json
// @Success 200 {object} models.CVE
// @Failure 404 {object} gin.H
// @Router /cve/{cveID} [get]
func GetCVEByCVEID(c *gin.Context) {
	cveID := c.Param("cveID")

	client, collection := mongodb.Connect()
	defer client.Disconnect(context.TODO())

	var result models.CVE
	err := collection.FindOne(context.TODO(), bson.M{"cveID": cveID}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CVE not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetCVEByDate godoc
// @Summary Get CVEs by date added
// @Description Get CVEs based on the date they were added to the CISA feed
// @ID get-cves-by-date
// @Param dateAdded path string true "Date Added (YYYY-MM-DD)"
// @Produce json
// @Success 200 {array} models.CVE
// @Failure 400 {object} gin.H
// @Router /cve/date/{dateAdded} [get]
func GetCVEByDate(c *gin.Context) {
	dateStr := c.Param("dateAdded")
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	client, collection := mongodb.Connect()
	defer client.Disconnect(context.TODO())

	filter := bson.M{"dateAdded": date}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching CVEs"})
		return
	}

	var results []models.CVE
	if err = cursor.All(context.TODO(), &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding CVEs"})
		return
	}

	c.JSON(http.StatusOK, results)
}
