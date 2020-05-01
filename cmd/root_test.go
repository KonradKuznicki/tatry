package cmd

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

func TestRootCMDFixture(t *testing.T) {
	gunit.Run(new(RootCMDFixture), t)
}

type RootCMDFixture struct {
	*gunit.Fixture
	testTmpDir string
}

func (this *RootCMDFixture) Setup() {
	var err error
	this.testTmpDir, err = ioutil.TempDir("./", "test-tmp")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.SetArgs([]string{"--cache-dir", this.testTmpDir, "-u", images[0] + "," + images[1], "-r", "1s", "-p", "2", "-i", "20s"})
}

func (this *RootCMDFixture) Teardown() {
	if err := os.RemoveAll(this.testTmpDir); err != nil {
		log.Fatal(err)
	}
}

func (this *RootCMDFixture) LongTestRun() {

	go rootCmd.Execute()

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
