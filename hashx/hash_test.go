package hashx_test

import (
	"testing"

	"github.com/mzccddkk/gdk/hashx"
	. "github.com/smartystreets/goconvey/convey"
)

var md5Test = map[string]string{
	"d41d8cd98f00b204e9800998ecf8427e": "",
	"0cc175b9c0f1b6a831c399e269772661": "a",
	"187ef4436122d1cc2f40dc2b92f0eba0": "ab",
	"900150983cd24fb0d6963f7d28e17f72": "abc",
	"e2fc714c4727ee9395f324cd2e7f331f": "abcd",
	"ab56b4d92b40713acc5af89985d4b786": "abcde",
	"e80b5017098950fc58aad83c8c14978e": "abcdef",
}

func TestMd5(t *testing.T) {
	for k, v := range md5Test {
		Convey("hashx.Md5(\""+v+"\") should equal "+k, t, func() {
			So(hashx.Md5(v), ShouldEqual, k)
		})
	}
}

var sha1Test = map[string]string{
	"da39a3ee5e6b4b0d3255bfef95601890afd80709": "",
	"86f7e437faa5a7fce15d1ddcb9eaeaea377667b8": "a",
	"da23614e02469a0d7c7bd1bdab5c9c474b1904dc": "ab",
	"a9993e364706816aba3e25717850c26c9cd0d89d": "abc",
	"81fe8bfe87576c3ecb22426f8e57847382917acf": "abcd",
	"03de6c570bfe24bfc328ccd7ca46b76eadaf4334": "abcde",
	"1f8ac10f23c5b5bc1167bda84b833e5c057a77d2": "abcdef",
}

func TestSha1(t *testing.T) {
	for k, v := range sha1Test {
		Convey("hashx.Sha1(\""+v+"\") should equal "+k, t, func() {
			So(hashx.Sha1(v), ShouldEqual, k)
		})
	}
}

var sha256Test = map[string]string{
	"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855": "",
	"ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb": "a",
	"fb8e20fc2e4c3f248c60c39bd652f3c1347298bb977b8b4d5903b85055620603": "ab",
	"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad": "abc",
	"88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589": "abcd",
	"36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c": "abcde",
	"bef57ec7f53a6d40beb640a780a639c83bc29ac8a9816f1fc6c5c6dcd93c4721": "abcdef",
}

func TestSha256(t *testing.T) {
	for k, v := range sha256Test {
		Convey("hashx.Sha256(\""+v+"\") should equal "+k, t, func() {
			So(hashx.Sha256(v), ShouldEqual, k)
		})
	}
}
