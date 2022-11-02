package main

// opio 는 웹에서 opio:// 로 시작하는 URL을 인식하고,
// opio 명령어에 URL 값을 넘겨 관련 응용프로그램을 실행하는 프로그램이다.

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

const (
	rvWindowsAppPath = "C:\\Program Files\\ShotGrid\\RV-2022.0.0\\bin\\rv.exe"
	rvLinuxAppPath   = "/opt/rv-Linux-x86-64-7.2.0/bin/rv"
	rvMacosAppPath   = "/Applications/RV64.app/Contents/MacOS/RV64"
	protocol         = "opio://"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("opio: ")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stdout, "명령를 실행하기 위한 인수가 충분하지 않습니다.\n")
		os.Exit(1)
	}
	// opio 프로토콜이 올바르게 써져있는지 체크함.
	if !strings.HasPrefix(flag.Args()[0], protocol) {
		fmt.Fprintf(os.Stdout, "인수가 %s 로 시작하지 않습니다. 종료합니다.\n", protocol)
		os.Exit(1)
	}
	uri := strings.TrimPrefix(flag.Args()[0], protocol)
	// URI를 통해서 문자를 받기 때문에 %3A -> ":", %2F -> "/" 같은 문자가 섞일 수 있다.
	// 이러한 문자를 QueryUnescape 함수를 통해서 1차 정리한다.
	scape, err := url.QueryUnescape(uri)
	if err != nil {
		log.Fatal(err)
	}

	Windows(scape)
}
