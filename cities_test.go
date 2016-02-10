package raddress

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCities(t *testing.T) {
	Convey("Test get city", t, func() {

		Convey("Test get city", func() {
			item := GetCityByName("Москва")
			So(item.Id, ShouldNotEqual, 0)
			So(item.RegionId, ShouldNotEqual, 0)
		})

		Convey("Get city region", func() {
			item := GetCityRegion("москва")
			So(item.Id, ShouldNotEqual, 0)
		})

		Convey("Get region", func() {
			items := GetRegionCities("Смоленска")
			So(len(items), ShouldNotEqual, 0)
		})

	})

}

func TestLoc(t *testing.T) {
	fmt.Println(Loc.RegionName(67))
	Convey("test Loc", t, func() {
		items := Loc.Cities(67)
		So(len(items), ShouldNotEqual, 0)
	})
}
