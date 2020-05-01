package os_background_interfaces

import (
	"log"
	"os/exec"
	"tatry/lib"
)

type GnomeUbuntu struct {
	fs lib.FS
}

func NewGnomeUbuntu(fs lib.FS) *GnomeUbuntu {
	return &GnomeUbuntu{fs: fs}
}

func (this *GnomeUbuntu) Set(path string) {

	out, err := exec.Command(
		"gsettings",
		"set",
		"org.gnome.desktop.background",
		"picture-uri",
		"file://"+this.fs.FullPath(path)).Output()

	if err != nil {
		log.Printf("%s", err)
		log.Println(string(out))
	}

}

func GnomeUbuntuCurrentBackground() string {

	out, err := exec.Command(
		"gsettings",
		"get",
		"org.gnome.desktop.background",
		"picture-uri").Output()

	if err != nil {
		log.Printf("%s", err)
		log.Println(string(out))
	}

	return string(out[1 : len(out)-2])
}
