package datafetch

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dutdavid/cve-api/models"
	"github.com/dutdavid/cve-api/mongodb"
)

func FetchAndStoreCVEs() error {
	url := "https://www.cisa.gov/sites/default/files/feeds/known_exploited_vulnerabilities.json"
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var cveCatalog struct {
		Vulnerabilities []models.CVE `json:"vulnerabilities"`
	}
	err = json.Unmarshal(data, &cveCatalog)
	if err != nil {
		return err
	}

	client, collection := mongodb.Connect()
	defer client.Disconnect(context.TODO())

	for _, cve := range cveCatalog.Vulnerabilities {
		_, err = collection.InsertOne(context.TODO(), cve)
		if err != nil {
			return err
		}
	}

	fmt.Println("CVEs stored in MongoDB!")
	return nil
}
