package os_background_interfaces

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"tatry/lib"
	"testing"
)

func TestGnomeUbuntuFixture(t *testing.T) {
	gunit.Run(new(GnomeUbuntuFixture), t)
}

type GnomeUbuntuFixture struct {
	*gunit.Fixture
	fs     lib.FS
	setter lib.BackgroundSetter
}

func (this *GnomeUbuntuFixture) Setup() {
	this.fs = lib.NewLocalFS("../test-resources")
	this.setter = NewGnomeUbuntu(this.fs)
}

func (this *GnomeUbuntuFixture) LongTestBackgroundSet() {

	path := "mountains-1362605.jpg"
	this.setter.Set(path)

	currentBackground := GnomeUbuntuCurrentBackground()

	this.So("file://./test-resources/"+path, should.Equal, currentBackground)

}
