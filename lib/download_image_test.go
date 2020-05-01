package lib

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
)

func TestDownloadImage(t *testing.T) {
	gunit.RunSequential(new(DownloadImageFixture), t)
}

type DownloadImageFixture struct {
	*gunit.Fixture
	downloader *DownloadImage
}

func (this *DownloadImageFixture) Setup() {
	this.downloader = NewDownloadImage()
}

func (this *DownloadImageFixture) TestProcessDownloads() {
	in := make(chan *Envelope, 10)
	out := make(chan *Envelope, 10)

	brokenImage := &Envelope{Cam: "kwgs"}
	okImage := &Envelope{Cam: "http://pogoda.topr.pl/download/current/hala.jpeg"}
	in <- okImage
	in <- brokenImage

	close(in)

	this.downloader.Process(in, out)

	this.So(<-out, should.Equal, okImage)
	this.So(brokenImage.err, should.NotBeNil)

}
