package raddress

import (
	"encoding/json"
)

var Loc Locality

type Locality struct {
	Data struct {
		Districts []struct {
			Areas []struct {
				Aoguid string `json:"aoguid"`
				Name   string `json:"name"`
				Okato  string `json:"okato"`
			} `json:"areas"`
			ID   int    `json:"id,string"`
			Name string `json:"name"`
		} `json:"districts"`
		Locality []Region `json:"locality"`
		Subjects []Region `json:"subjects"`
	} `json:"data"`
}

func init() {
	if BinDataEnabled {
		e := json.Unmarshal(MustAsset("russia.json"), &Loc)
		if e != nil {
			panic(e)
		}
	}
}

func (l Locality) Districts() (res Districts) {
	for _, v := range l.Data.Districts {
		res = append(res, District{v.ID, v.Name})
	}
	return
}

func (l Locality) Localities(regionId int) (res Regions) {
	for _, region := range l.Data.Locality {
		if region.Regioncode == regionId {
			res = append(res, region)
		}
	}
	return
}
