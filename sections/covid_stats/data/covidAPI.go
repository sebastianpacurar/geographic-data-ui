package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	covStats   = &CovidStats{}
	CachedDay  = covStats.Day
	CachedDays = make(map[string][]CovidData)
)

type (
	CovidStats struct {
		Day      []CovidData
		IsCached bool
	}

	CovidData struct {
		Country     string `json:"countryRegion"`
		AdminArea   string `json:"admin2"`
		State       string `json:"provinceState"`
		CombinedKey string `json:"combinedKey"`
		LastUpdate  string `json:"lastUpdate"`
		Lat         string `json:"lat"`
		Lng         string `json:"lng"`
		Confirmed   string `json:"confirmed"`
		Deaths      string `json:"deaths"`
	}
)

// fetchDates - Fetches All Country Data except the flag
func fetchDates(date string) ([]byte, error) {
	client := &http.Client{}
	URL := fmt.Sprintf("https://covid19.mathdro.id/api/daily/%s", date)
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Fatalln(fmt.Sprintf("http.Get(\"%s\") failed: %s", URL, err))
		return []byte{}, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatalln("error at deferred func, at the end: ", err)
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("error at res.Body reading: ", err)
		return []byte{}, err
	}
	return body, nil
}

// TODO: not working - problematic for continuous recalls
// InitDayData - parse the data and store it in a map[date]CovidData, where date is the day as string
func InitDayData(day string) error {
	// if the specific day is not already cached, then add it to collection with its relevant data
	if _, ok := CachedDays[day]; !ok {
		data, err := fetchDates(day)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &CachedDay)
		if err != nil {
			log.Print("json Unmarshal CovidStats data ", err)
			return err
		}

		// add all countries data to Date map
		CachedDays[day] = CachedDay
	}
	return nil
}
