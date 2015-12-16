package raddress

import (
	"github.com/ungerik/go-dry"
)

var (
	RussianRegions     Regions
	DefaultRegionsPath = "regions.json"

	BinDataEnabled = true
)

type Region struct {
	Aoguid     string `json:"aoguid"`
	Disid      int    `json:"disid,string"`
	Name       string `json:"name"`
	Okato      string `json:"okato"`
	Parentguid string `json:"parentguid"`
	Regioncode int    `json:"regioncode,string"`
}

func (r Region) Childs() (res Regions) {
	for _, v := range Loc.Data.Locality {
		if v.Parentguid == r.Aoguid {
			res = append(res, v)
		}
	}
	return
}

type Regions []Region

func Load(path ...string) error {
	p := DefaultRegionsPath
	if len(path) > 0 {
		p = path[0]
	}
	return dry.FileUnmarshallJSON(p, &RussianRegions)
}

func MustLoad(path ...string) {
	e := Load(path...)
	if e != nil {
		panic(e)
	}
}
