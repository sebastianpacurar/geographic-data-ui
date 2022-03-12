package table

const (
	NAME                  = "Name"
	CAPITALS              = "Capitals"
	REGION                = "Region"
	SUBREGION             = "Subregion"
	CONTINENTS            = "Continents"
	IDD_ROOT              = "IDD Root"
	IDD_SUFFIXES          = "IDD Suffixes"
	TOP_LEVEL_DOMAINS     = "Top Level Domains"
	INDEPENDENT           = "Independent"
	STATUS                = "Status"
	UNITED_NATIONS_MEMBER = "UN Member"
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
	OFFICIAL_NAME         = "Official Name"
	LANGUAGES             = "Languages"
)

var (
	// SearchBy - used as widget.Bool Value to search countries by a column headcell, and defaults to Country.Name.Common
	SearchBy = NAME

	// ColNames - the current columns used to display data (NAME is included by default as sticky column)
	ColNames = []string{
		CAPITALS, REGION, SUBREGION, CONTINENTS, IDD_ROOT, IDD_SUFFIXES, TOP_LEVEL_DOMAINS,
		INDEPENDENT, STATUS, UNITED_NATIONS_MEMBER, LANDLOCKED, CCA2, CCA3, CCN3, CIOC, FIFA, AREA, POPULATION, LATITUDE,
		LONGITUDE, START_OF_WEEK, CAR_SIGNS, CAR_SIDE, OFFICIAL_NAME, LANGUAGES,
	}

	// SearchByCols - used to Search By a specific column
	SearchByCols = []string{
		NAME, CAPITALS, TOP_LEVEL_DOMAINS, INDEPENDENT, UNITED_NATIONS_MEMBER, LANDLOCKED, CCA2, CCA3,
		CCN3, LATITUDE, LONGITUDE, START_OF_WEEK, CAR_SIGNS, CAR_SIDE, OFFICIAL_NAME,
	}

	// ColState - used to show/hide columns from CP
	ColState = map[string]bool{
		NAME:                  true, // should always be true since it refers to the sticky column
		CAPITALS:              true,
		REGION:                true,
		SUBREGION:             true,
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
		OFFICIAL_NAME:         false,
		LANGUAGES:             false,
	}

	//ColPos = map[string]int{
	//	CAPITALS:              0,
	//	REGION:                1,
	//	SUBREGION:             2,
	//	CONTINENTS:            3,
	//	IDD_ROOT:              4,
	//	IDD_SUFFIXES:          5,
	//	TOP_LEVEL_DOMAINS:     6,
	//	INDEPENDENT:           7,
	//	STATUS:                8,
	//	UNITED_NATIONS_MEMBER: 9,
	//	LANDLOCKED:            10,
	//	CCA2:                  11,
	//	CCA3:                  12,
	//	CCN3:                  13,
	//	CIOC:                  14,
	//	FIFA:                  15,
	//	AREA:                  16,
	//	POPULATION:            17,
	//	LATITUDE:              18,
	//	LONGITUDE:             19,
	//	START_OF_WEEK:         20,
	//	CAR_SIGNS:             21,
	//	CAR_SIDE:              22,
	//	OFFICIAL_NAME:         23,
	//	LANGUAGES:             24,
	//}
)
