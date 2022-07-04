package slicex_test

import (
	"testing"

	"github.com/mzccddkk/gdk/slicex"
	. "github.com/smartystreets/goconvey/convey"
)

var columnTest = []map[interface{}]interface{}{
	{
		"id":         1858,
		"first_name": "John",
		"last_name":  "Doe",
	},
	{
		"id":         3645,
		"first_name": "Sally",
		"last_name":  "Smith",
	},
	{
		"id":         1258,
		"first_name": "Jane",
		"last_name":  "Jones",
	},
	{
		"id":         2587,
		"first_name": "Peter",
		"last_name":  "Doe",
	},
}

func TestColumn(t *testing.T) {
	Convey("slicex.Column(columnTest, \"first_name\")", t, func() {
		actual := slicex.Column(columnTest, "first_name")
		expected := []interface{}{"John", "Sally", "Jane", "Peter"}
		So(actual, ShouldResemble, expected)
	})

	Convey("slicex.Column(columnTest, \"first_name\", \"id\")", t, func() {
		actual := slicex.Column(columnTest, "first_name", "id")
		expected := map[interface{}]interface{}{1858: "John", 3645: "Sally", 1258: "Jane", 2587: "Peter"}
		So(actual, ShouldResemble, expected)
	})
}

var inTest = []interface{}{1858, 3645, 1258, 2587}

func TestIn(t *testing.T) {
	Convey("Given a int slice", t, func() {
		Convey("When the first parameter is 3645", func() {
			Convey("Then the result should be true", func() {
				So(slicex.In(3645, inTest), ShouldBeTrue)
			})
		})
		Convey("When the first parameter is 2869", func() {
			Convey("Then the result should be false", func() {
				So(slicex.In(2869, inTest), ShouldBeFalse)
			})
		})
		Convey("When the first parameter is string \"2587\"", func() {
			Convey("Then the result should be false", func() {
				So(slicex.In("2587", inTest), ShouldBeFalse)
			})
		})
	})
}

func TestDiff(t *testing.T) {
	Convey("Given two slice", t, func() {
		slice1 := []interface{}{"green", "red", "blue"}
		slice2 := []interface{}{"green", "yellow", "blue", 2}
		Convey("When the comparison is done", func() {
			Convey("Then the result should be resemble", func() {
				expected := []interface{}{"red"}
				So(slicex.Diff(slice1, slice2), ShouldResemble, expected)
			})
		})
	})
}

func TestIntersect(t *testing.T) {
	Convey("Given two slice", t, func() {
		slice1 := []interface{}{"green", "red", "blue"}
		slice2 := []interface{}{"green", "yellow", "blue", 2}
		Convey("When the comparison is done", func() {
			Convey("Then the result should be resemble", func() {
				expected := []interface{}{"green", "blue"}
				So(slicex.Intersect(slice1, slice2), ShouldResemble, expected)
			})
		})
	})
}
