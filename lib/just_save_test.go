package lib

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"testing"
	"time"
)

func TestJustSaveFixture(t *testing.T) {
	gunit.Run(new(JustSaveFixture), t)
}

type JustSaveFixture struct {
	*gunit.Fixture
	justSave *JustSave
	envelope *Envelope
	testChan chan *Envelope
	spyFS    *SpyFS
}

func (this *JustSaveFixture) Setup() {
	this.spyFS = &SpyFS{}
	this.justSave = NewJustSave(this.spyFS)
}

func (this *JustSaveFixture) TestJustSave() {

	this.deployEnvelope()

	this.justSave.Process(this.testChan)

	this.AsserImageSaved()
}

func (this *JustSaveFixture) deployEnvelope() {
	this.createEnvelope()

	this.testChan = make(chan *Envelope, 10)
	this.testChan <- this.envelope
	close(this.testChan)
}

func (this *JustSaveFixture) createEnvelope() {
	this.envelope = &Envelope{
		Cam:  "http://my.image/image.jpg",
		err:  nil,
		Out:  []byte("test"),
		Time: time.Now(),
	}
}

func (this *JustSaveFixture) AsserImageSaved() {
	this.So(this.spyFS.fileName, should.Equal, this.envelope.Cam)
	this.So(this.envelope.Out, should.Resemble, this.envelope.Out)
}

//////////////////////////////////////////////////////////////////////////

type SpyFS struct {
	fileName    string
	fileContent []byte
}

func (this *SpyFS) FullPath(name string) string {
	return "full/path/" + this.fileName
}

func (this *SpyFS) Write(name string, content []byte) error {
	this.fileName = name
	this.fileContent = content
	return nil
}

func (this *SpyFS) Read(name string) ([]byte, error) {
	return this.fileContent, nil
}
