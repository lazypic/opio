package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Windows 함수는 URL로 전달받은 문자를 실행하는 함수이다.
func Windows(scape string) {
	switch strings.ToLower(filepath.Ext(scape)) {
	case ".mov":
		os.Setenv("RV_SUPPORT_PATH", "") // 회사 RV 파이프라인툴을 로딩하기 위해서 필요하다.
		if strings.Contains(scape, ";") {
			var movlist []string
			pathlist := strings.Split(scape, ";")
			movlist = append(movlist, pathlist...)
			err := exec.Command(rvWindowsAppPath, movlist...).Run()
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		err := exec.Command(rvWindowsAppPath, scape).Run()
		if err != nil {
			log.Fatal(err)
		}
		return
	case ".rv":
		os.Setenv("RV_SUPPORT_PATH", "") // 회사 RV 파이프라인툴을 로딩하기 위해서 필요하다.
		err := exec.Command(rvWindowsAppPath, scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	default:
		err := exec.Command("cmd", "/C", "start", "", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
