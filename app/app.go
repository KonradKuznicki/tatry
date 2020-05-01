package app

import (
	"tatry/lib"
	"tatry/os_background_interfaces"
	"time"
)

type Config struct {
	CacheLocation              string        `mapstructure:"cache-dir"`
	Cams                       []string      `mapstructure:"URLs"`
	DownloadConcurrency        int           `mapstructure:"download-parallelism"`
	DownloadInterval           time.Duration `mapstructure:"download-interval"`
	BackgroundRotationInterval time.Duration `mapstructure:"background-rotation"`
}

type App struct {
	config *Config
	fs     lib.FS
}

func NewApp(config *Config) *App {
	return &App{config: config, fs: lib.NewLocalFS(config.CacheLocation)}
}

func (this *App) Run() {

	go this.startBackgroundDownloader()

	go this.startBackgroundRotator()

	this.hangForEver()
}

func (this *App) hangForEver() {
	var noChan chan interface{}
	noChan <- nil
}

func (this *App) startBackgroundRotator() {
	backgroundSetter := os_background_interfaces.NewGnomeUbuntu(this.fs)
	backgroundRotator := lib.NewBackgroundRotation(this.config.Cams, backgroundSetter)
	backgroundRotator.Process(tickChannel(this.config.BackgroundRotationInterval))
}

func (this *App) startBackgroundDownloader() {
	downloader := lib.NewDownloadImage()
	updater := lib.NewJustSave(this.fs)
	pipeline := lib.NewPipeline(this.config.Cams, this.config.DownloadConcurrency, downloader, updater)
	pipeline.Process(tickChannel(this.config.DownloadInterval))
}

func tickChannel(interval time.Duration) chan time.Time {
	ticks := make(chan time.Time, 10)

	go func() {
		ticks <- time.Now()
		for {
			time.Sleep(interval)
			ticks <- time.Now()
		}
	}()

	return ticks
}
