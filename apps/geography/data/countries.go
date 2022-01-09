package data

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	allCountries = &Countries{}
	Cached       = allCountries.CountriesList

	// AllFlags - used only to get all flags correctly from stable v2 api version
	AllFlags = allCountries.FlagsList
)

type (
	Countries struct {
		CountriesList []Country
		FlagsList     []CountryFlag // used only to get all flags correctly from stable v2 api version
		IsCached      bool
	}

	// CountryFlag - Used with v2 API - Caution: country naming discrepancies!
	CountryFlag struct {
		FlagField Flag `json:"flags"`
		// FlagImg - obtained after decoding the FlagField png
		FlagImg image.Image
	}
	Flag struct {
		Png string `json:"png"`
	}

	// Country - Used with v3.1 API
	Country struct {
		Name           Name                       `json:"name"`
		TopLevelDomain []string                   `json:"tld"`
		Cca2           string                     `json:"cca2"`
		Ccn3           string                     `json:"ccn3"`
		Cca3           string                     `json:"cca3"`
		Cioc           string                     `json:"cioc"`
		Fifa           string                     `json:"fifa"`
		Independent    bool                       `json:"independent"`
		Status         string                     `json:"status"`
		UNMember       bool                       `json:"unMember"`
		Currencies     map[string]Currency        `json:"currencies"`
		Idd            InternationalDirectDialing `json:"idd"`
		Car            Car                        `json:"car"`
		Capital        []string                   `json:"capital"`
		AltSpellings   []string                   `json:"altSpellings"`
		Translations   map[string]TranslationLang `json:"translations"`
		Languages      map[string]string          `json:"languages"`
		LatLng         []float64                  `json:"latlng"`
		Landlocked     bool                       `json:"landlocked"`
		Borders        []string                   `json:"borders"`
		Area           float64                    `json:"area"`
		Demonyms       map[string]Demonym         `json:"demonyms"`
		Population     int32                      `json:"population"`
		StartOfWeek    string                     `json:"startOfWeek"`
		Region         string                     `json:"region"`
		Subregion      string                     `json:"subregion"`
		Continents     []string                   `json:"continents"`

		// Active - used for search rows/cards
		Active bool
		// Selected - used for CP
		Selected bool
		// for CP details and contextual view
		IsCPViewed   bool
		IsCtxtActive bool

		CountryFlag
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

	Currency struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	Car struct {
		Signs []string `json:"signs"`
		Side  string   `json:"side"`
	}
)

func (c *Countries) GetSelected() []Country {
	res := make([]Country, 0)
	for i := range Cached {
		if Cached[i].Selected {
			res = append(res, Cached[i])
		}
	}
	return res
}

func (c *Countries) GetSelectedCount() int {
	count := 0
	for i := range Cached {
		if Cached[i].Selected {
			count++
		}
	}
	return count
}

// ProcessFlagFromUrl - decode image directly from url  (very slow)
func (c *Country) ProcessFlagFromUrl(url string) error {
	resp, e := http.Get(url)
	if e != nil {
		log.Fatalln(e)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	img, _, _ := image.Decode(resp.Body)
	c.FlagImg = img
	return nil
}

func (c *Countries) InitCountries() error {
	if !c.IsCached {
		countries, err := c.fetchCountries("all")
		if err != nil {
			log.Fatalln("error fetching data from RESTCountries API ", err)
			return err
		}
		err = json.Unmarshal(countries, &Cached)
		if err != nil {
			log.Fatalln("json Unmarshal RESTCountries for mutable: ", err)
			return err
		}

		flags, err := c.fetchFlags()
		if err != nil {
			log.Fatalln("error fetching data from RESTCountries API ", err)
			return err
		}
		err = json.Unmarshal(flags, &AllFlags)
		if err != nil {
			log.Fatalln("json Unmarshal RESTCountries for mutable: ", err)
			return err
		}
		c.IsCached = true
	}
	return nil
}

func (c *Country) WriteFlagToFile() error {
	count, err := fileCount("output/geography/flags")
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error at counting files in output/geography/flags: %s", err))
	}

	if count < 250 {
		resp, e := http.Get(c.FlagField.Png)
		if e != nil {
			log.Fatalln(e)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(resp.Body)

		file, err := os.Create(fmt.Sprintf("output/geography/flags/%s.png", c.Name.Common))
		if err != nil {
			log.Fatal(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(file)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

// ReadFlagFromFile - Currently faster than ProcessFlagFromUrl, although it demands the png files present in output/flags
func (c *Country) ReadFlagFromFile() error {
	path := fmt.Sprintf("output/geography/flags/%s.png", c.Name.Common)

	file, err := os.Open(path)
	if err != nil {
		_ = c.WriteFlagToFile()
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "output/geography/flags/placeholder/no-flag.png"
	}

	file, err = os.Open(path)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error opening path: %s", path))
	}

	img, _, _ := image.Decode(file)
	c.FlagImg = img
	return nil
}

// fetchFlags - Fetches only the flags and saves them in CountryFlag
func (c *Countries) fetchFlags() ([]byte, error) {
	URL := fmt.Sprintf("https://restcountries.com/v2/all")
	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(fmt.Sprintf("http.Get(\"%s\") failed: %s", URL, err))
		return []byte{}, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return []byte{}, err
	}
	return body, nil
}

// fetchCountries - Fetches All Country Data except the flag
func (c *Countries) fetchCountries(location string) ([]byte, error) {
	URL := fmt.Sprintf("https://restcountries.com/v3.1/%s", location)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(fmt.Sprintf("http.Get(\"%s\") failed: %s", URL, err))
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

func fileCount(path string) (int, error) {
	i := 0
	entry, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error when reading directory %s", err))
	}
	for _, f := range entry {
		if !f.IsDir() {
			i++
		}
	}
	return i, nil
}
