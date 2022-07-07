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
	{
		"sha512",
		"Sample #1",
		"secret1",
		"d3758cfb6edc076d699d2140113a73ebdfa50238c2bcb3ccc6ceb28d294aa35e2164fe7e49bdaedb94698d396d005aace954030c1265f58c07ccfd79a37bcf6d",
	},
	{
		"sha512",
		"Sample #2",
		"secret2",
		"357dec399a3c02ce7ad28dbb9c2ea3c4837ea92977f3ba612d809b368409a46f925ef2d9aaca42881a71e90717084b766c98b6d67d4ca7efe8aeb497667cab2d",
	},
	{
		"sha512",
		"Sample #3",
		"secret3",
		"8609145d118bb5209baaa1690426f237f6c142975841304f36074629ef15aa7cc6729a383164677e33f80fe63eb3c9373a24a97b8e4db794ef281c8e920120b9",
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
		if v.hash == "sha512" {
			i := fmt.Sprintf("hmacx.Sha512(\"%s\", \"%s\") should equal %s", v.in, v.key, v.out)
			Convey(i, t, func() {
				So(hmacx.Sha512(v.in, v.key), ShouldEqual, v.out)
			})
		}
	}
}
