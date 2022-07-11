package hmacx_test

import (
	"testing"

	"github.com/mzccddkk/gdk/hmacx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHmacFile(t *testing.T) {
	Convey("Test Md5File hmacx.Md5File(\"../testdata/test\")", t, func() {
		hashed, _ := hmacx.Md5File("../testdata/test", "secret")
		So(hashed, ShouldEqual, "deba4bdf87356d276ea2910caa00f0a0")
	})
	Convey("Test Md5File  hmacx.Md5File(\"../testdata/test_not_found\")", t, func() {
		hashed, _ := hmacx.Md5File("../testdata/test_not_found", "secret")
		So(hashed, ShouldEqual, "")
	})

	Convey("Test Sha1File hmacx.Sha1File(\"../testdata/test\")", t, func() {
		hashed, _ := hmacx.Sha1File("../testdata/test", "secret")
		So(hashed, ShouldEqual, "7e851b6b6f3e3dc1b1a571cb8bd75e3ef8748c4b")
	})
	Convey("Test Sha1File  hmacx.Sha1File(\"../testdata/test_not_found\")", t, func() {
		hashed, _ := hmacx.Sha1File("../testdata/test_not_found", "secret")
		So(hashed, ShouldEqual, "")
	})

	Convey("Test Sha256File hmacx.Sha256File(\"../testdata/test\")", t, func() {
		hashed, _ := hmacx.Sha256File("../testdata/test", "secret")
		So(hashed, ShouldEqual, "21e9d91585d28e90966a4597412c7e0104ff2beb73e6f172873d012271868364")
	})
	Convey("Test Sha256File  hmacx.Sha256File(\"../testdata/test_not_found\")", t, func() {
		hashed, _ := hmacx.Sha256File("../testdata/test_not_found", "secret")
		So(hashed, ShouldEqual, "")
	})

	Convey("Test Sha512File hmacx.Sha512File(\"../testdata/test\")", t, func() {
		hashed, _ := hmacx.Sha512File("../testdata/test", "secret")
		So(hashed, ShouldEqual, "bf5f572ab3473a30261db61f047514c54b895551ac40e9a468d6f342f26102aa14816c23183473c2c874a0cfbdc7cb94287af01a0154167360d425e796d1d4ec")
	})
	Convey("Test Sha512File  hmacx.Sha512File(\"../testdata/test_not_found\")", t, func() {
		hashed, _ := hmacx.Sha512File("../testdata/test_not_found", "secret")
		So(hashed, ShouldEqual, "")
	})
}
