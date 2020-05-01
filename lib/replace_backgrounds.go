package lib

import "time"

type BackgroundRotation struct {
	cams              []string
	currentBackground int
	backgroundSetter  BackgroundSetter
}

func NewBackgroundRotation(cams []string, backgroundSetter BackgroundSetter) *BackgroundRotation {
	return &BackgroundRotation{cams: cams, backgroundSetter: backgroundSetter}
}

func (this *BackgroundRotation) Process(refresh chan time.Time) {
	for range refresh {
		this.changeBackground()
	}
}

func (this *BackgroundRotation) changeBackground() {
	this.backgroundSetter.Set(this.rotatePath())
}

func (this *BackgroundRotation) rotatePath() string {
	this.currentBackground++
	if this.currentBackground >= len(this.cams) {
		this.currentBackground = 0
	}

	return this.cams[this.currentBackground]
}
