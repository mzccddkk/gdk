package hashx_test

import (
	"fmt"
	"testing"

	"github.com/mzccddkk/gdk/hashx"
	. "github.com/smartystreets/goconvey/convey"
)

type hashTest struct {
	hash string
	in   string
	out  string
}

var hashTests = []hashTest{
	{
		"md5",
		"Sample #1",
		"0c5dd21f4e300bf13d57cbd1a7629822",
	},
	{
		"md5",
		"Sample #2",
		"f7a5c31ed2326135be3cef0aee334cdf",
	},
	{
		"md5",
		"Sample #3",
		"a03dcfd6a77d4f4997202e0ab75561ed",
	},
	{
		"sha1",
		"Sample #1",
		"9cbd5edf990b90162d2b2568f546d7417430c695",
	},
	{
		"sha1",
		"Sample #2",
		"0817d04b9ba6253549a57e5c90051eb73c79352a",
	},
	{
		"sha1",
		"Sample #3",
		"86882a85ee40b2b71a40da4040df9be08a83c351",
	},
	{
		"sha256",
		"Sample #1",
		"b733ead66b732ccb4a64e010a396642764c6a45a376ce5220e89dcbba8574f1c",
	},
	{
		"sha256",
		"Sample #2",
		"920c288bff87f8355c385c5afea5c01abe65376a5cf8a0553124a1e27d9e640f",
	},
	{
		"sha256",
		"Sample #3",
		"5bda2a35727fd63019d4193dc9b21771daa70f4c77748381a5c9361f9ff9e667",
	},
	{
		"sha512",
		"Sample #1",
		"bfa7330b2402288fe01662a5f064b63302d81dd217c9b383144854942ff2cbc86b633c3de8dd9dffa270316d36140a334a2449c555524f9f76351327e3c53134",
	},
	{
		"sha512",
		"Sample #2",
		"c450392be555deb09d257757d58e539f331d8f281e254b57fd03ddd7aea41d0f6104ca617a3ccfb8099e6e6dd9a961e7d944d066196f81ab12b41476c2f86755",
	},
	{
		"sha512",
		"Sample #3",
		"3f9f4ffd3a0fda20f785e49cd76cd238ba8f08cc8b6d0c77fb7f7923d3a52a2fb5e9b70994886d2197c15cab08ec568d06b19fcda946d2397a90fcc4b772426e",
	},
}

func TestHash(t *testing.T) {
	for _, v := range hashTests {
		if v.hash == "md5" {
			i := fmt.Sprintf("hashx.Md5(\"%s\") should equal %s", v.in, v.out)
			Convey(i, t, func() {
				So(hashx.Md5(v.in), ShouldEqual, v.out)
			})
		}
		if v.hash == "sha1" {
			i := fmt.Sprintf("hashx.Sha1(\"%s\") should equal %s", v.in, v.out)
			Convey(i, t, func() {
				So(hashx.Sha1(v.in), ShouldEqual, v.out)
			})
		}
		if v.hash == "sha256" {
			i := fmt.Sprintf("hashx.Sha256(\"%s\") should equal %s", v.in, v.out)
			Convey(i, t, func() {
				So(hashx.Sha256(v.in), ShouldEqual, v.out)
			})
		}
		if v.hash == "sha512" {
			i := fmt.Sprintf("hashx.Sha512(\"%s\") should equal %s", v.in, v.out)
			Convey(i, t, func() {
				So(hashx.Sha512(v.in), ShouldEqual, v.out)
			})
		}
	}
}
