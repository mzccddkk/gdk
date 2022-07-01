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
