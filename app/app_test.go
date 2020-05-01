package app

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"io/ioutil"
	"log"
	"os"
	"tatry/os_background_interfaces"
	"testing"
	"time"
)

func TestIntegrationFixture(t *testing.T) {
	gunit.Run(new(AppFixture), t)
}

type AppFixture struct {
	*gunit.Fixture
	app        *App
	testTmpDir string
}

func (this *AppFixture) Setup() {
	var err error
	this.testTmpDir, err = ioutil.TempDir("./", "test-tmp")
	if err != nil {
		log.Fatal(err)
	}

	c := &Config{
		CacheLocation:              this.testTmpDir,
		Cams:                       images,
		DownloadConcurrency:        3,
		DownloadInterval:           time.Second * 20,
		BackgroundRotationInterval: time.Second,
	}

	this.app = NewApp(c)
}

func (this *AppFixture) Teardown() {
	if err := os.RemoveAll(this.testTmpDir); err != nil {
		log.Fatal(err)
	}
}

func (this *AppFixture) LongTestRun() {

	go this.app.Run()

	time.Sleep(time.Second * 3)

	content, err := ioutil.ReadFile(os_background_interfaces.GnomeUbuntuCurrentBackground()[7:])

	this.So(err, should.BeNil)
	this.So(content, should.NotBeEmpty)
}

/////////////////////////////////////////////////////////

var images = []string{
	"http://pogoda.topr.pl/download/current/kwgs.jpeg",
	"http://pogoda.topr.pl/download/current/hala.jpeg",
}
