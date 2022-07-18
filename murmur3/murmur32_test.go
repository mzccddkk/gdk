package murmur3_test

import (
	"strconv"
	"testing"

	"github.com/mzccddkk/gdk/murmur3"
	. "github.com/smartystreets/goconvey/convey"
)

var sum32Test = map[uint32]string{
	1635893381: "abcdef",
	2665153934: "abcdef中",
	1919966354: "abcd中",
}

func TestSum32(t *testing.T) {
	for k, v := range sum32Test {
		Convey("murmur3.Sum32(\""+v+"\") should equal "+strconv.Itoa(int(k)), t, func() {
			So(murmur3.SumA(v), ShouldEqual, k)
		})
	}
}
