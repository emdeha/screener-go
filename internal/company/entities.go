package company

type Company struct {
	CIK        int    `bson:"cik"`
	EntityName string `bson:"entityName"`
	Facts      Facts  `bson:"facts"`
}

type Facts struct {
	DEI    interface{} `bson:"dei"`
	USGaap interface{} `bson:"us-gaap"`
}
