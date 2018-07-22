package config

const (
	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// ElasticSearch
	ElasticIndex = "dating_profile"

	// RPC service method
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	ItemSaverPort = 1234
	WorkerPort0   = 9000
)
