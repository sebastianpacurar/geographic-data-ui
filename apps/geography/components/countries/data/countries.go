package data

import (
	"encoding/json"
	"gioui-experiment/apps/geography/components/api"
	"log"
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

func (c *Countries) InitCountries() error {
	if !c.IsCached {
		data, err := api.FetchCountries("all")
		if err != nil {
			log.Fatalln("error fetching data from RESTCountries API ", err.Error())
			return err
		}
		err = json.Unmarshal(data, &Data)
		if err != nil {
			log.Fatalln("json Unmarshal RESTCountries response: ", err.Error())
			return err
		}
		c.IsCached = true
	}
	return nil
}

//func (c *Countries) FlagSrcToPNG(url, srcName string) {
//	if _, err := os.Stat(fmt.Sprintf("/flags/%s.png", srcName)); err != nil {
//		r, re := http.Get(url)
//		if re != nil {
//			log.Fatalln(fmt.Sprintf("error fetching the image of %s", srcName))
//		}
//		defer func(Body io.ReadCloser) {
//			err := Body.Close()
//			if err != nil {
//				log.Fatalln(fmt.Sprintf("error when closing response body reader for %s", srcName))
//			}
//		}(r.Body)
//
//		f, fe := os.Create(fmt.Sprintf("/FlagSrcs/%s.png", srcName))
//		if fe != nil {
//			log.Fatalln(fmt.Sprintf("error when creating country FlagSrc PNG as %s.png", srcName))
//		}
//		defer func(f *os.File) {
//			err := f.Close()
//			if err != nil {
//				log.Fatalln(fmt.Sprintf("error when closing the file creation process: %s", err.Error()))
//			}
//		}(f)
//
//		_, err := io.Copy(f, r.Body)
//		if err != nil {
//			log.Fatal(fmt.Sprintf("error when copying the bytes to png file: %s", err.Error()))
//		}
//	}
//}
