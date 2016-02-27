package raddress

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/fatih/color"
	"strings"
	"sync"
)

var csvApi *csvapi

type csvapi struct {
	regions []Item
	cities  []Item
}

type Item struct {
	Id       int
	RegionId int
	Name     string
}

var (
	CsvRegion = "cr"
	CsvCity   = "cc"
)

func init() {
	csvApi = new(csvapi)
	loadCsv(CsvRegion)
	loadCsv(CsvCity)
}

func loadCsv(t string) {

	var (
		r  *bytes.Reader
		cr *csv.Reader

		fileName string
		arr      []Item
	)

	switch t {
	case CsvRegion:
		fileName = "region"
	case CsvCity:
		fileName = "city"
	}

	r = bytes.NewReader(MustAsset("data/" + fileName + ".csv"))
	cr = csv.NewReader(r)
	cr.Comma = ';'

	// read header
	_, e := cr.Read()
	if e != nil {
		panic(e)
	}

	for line, e := cr.Read(); e == nil; line, e = cr.Read() {
		if len(line) != 4 {
			color.Red("Line should contain 4 elements, got %d", len(line))
			continue
		}

		// shoose Russia regions
		if line[1] != "3159" {
			continue
		}
		arr = append(arr,
			Item{Id: com.StrTo(line[0]).MustInt(),
				RegionId: com.StrTo(line[2]).MustInt(),
				Name:     line[3]})
	}
	switch t {
	case CsvRegion:
		csvApi.regions = append(csvApi.regions, arr...)
	case CsvCity:
		fileName = "city"
		csvApi.cities = append(csvApi.cities, arr...)
	}
}

func GetCityByName(city string) Item {
	city = strings.ToLower(city)
	for _, v := range csvApi.cities {
		if strings.ToLower(v.Name) == city {
			return v
		}
	}
	return Item{}
}

func GetCityRegion(city string) Item {
	c := GetCityByName(city)
	return GetRegion(c.RegionId)
}

func GetRegion(id int) Item {
	for _, v := range csvApi.regions {
		if v.Id == id {
			return v
		}
	}
	return Item{}
}

var (
	mx     sync.Mutex
	ccache = map[string][]Item{}
)

func GetRegionCities(region string) (res []Item) {
	mx.Lock()
	if res, ok := ccache[region]; ok {
		mx.Unlock()
		return res
	} else {
		mx.Unlock()
	}

	region = strings.ToLower(region)
	region = strings.TrimSuffix(region, " район")
	fmt.Println(region)
	for _, v := range csvApi.regions {
		if strings.HasPrefix(strings.ToLower(v.Name), region) {
			for _, city := range csvApi.cities {
				if city.RegionId == v.Id {
					if hasStr(city.Name, allowedCities) {
						res = append(res, city)
					}
				}
			}
		}
	}
	mx.Lock()
	ccache[region] = res
	mx.Unlock()
	return
}

var allowedCities []string

func AllowCities(arr []string) {
	allowedCities = arr
}

func hasStr(str string, arr []string) bool {
	if arr == nil {
		return true
	}
	for _, v := range arr {
		if strings.ToLower(v) == strings.ToLower(str) {
			return true
		}
	}
	return false
}
