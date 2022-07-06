package hmacx_test

import (
	"fmt"
	"testing"

	"github.com/mzccddkk/gdk/hmacx"
	. "github.com/smartystreets/goconvey/convey"
)

type hmacTest struct {
	hash string
	in   string
	key  string
	out  string
}

var hmacTests = []hmacTest{
	{
		"md5",
		"Sample #1",
		"secret1",
		"51e69572c10fe283f4667e359339232d",
	},
	{
		"md5",
		"Sample #2",
		"secret2",
		"b8286dfa290a1c4197b245438c5b6da4",
	},
	{
		"md5",
		"Sample #3",
		"secret3",
		"4586bdf09385892fd308663105498d37",
	},
	{
		"sha1",
		"Sample #1",
		"secret1",
		"8247d7c2d71c1c08c135f6b3a77e6fd2edb3d38d",
	},
	{
		"sha1",
		"Sample #2",
		"secret2",
		"f0119c91f2c8dda90a812d65f275c0e29d2da0b7",
	},
	{
		"sha1",
		"Sample #3",
		"secret3",
		"b3440c6b8f37b9bf6bc998029d05f4f3045fa0a2",
	},
	{
		"sha256",
		"Sample #1",
		"secret1",
		"e7a08cdedad6a905bf9d1b9334d2726c984e0b8b62e2f11ddfb18a20acde8437",
	},
	{
		"sha256",
		"Sample #2",
		"secret2",
		"06c772bde71d7dffd580c2947fa29cb0a560925da3f05ae9d2fb5a803df152a1",
	},
	{
		"sha256",
		"Sample #3",
		"secret3",
		"a2857672e4e04297d33b0b3d31e957613902e30143ee29e23e0134a51825d76b",
	},
}

func TestHMAC(t *testing.T) {
	for _, v := range hmacTests {
		if v.hash == "md5" {
			i := fmt.Sprintf("hmacx.Md5(\"%s\", \"%s\") should equal %s", v.in, v.key, v.out)
			Convey(i, t, func() {
				So(hmacx.Md5(v.in, v.key), ShouldEqual, v.out)
			})
		}
		if v.hash == "sha1" {
			i := fmt.Sprintf("hmacx.Sha1(\"%s\", \"%s\") should equal %s", v.in, v.key, v.out)
			Convey(i, t, func() {
				So(hmacx.Sha1(v.in, v.key), ShouldEqual, v.out)
			})
		}
		if v.hash == "sha256" {
			i := fmt.Sprintf("hmacx.Sha256(\"%s\", \"%s\") should equal %s", v.in, v.key, v.out)
			Convey(i, t, func() {
				So(hmacx.Sha256(v.in, v.key), ShouldEqual, v.out)
			})
		}
	}
}
