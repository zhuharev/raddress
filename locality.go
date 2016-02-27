package raddress

import (
	"encoding/json"
	"strings"
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
		e := json.Unmarshal(MustAsset("data/russia.json"), &Loc)
		if e != nil {
			panic(e)
		}
	}
	for i, l := range Loc.Data.Locality {
		Loc.Data.Locality[i].Name = fix(l.Name)
	}
	for i, l := range Loc.Data.Subjects {
		Loc.Data.Subjects[i].Name = fix(l.Name)
	}
}

func fix(name string) string {
	rp := strings.NewReplacer(" р-н", " район",
		" г", "",
		" обл", " область",
	)

	name = rp.Replace(name)
	if strings.HasSuffix(name, " Респ") {
		name = strings.TrimSuffix(name, " Респ")
		if strings.HasSuffix(name, "ая") {
			name = name + " Республика"
		} else {
			name = "Республика " + name
		}
	}
	if strings.HasPrefix(strings.ToLower(name), "москва") {
		name = "Москва и область"
	}
	return name
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

func (l Locality) Cities(regionId int) []Item {
	name := strings.ToLower(l.RegionName(regionId))
	name = strings.Replace(name, "область", "", -1)
	name = strings.TrimSpace(name)
	return GetRegionCities(name)
}

func (l Locality) RegionName(id int) string {
	for _, v := range l.Data.Subjects {
		if v.Regioncode == id {
			return v.Name
		}
	}
	return ""
}

func (l Locality) RegionId(name string) int {
	fn := strings.Split(name, " ")[0]
	for _, v := range l.Data.Subjects {
		if strings.HasPrefix(v.Name, fn) {
			return v.Regioncode
		}
	}
	return 0
}
