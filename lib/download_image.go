package lib

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Envelope struct {
	Cam  string
	err  error
	Out  []byte
	Time time.Time
}

type DownloadImage struct {
}

func NewDownloadImage() *DownloadImage {
	return &DownloadImage{}
}

func (this *DownloadImage) Process(in chan *Envelope, out chan *Envelope) {
	for envelope := range in {
		log.Println("downloading", envelope.Cam)
		buf, err := this.fetch(envelope.Cam)
		envelope.Out = buf
		envelope.err = err
		if err != nil {
			log.Println("download error", envelope.Cam, envelope.err)
		} else {
			log.Println("downloaded", envelope.Cam)
			out <- envelope
		}
	}
}

func (this *DownloadImage) fetch(cam string) ([]byte, error) {
	req, err := http.NewRequest("GET", cam, nil)
	if err != nil {
		return []byte{}, errors.Wrap(err, "couldn't create image request")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.Wrap(err, "couldn't download image")
	}
	defer func() {
		_ = resp.Body.Close() // silence error
	}()

	return ioutil.ReadAll(resp.Body)
}
