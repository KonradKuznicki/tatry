package lib

import (
	"log"
)

type JustSave struct {
	fs FS
}

func NewJustSave(fs FS) *JustSave {
	return &JustSave{fs: fs}
}

func (this *JustSave) Process(in chan *Envelope) {
	log.Println("writer listening")
	for envelope := range in {
		err := this.fs.Write(envelope.Cam, envelope.Out)

		if err != nil {
			log.Printf("%s", err)
		}

	}
}
