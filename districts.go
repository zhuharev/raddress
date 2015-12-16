package raddress

type District struct {
	Id   int
	Name string
}

func (d District) Childs() (rs Regions) {
	for _, v := range Loc.Data.Subjects {
		if v.Disid == d.Id {
			rs = append(rs, v)
		}
	}
	return
}

/*func (r Region) Localities() (rs Regions) {
	for _, v := range Locality.Data.Locality {
		if v.Parentguid == r.Aoguid {
			rs = append(rs, v)
		}
	}
	return
}*/

type Districts []District

func (d Districts) Subjects(id int) (rs Regions) {
	for _, v := range Loc.Data.Subjects {
		if v.Regioncode == id {
			rs = append(rs, v)
		}
	}
	return
}

/*func (d District) Localities(regionId int) (rs Regions) {
	for _, v := range Loc.Data.Locality {
		if v.Regioncode == regionId && v.Parentguid == v.Aoguid {
			rs = append(rs, v)
		}
	}
	return
}*/

var (
	districts []string = []string{
		"Центральный",
		"Южный",
		"Северо-Западный",
		"Дальневосточный",
		"Сибирский",
		"Уральский",
		"Приволжский",
		"Северо-Кавказский",
	}
	suffix = " федеральный округ"
)

func GetDistrictById(id int) string {
	return districts[id-1] + suffix
}

func GetDistricts() (d Districts) {
	for k, v := range districts {
		d = append(d, District{k + 1, v + suffix})
	}
	return
}
