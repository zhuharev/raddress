package raddress

type District struct {
	Id   int
	Name string
}

type Districts []District

func (d Districts) Childs(id int) (rs Regions) {
	for _, v := range RussianRegions {
		if v.DistrictId == id {
			rs = append(rs, v)
		}
	}
	return
}

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
