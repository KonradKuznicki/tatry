package lib

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
)

func TestSafeFileNameFixture(t *testing.T) {
	gunit.Run(new(SafeFileNameFixture), t)
}

type SafeFileNameFixture struct {
	*gunit.Fixture
}

func (this *SafeFileNameFixture) Setup() {
}

func (this *SafeFileNameFixture) TestFileName() {
	this.So(SafeFileName("http://asdf.fff.com/aś/sdf1234.jpe"), should.Equal, "http_asdf.fff.com_a_sdf1234.jpe")
}
