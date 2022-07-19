package murmur3_test

import (
	"testing"

	"github.com/mzccddkk/gdk/murmur3"
	. "github.com/smartystreets/goconvey/convey"
)

type murmurTest struct {
	str  string
	seed uint32

	ha  uint32
	has uint32

	h1c  uint32
	h2c  uint32
	h3c  uint32
	h4c  uint32
	h1cs uint32
	h2cs uint32
	h3cs uint32
	h4cs uint32

	h1f  uint64
	h2f  uint64
	h1fs uint64
	h2fs uint64
}

var murmurTests = []murmurTest{
	{
		"Hello World",
		123456,
		427197390,
		4125902693,
		2462775914,
		1499806821,
		1090519024,
		188476316,
		3637336348,
		2044961496,
		2622412725,
		874480687,
		1901405986810282715,
		9504319040210908199,
		1440668659686177044,
		7465135666505861501,
	},
	{
		"The quick brown fox jumps over the lazy dog",
		123456,
		776992547,
		3170709897,
		871072587,
		2574835359,
		1926730536,
		3913369451,
		2082714780,
		3697772658,
		2850304761,
		156758268,
		16378391709484522348,
		8809951995912426311,
		516196728027175252,
		7547558758340691378,
	},
}

func TestMurmur3(t *testing.T) {
	for _, v := range murmurTests {
		Convey("TestMurmur3", t, func() {
			ha := murmur3.SumA(v.str)
			So(ha, ShouldEqual, v.ha)

			has := murmur3.SumAWithSeed(v.str, v.seed)
			So(has, ShouldEqual, v.has)

			h1c, h2c, h3c, h4c := murmur3.SumC(v.str)
			So(h1c, ShouldEqual, v.h1c)
			So(h2c, ShouldEqual, v.h2c)
			So(h3c, ShouldEqual, v.h3c)
			So(h4c, ShouldEqual, v.h4c)

			h1cs, h2cs, h3cs, h4cs := murmur3.SumCWithSeed(v.str, v.seed)
			So(h1cs, ShouldEqual, v.h1cs)
			So(h2cs, ShouldEqual, v.h2cs)
			So(h3cs, ShouldEqual, v.h3cs)
			So(h4cs, ShouldEqual, v.h4cs)

			h1f, h2f := murmur3.SumF(v.str)
			So(h1f, ShouldEqual, v.h1f)
			So(h2f, ShouldEqual, v.h2f)

			h1fs, h2fs := murmur3.SumFWithSeed(v.str, v.seed)
			So(h1fs, ShouldEqual, v.h1fs)
			So(h2fs, ShouldEqual, v.h2fs)
		})
	}
}
