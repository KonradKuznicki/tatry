package lib

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
	"time"
)

func TestReplaceBackgroundFixture(t *testing.T) {
	gunit.Run(new(ReplaceBackgroundFixture), t)
}

type ReplaceBackgroundFixture struct {
	*gunit.Fixture
	rotator                *BackgroundRotation
	spyBackgroundInterface *SpyBackgroundInterface
}

func (this *ReplaceBackgroundFixture) Setup() {
	this.spyBackgroundInterface = NewSpyBackgroundInterface()
	this.rotator = NewBackgroundRotation([]string{"img1", "img2"}, this.spyBackgroundInterface)
}

func (this *ReplaceBackgroundFixture) TestBackgroundRotation() {
	tickChan := this.deployTicks(4)

	this.rotator.Process(tickChan)

	this.So(this.spyBackgroundInterface.called, should.Equal, 4)
	this.So(this.spyBackgroundInterface.pathsSet, should.Resemble, []string{"img2", "img1", "img2", "img1"})
}

func (this *ReplaceBackgroundFixture) deployTicks(tickCount int) chan time.Time {
	tickChan := make(chan time.Time, tickCount)
	for i := 0; i < tickCount; i++ {
		tickChan <- time.Now()
	}

	close(tickChan)
	return tickChan
}

//////////////////////////////////////////////////////////////////////////

type SpyBackgroundInterface struct {
	called   int
	pathsSet []string
}

func NewSpyBackgroundInterface() *SpyBackgroundInterface {
	return &SpyBackgroundInterface{
		pathsSet: []string{},
	}
}

func (this *SpyBackgroundInterface) Set(path string) {
	this.called++
	this.pathsSet = append(this.pathsSet, path)
}
