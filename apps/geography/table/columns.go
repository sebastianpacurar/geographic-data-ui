package table

const (
	NAME                  = "Name"
	OFFICIAL_NAME         = "Official Name"
	CAPITAL               = "Capital"
	REGION                = "Region"
	SUBREGION             = "Subregion"
	LANGUAGES             = "Languages"
	CONTINENTS            = "Continents"
	IDD_ROOT              = "IDD Root"
	IDD_SUFFIXES          = "IDD Suffixes"
	TOP_LEVEL_DOMAINS     = "Top Level Domains"
	INDEPENDENT           = "Independent"
	STATUS                = "Status"
	UNITED_NATIONS_MEMBER = "United Nations Member"
	LANDLOCKED            = "Landlocked"
	CCA2                  = "CCA 2"
	CCA3                  = "CCA 3"
	CCN3                  = "CCN 3"
	CIOC                  = "IOC Code"
	FIFA                  = "FIFA Code"
	AREA                  = "Area"
	POPULATION            = "Population"
	LATITUDE              = "Latitude"
	LONGITUDE             = "Longitude"
	START_OF_WEEK         = "Start of Week"
	CAR_SIGNS             = "Car Signs"
	CAR_SIDE              = "Car Side"
)

// ColState - used to show/hide columns from CP
var ColState = map[string]bool{
	NAME:                  true,
	OFFICIAL_NAME:         false,
	CAPITAL:               true,
	REGION:                true,
	SUBREGION:             true,
	LANGUAGES:             false,
	CONTINENTS:            false,
	IDD_ROOT:              false,
	IDD_SUFFIXES:          false,
	TOP_LEVEL_DOMAINS:     true,
	INDEPENDENT:           true,
	STATUS:                false,
	UNITED_NATIONS_MEMBER: true,
	LANDLOCKED:            false,
	CCA2:                  true,
	CCA3:                  false,
	CCN3:                  false,
	CIOC:                  false,
	FIFA:                  false,
	AREA:                  true,
	POPULATION:            true,
	LATITUDE:              true,
	LONGITUDE:             true,
	START_OF_WEEK:         false,
	CAR_SIGNS:             true,
	CAR_SIDE:              true,
}
