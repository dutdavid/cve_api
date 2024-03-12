package scheduler

import (
	"fmt"
	"time"

	"github.com/dutdavid/cve_api/datafetch"
)

func StartScheduler() {
	ticker := time.NewTicker(24 * time.Hour)

	for range ticker.C {
		err := datafetch.FetchAndStoreCVEs()
		if err != nil {
			// Add error handling, e.g., logging
			fmt.Println("Error fetching CVEs:", err)
		}
	}
}
