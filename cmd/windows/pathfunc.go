package main

import (
	"log"
	"os/user"
	"strings"
)

// Home2Abspath 함수는 ~ 문자로 경로가 시작하면 물리적인 경로로 바꾸어준다.
func Home2Abspath(p string) string {
	if !strings.HasPrefix(p, "~") {
		return p
	}
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + strings.TrimPrefix(p, "~")
}

// Win2lin 함수는 윈도우즈 경로를 리눅스 경로로 바꾼다. 만약, 변환되지 않으면 패스를 그대로 출력한다.
func Win2lin(path string) string {
	if strings.HasPrefix(path, "W:\\") {
		return "/show/" + strings.Replace(path[3:], "\\", "/", len(path[3:]))
	} else if strings.HasPrefix(path, "/show/") {
		return path
	} else if strings.HasPrefix(path, "/lustre") { // lustre, lustre2, lustre3, lustre4 로 시작할 때..
		return path
	} else if strings.HasPrefix(path, "\\\\10.0.200.100\\show_") {
		return "/show/" + strings.Replace(path[20:], "\\", "/", len(path[20:]))
	} else if strings.HasPrefix(path, "\\\\10.0.200.100\\lustre_Digitalidea_source\\") {
		return "/lustre2/Digitalidea_source/" + strings.Replace(path[41:], "\\", "/", len(path[41:]))
	} else {
		return path
	}
}

//Lin2win 함수는 리눅스 경로를 윈도우즈 경로로 바꾼다.
func Lin2win(path string) string {
	if strings.HasPrefix(path, "/lustre2/Digitalidea_source/flib") { //flib
		return "\\\\10.0.200.100\\lustre_Digitalidea_source\\flib" + strings.Replace(path[32:], "/", "\\", len(path[32:]))
	} else if strings.HasPrefix(path, "/lustre/Digitalidea_source/flib") { //flib
		return "\\\\10.0.200.100\\lustre_Digitalidea_source\\flib" + strings.Replace(path[31:], "/", "\\", len(path[31:]))
	} else if strings.HasPrefix(path, "/show") {
		return "\\\\10.0.200.100\\show_" + strings.Replace(path[6:], "/", "\\", len(path[6:]))
	} else if strings.HasPrefix(path, "/lustre/show") {
		return "\\\\10.0.200.100\\show_" + strings.Replace(path[13:], "/", "\\", len(path[13:]))
	} else if strings.HasPrefix(path, "/lustre2/show") {
		return "\\\\10.0.200.100\\show_" + strings.Replace(path[14:], "/", "\\", len(path[14:]))
	} else if strings.HasPrefix(path, "/lustre3/show") {
		return "\\\\10.0.200.100\\show_" + strings.Replace(path[14:], "/", "\\", len(path[14:]))
	} else if strings.HasPrefix(path, "/lustre4/show") {
		return "\\\\10.0.200.100\\show_" + strings.Replace(path[14:], "/", "\\", len(path[14:]))
	} else if strings.HasPrefix(path, "/MMHUB_nas01/show") {
		return "Z:\\show\\" + strings.Replace(path[18:], "/", "\\", len(path[18:]))
	} else {
		return path
	}
}

func Lin2wins(paths []string) []string {
	var results []string
	for _, path := range paths {
		results = append(results, Lin2win(path))
	}
	return results
}
