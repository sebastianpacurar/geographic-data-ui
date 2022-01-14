package data

import (
	"bytes"
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
)

type (
	Countries struct {
		CountriesList []Country
		IsCached      bool
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
		Flag           `json:"flags"`

		// FlagImg - obtained after decoding the FlagField png
		FlagImg image.Image
		// Active - used for search rows/cards
		Active bool
		// Selected - used for CP
		Selected bool
		// for CP details and contextual view
		IsCPViewed   bool
		IsCtxtActive bool
		// for Continent Tab Selection
		ActiveContinent bool
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

	Flag struct {
		Png string `json:"png"`
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
func downloadFlagFromUrl(url string) ([]byte, error) {
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
	var data bytes.Buffer
	_, e = io.Copy(&data, resp.Body)
	if e != nil {
		return nil, e
	}
	return data.Bytes(), nil
}

// ProcessFlags - decode png image for 4 countries at once
func ProcessFlags() {
	firstBatch := make(chan image.Image, len(Cached)/4)
	secondBatch := make(chan image.Image, len(Cached)/4)
	thirdBatch := make(chan image.Image, len(Cached)/4)
	fourthBatch := make(chan image.Image, len(Cached)/4)

	go downloadAndDecodeFlag(Cached[:63], firstBatch)
	go downloadAndDecodeFlag(Cached[63:126], secondBatch)
	go downloadAndDecodeFlag(Cached[126:189], thirdBatch)
	go downloadAndDecodeFlag(Cached[189:], fourthBatch)
}

// downloadAndDecodeFlag - download, decode and attach flag to country through channels
func downloadAndDecodeFlag(countries []Country, done chan image.Image) {
	for i := range countries {
		b, err := downloadFlagFromUrl(countries[i].Png)
		if err != nil {
			done <- nil
			return
		}
		img, _, _ := image.Decode(bytes.NewReader(b))
		done <- img
		countries[i].FlagImg = <-done
	}
	close(done)
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
		c.IsCached = true
	}
	return nil
}

// WriteFlagToFile - needed for slower OSes to store the flags locally, for quicker retrieval on next app start
func (c *Country) WriteFlagToFile() error {
	count, err := fileCount("output/geography/flags")
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error at counting files in output/geography/flags: %s", err))
	}

	if count < 250 {
		resp, e := http.Get(c.Png)
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

// ReadFlagFromFile - Currently faster than ProcessFlagFromUrl, useful for slower OSes
func ReadFlagFromFile() []byte {
	file, err := ioutil.ReadFile("output/geography/flags/placeholder/no-flag.png")
	if err != nil {
		log.Fatalln("Error opening no-flag.png path")
	}
	return file
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
