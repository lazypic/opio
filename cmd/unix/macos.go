package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// MacOS 함수는 URL로 전달받은 문자를 실행하는 함수이다.
func MacOS(scape string) {
	scape = Home2Abspath(scape)
	switch strings.ToLower(filepath.Ext(scape)) {
	case ".nk":
		os.Setenv("NUKE_PATH", Home2Abspath("~/nuke"))
		os.Setenv("NUKE_FONT_PATH", Home2Abspath("~/nuke/font"))
		os.Setenv("OCIO", Home2Abspath("~/OpenColorIO-Configs/aces_1.2/config.ocio"))
		// 맥은 인터넷 연결되 되어있을 가능성이 높다. 항상 논커머셜로 실행한다.
		err := exec.Command("/Applications/Nuke13.2v3/Nuke13.2v3.app/Contents/MacOS/Nuke13.2", "--nukex", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".mp4", ".mov", ".jpg": // DJV 를 사용한다.
		err := exec.Command("/Applications/DJV2.app/Contents/MacOS/DJV2", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".blend":
		os.Setenv("OCIO", Home2Abspath("~/OpenColorIO-Configs/aces_1.2/config.ocio"))
		err := exec.Command(Home2Abspath("~/app/blender2.83/blender.app/Contents/MacOS/Blender"), "--python", Home2Abspath("~/blender/init.py"), scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".kra":
		os.Setenv("OCIO", Home2Abspath("~/OpenColorIO-Configs/aces_1.2/config.ocio"))
		err := exec.Command("/Applications/krita.app/Contents/MacOS/krita", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".xcf":
		err := exec.Command("/Applications/GIMP-2.10.app/Contents/MacOS/gimp", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	default:
		// 일반적으로 .abc, .hwp, 키노트등의 포멧은 open 명령어로 잘 작동된다.
		err := exec.Command("open", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
