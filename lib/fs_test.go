package lib

import (
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestLocalFSFixture(t *testing.T) {
	gunit.Run(new(LocalFSFixture), t)
}

type LocalFSFixture struct {
	*gunit.Fixture
	fs         FS
	testTmpDir string
}

func (this *LocalFSFixture) Setup() {
	var err error
	this.testTmpDir, err = ioutil.TempDir("./", "test-tmp")
	if err != nil {
		log.Fatal(err)
	}

	this.fs = NewLocalFS(this.testTmpDir)
}

func (this *LocalFSFixture) Teardown() {
	if err := os.RemoveAll("./" + this.testTmpDir); err != nil {
		log.Fatal(err)
	}
}

func (this *LocalFSFixture) TestOperation() {
	testName := "http://ttt.aaa/@#$%śś1234|\".jpg"
	testConent := []byte("asdf")

	_ = this.fs.Write(testName, testConent)
	readContent, _ := this.fs.Read(testName)

	this.Println(this.fs.FullPath(testName))
	checkFile, err := ioutil.ReadFile(this.testTmpDir + "/http_ttt.aaa_1234_.jpg")

	this.So(err, should.BeNil)
	this.So(testConent, should.Resemble, checkFile)
	this.So(testConent, should.Resemble, readContent)
}
