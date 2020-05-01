package lib

import (
	"log"
	"sync"
	"time"
)

type Updater interface {
	Process(in chan *Envelope)
}

type Downloader interface {
	Process(in chan *Envelope, out chan *Envelope)
}

type DownloaderPipeline struct {
	downloadChannel chan *Envelope
	updateChannel   chan *Envelope
	waitGroup       sync.WaitGroup

	cameras             []string
	downloadConcurrency int
	downloader          Downloader
	updater             Updater
}

func (this *DownloaderPipeline) Process(tickChannel chan time.Time) {

	go this.dispatchImageURL(tickChannel, this.downloadChannel)

	this.dispatchConcurrentDownloaders()

	this.updater.Process(this.updateChannel)
}

func (this *DownloaderPipeline) dispatchConcurrentDownloaders() {
	for i := 0; i < this.downloadConcurrency; i++ {
		this.waitGroup.Add(1)
		go func() {
			this.downloader.Process(this.downloadChannel, this.updateChannel)
			this.waitGroup.Done()
		}()
	}
	go func() {
		this.waitGroup.Wait()
		close(this.updateChannel)
	}()
}

func (this *DownloaderPipeline) dispatchImageURL(tickChannel chan time.Time, downloadChannel chan *Envelope) {
	for range tickChannel {
		for _, cam := range this.cameras {
			log.Println("queueing refresh", cam)
			downloadChannel <- &Envelope{
				Time: time.Now(),
				Cam:  cam,
			}
		}
	}
	close(downloadChannel)
}

func NewPipeline(cameras []string, downloadConcurrency int, downloader Downloader, updater Updater) *DownloaderPipeline {
	return &DownloaderPipeline{
		downloadChannel: make(chan *Envelope, 100),
		updateChannel:   make(chan *Envelope, 100),

		cameras:             cameras,
		downloadConcurrency: downloadConcurrency,
		downloader:          downloader,
		updater:             updater,
	}
}
