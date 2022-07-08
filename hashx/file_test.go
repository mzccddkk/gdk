package hashx_test

import (
	"testing"

	"github.com/mzccddkk/gdk/hashx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHashFile(t *testing.T) {
	Convey("Test Md5File hashx.Md5File(\"./test\")", t, func() {
		hashed, _ := hashx.Md5File("./test")
		So(hashed, ShouldEqual, "a5890ace30a3e84d9118196c161aeec2")
	})
	Convey("Test Md5File  hashx.Md5File(\"./test_not_found\")", t, func() {
		hashed, _ := hashx.Md5File("./test_not_found")
		So(hashed, ShouldEqual, "")
	})

	Convey("Test Sha1File hashx.Sha1File(\"./test\")", t, func() {
		hashed, _ := hashx.Sha1File("./test")
		So(hashed, ShouldEqual, "5d03965084a5db13c178cbb1ffc120b360353685")
	})
	Convey("Test Sha1File  hashx.Sha1File(\"./test_not_found\")", t, func() {
		hashed, _ := hashx.Sha1File("./test_not_found")
		So(hashed, ShouldEqual, "")
	})

	Convey("Test Sha256File hashx.Sha256File(\"./test\")", t, func() {
		hashed, _ := hashx.Sha256File("./test")
		So(hashed, ShouldEqual, "5881707e54b0112f901bc83a1ffbacac8fab74ea46a6f706a3efc5f7d4c1c625")
	})
	Convey("Test Sha256File  hashx.Sha256File(\"./test_not_found\")", t, func() {
		hashed, _ := hashx.Sha256File("./test_not_found")
		So(hashed, ShouldEqual, "")
	})

	Convey("Test Sha512File hashx.Sha512File(\"./test\")", t, func() {
		hashed, _ := hashx.Sha512File("./test")
		So(hashed, ShouldEqual, "3a618d54b8dd96ad42eedeaeac753d7398254e8bb0027c3517d906532608abdadf8b372d5025f261aaa15cbd17229b1e1f98d6df60c91277a422addfd9b7cb4c")
	})
	Convey("Test Sha512File  hashx.Sha512File(\"./test_not_found\")", t, func() {
		hashed, _ := hashx.Sha512File("./test_not_found")
		So(hashed, ShouldEqual, "")
	})
}
