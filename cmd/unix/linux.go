package main

// opio 는 웹에서 opio:// 로 시작하는 URL을 인식하고,
// opio 명령어에 URL 값을 넘겨 관련 응용프로그램을 실행하는 프로그램이다.

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// Linux 함수는 URL로 전달받은 문자를 실행하는 함수이다.
func Linux(scape string) {
	switch strings.ToLower(filepath.Ext(scape)) {
	case ".nk":
		// 회사 셋팅에서 사용자 .bashrc에 보면 IP팀이 umask 0002라고 설정해놓았다.
		// opio를 통해서 뉴크를 실행하기 때문에 opio 도 umask 설정이 필요하다.
		// 이렇게 설정이되어야 뉴크실행후 뉴크가 만드는 폴더에 대해서 권한문제가 발생하지 않는다.
		syscall.Umask(0002) // 윈도우는 지원 안함.
		err := exec.Command("gnome-terminal", "-x", "nuke", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".jpg", ".png", ".exr", ".tga", ".psd", ".dpx", ".tif":
		imglist := []string{}
		imglist = append(imglist, strings.Split(scape, ";")...)
		imgext := []string{".jpg", ".png", ".exr", ".tga", ".psd", ".dpx", ".tif"}
		imagelist := []string{}
		for _, img := range imglist {
			for _, ext := range imgext {
				if !strings.Contains(img, ext) {
					continue
				}
				imagelist = append(imagelist, img)
			}
		}
		os.Setenv("RV_SUPPORT_PATH", "/rv/supportPath")                       // 회사 RV 파이프라인툴을 로딩하기 위해서 필요하다.
		os.Setenv("PKG_CONFIG_PATH", "/opencv/v3.2.0/lib/pkgconfig")          // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("LD_LIBRARY_PATH", "/opencv/v3.2.0/lib")                    // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("PYTHONPATH", "/opencv/v3.2.0/lib/python2.7/site-packages") // import cv2, import numpy 를 로딩하기 위해서 필요하다.
		os.Setenv("OCIO", "/ocio/aces_1.0.3/config.ocio")                     // RV 실행시 OCIO를 로딩하기 위해서 필요하다.
		err := exec.Command(rvLinuxAppPath, imagelist...).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".rv":
		os.Setenv("RV_ENABLE_MIO_FFMPEG", "1")                                     // Prores코덱을 위해서 활성화 한다.
		os.Setenv("RV_SUPPORT_PATH", "/rv/supportPath")                            // 회사 RV 파이프라인툴을 로딩하기 위해서 필요하다.
		os.Setenv("PKG_CONFIG_PATH", "/Tool/opencv/v3.2.0/lib/pkgconfig")          // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("LD_LIBRARY_PATH", "/Tool/opencv/v3.2.0/lib")                    // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("PYTHONPATH", "/Tool/opencv/v3.2.0/lib/python2.7/site-packages") // import cv2, import numpy 를 로딩하기 위해서 필요하다.
		os.Setenv("OCIO", "/Tool/ocio/aces_1.0.3/config.ocio")                     // RV 실행시 OCIO를 로딩하기 위해서 필요하다.
		err := exec.Command(rvLinuxAppPath, scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".mov":
		os.Setenv("RV_ENABLE_MIO_FFMPEG", "1")                                     // Prores코덱을 위해서 활성화 한다.
		os.Setenv("RV_SUPPORT_PATH", "/rv/supportPath")                            // 회사 RV 파이프라인툴을 로딩하기 위해서 필요하다.
		os.Setenv("PKG_CONFIG_PATH", "/Tool/opencv/v3.2.0/lib/pkgconfig")          // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("LD_LIBRARY_PATH", "/Tool/opencv/v3.2.0/lib")                    // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("PYTHONPATH", "/Tool/opencv/v3.2.0/lib/python2.7/site-packages") // import cv2, import numpy 를 로딩하기 위해서 필요하다.
		os.Setenv("OCIO", "/Tool/ocio/aces_1.0.3/config.ocio")                     // RV 실행시 OCIO를 로딩하기 위해서 필요하다.
		playlist := []string{}
		playlist = append(playlist, strings.Split(scape, ";")...)

		// 플레이 리스트를 받아서 입체 체크를 한다.
		movlist := []string{}
		isStereo := false
		for _, mov := range playlist {
			cmdlist, hasStereo := ToRvStereo(mov)
			if !hasStereo {
				movlist = append(movlist, mov)
				continue
			}
			// RV는 left, right 미디어를 같은 그룹을 묶을 때 "[,]"를 사용한다.
			movlist = append(movlist, "[")
			movlist = append(movlist, cmdlist...)
			movlist = append(movlist, "]")
			isStereo = true
		}
		if isStereo {
			// RV에서 입체 재생을 위해서는 옵션 마지막에 "-stereo scanline" 옵션 필요함.
			movlist = append(movlist, "-stereo")
			movlist = append(movlist, "scanline")
		}
		err := exec.Command(rvLinuxAppPath, movlist...).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".mp4":
		os.Setenv("RV_ENABLE_MIO_FFMPEG", "1")                                // Prores코덱을 위해서 활성화 한다.
		os.Setenv("RV_SUPPORT_PATH", "/rv/supportPath")                       // 회사 RV 파이프라인툴을 로딩하기 위해서 필요하다.
		os.Setenv("PKG_CONFIG_PATH", "/opencv/v3.2.0/lib/pkgconfig")          // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("LD_LIBRARY_PATH", "/opencv/v3.2.0/lib")                    // RV플러그인중 OpenCV를 로딩하기 위해서 필요함.
		os.Setenv("PYTHONPATH", "/opencv/v3.2.0/lib/python2.7/site-packages") // import cv2, import numpy 를 로딩하기 위해서 필요하다.
		os.Setenv("OCIO", "/ocio/aces_1.0.3/config.ocio")                     // RV 실행시 OCIO를 로딩하기 위해서 필요하다.
		err := exec.Command(rvLinuxAppPath, scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".avi", ".mkv":
		err := exec.Command("/usr/bin/vlc", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".ttf": // 폰트는 폰트브라우저로 연다.
		err := exec.Command("/usr/bin/gnome-font-viewer", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".pdf":
		err := exec.Command("/usr/bin/evince", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".blend":
		err := exec.Command("/Applications/Linux/blender/blender-2.75a-linux-glibc211-x86_64/blender", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".obj":
		err := exec.Command("/Applications/Linux/blender/blender-2.75a-linux-glibc211-x86_64/blender", "--python", "/blender/python/loadobj.py", "--", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".sh":
		err := exec.Command("mate-terminal", "-x", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".hip":
		syscall.Umask(0002) // 윈도우는 지원 안함.
		err := exec.Command("mate-terminal", "-x", "h", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".ma", ".mb":
		syscall.Umask(0002) // 윈도우는 지원 안함.
		err := exec.Command("mate-terminal", "-x", "m", "-f", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".usd", ".usda", ".usdc", ".usdz", ".abc":
		syscall.Umask(0002) // 윈도우는 지원 안함.
		err := exec.Command("mate-terminal", "-x", "uview", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".3de":
		syscall.Umask(0002) // 윈도우는 지원 안함.
		err := exec.Command("mate-terminal", "-x", "e", "-open", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".katana":
		syscall.Umask(0002) // 윈도우는 지원 안함.
		err := exec.Command("mate-terminal", "-x", "katana", "--asset", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	case ".project":
		syscall.Umask(0002) // 윈도우는 지원 안함.
		// clarisse는 옵션이 없이 파일로 캐치하지만, 편의성을 위해 커맨드명령이 있을때 실행한다.
		err := exec.Command("mate-terminal", "-x", "clarisse", "-config_file", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	default:
		err := exec.Command("nautilus", scape).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
