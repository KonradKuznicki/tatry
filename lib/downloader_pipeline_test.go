package lib

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	gunit.Run(new(PipelineFixture), t)
}

type PipelineFixture struct {
	*gunit.Fixture
	pipeline      *DownloaderPipeline
	spyDownloader *SpyDownloader
	spyUpdater    *SpyUpdater
}

func (this *PipelineFixture) Setup() {
	this.spyDownloader = &SpyDownloader{}
	this.spyUpdater = &SpyUpdater{}
	this.pipeline = NewPipeline([]string{"http://test/image.jpeg", "http://test/image.jpeg"}, 3, this.spyDownloader, this.spyUpdater)
}

func (this *PipelineFixture) TestPipeline() {
	tickChan := make(chan time.Time, 10)
	tickChan <- time.Now()
	tickChan <- time.Now()
	close(tickChan)

	this.pipeline.Process(tickChan)

	this.So(this.spyDownloader.processed, should.Equal, 4)
	this.So(this.spyUpdater.processed, should.Equal, 4)

}

//////////////////////////////////////////////////////////////////////

type SpyDownloader struct {
	processed int
}

func (this *SpyDownloader) Process(in chan *Envelope, out chan *Envelope) {
	for envelope := range in {
		this.processed++
		out <- envelope
	}
}

type SpyUpdater struct {
	processed int
}

func (this *SpyUpdater) Process(in chan *Envelope) {
	for range in {
		this.processed++
	}
}
