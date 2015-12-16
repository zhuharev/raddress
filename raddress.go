package raddress

import (
	"encoding/json"
	"fmt"
	"github.com/ungerik/go-dry"
)

var (
	RussianRegions     Regions
	DefaultRegionsPath = "regions.json"

	BinDataEnabled = true
)

type Region struct {
	Name       string `json:"name"`
	RegionCode int    `json:"rid"`
	DistrictId int    `json:"did"`
}

type Regions []Region

func init() {
	if BinDataEnabled {
		var regs []Region
		fmt.Println(string(MustAsset("regions.json")))
		e := json.Unmarshal(MustAsset("regions.json"), &regs)
		if e != nil {
			panic(e)
		}
		RussianRegions = Regions(regs)
	}
}

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
