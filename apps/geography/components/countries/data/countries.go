package data

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	allCountries = &Countries{}
	Data         = allCountries.List
)

type (
	Countries struct {
		List     []Country
		IsCached bool
	}

	Country struct {
		Name           Name                       `json:"name"`
		TopLevelDomain []string                   `json:"tld"`
		Cca2           string                     `json:"cca2"`
		Ccn3           string                     `json:"ccn3"`
		Cca3           string                     `json:"cca3"`
		Cioc           string                     `json:"cioc"`
		Independent    bool                       `json:"independent"`
		Status         string                     `json:"status"`
		UNMember       bool                       `json:"unMember"`
		Currencies     Currency                   `json:"currencies"`
		Idd            InternationalDirectDialing `json:"idd"`
		Capital        []string                   `json:"capital"`
		AltSpellings   []string                   `json:"altSpellings"`
		Translations   map[string]TranslationLang `json:"translations"`
		LatLng         []float64                  `json:"latlng"`
		Landlocked     bool                       `json:"landlocked"`
		Borders        []string                   `json:"borders"`
		Area           float64                    `json:"area"`
		Demonyms       map[string]Demonym         `json:"demonyms"`
		Population     int32                      `json:"population"`
		StartOfWeek    string                     `json:"startOfWeek"`
		FlagSrc        Flag                       `json:"flags"`

		// Active - used for search rows/cards
		Active bool
		// Selected - used for CP
		Selected bool
		// for CP details and contextual view
		IsCPViewed bool
		IsViewed   bool
	}

	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	}

	InternationalDirectDialing struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	}

	TranslationLang struct {
		Official string `json:"official"`
		Common   string `json:"common"`
	}

	Demonym struct {
		Female string `json:"f"`
		Male   string `json:"m"`
	}

	Flag struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	}

	Currency struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}
)

func (c *Countries) GetSelected() []Country {
	res := make([]Country, 0)
	for i := range Data {
		if Data[i].Selected {
			res = append(res, Data[i])
		}
	}
	return res
}

func (c *Countries) GetSelectedCount() int {
	count := 0
	for i := range Data {
		if Data[i].Selected {
			count++
		}
	}
	return count
}

func (c *Countries) InitCountries() error {
	if !c.IsCached {
		data, err := c.fetchCountries("all")
		if err != nil {
			log.Fatalln("error fetching data from RESTCountries API ", err.Error())
			return err
		}
		err = json.Unmarshal(data, &Data)
		if err != nil {
			log.Fatalln("json Unmarshal RESTCountries for mutable: ", err.Error())
			return err
		}
		c.IsCached = true
	}
	return nil
}

func (c *Countries) fetchCountries(location string) ([]byte, error) {
	URL := fmt.Sprintf("https://restcountries.com/v3.1/%s", location)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(fmt.Sprintf("http.Get(\"%s\") failed: ", URL), err.Error())
		return []byte{}, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatalln("error at deferred func, at the end: ", err.Error())
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("error at res.Body reading: ", err.Error())
		return []byte{}, err
	}
	return body, nil
}
