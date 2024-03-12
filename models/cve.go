package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CVE struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	CVEID              string             `bson:"cveID"`
	VendorProject      string             `bson:"vendorProject"`
	Product            string             `bson:"product"`
	VulnerabilityName  string             `bson:"vulnerabilityName"`
	DateAdded          time.Time          `bson:"dateAdded"`
	ShortDescription   string             `bson:"shortDescription"`
	RequiredAction     string             `bson:"requiredAction"`
	DueDate            time.Time          `bson:"dueDate"`
	KnownRansomwareUse string             `bson:"knownRansomwareCampaignUse"`
	Notes              string             `bson:"notes"`
}
