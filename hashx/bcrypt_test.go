package hashx_test

import (
	"testing"

	"github.com/mzccddkk/gdk/hashx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBcrypt(t *testing.T) {
	Convey("Bcrypt 123456 with default cost should be verified", t, func() {
		hashedPassword := hashx.Bcrypt("123456")
		So(hashx.BcryptVerify(hashedPassword, "123456"), ShouldBeTrue)
	})
	Convey("Bcrypt 123456 with 32 cost should not be verified", t, func() {
		hashedPassword := hashx.BcryptWithCost("123456", 32)
		So(hashx.BcryptVerify(hashedPassword, "123456"), ShouldBeFalse)
	})
}
