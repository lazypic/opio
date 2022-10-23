# OPIO

OpenPipelineIO에서 사용하는 웹 프로토콜입니다.
응용프로그램을 웹에서 실행하기 위해서 제작되었습니다.

### opio 설치

#### Windows10

- `C:\bin` 폴더를 생성합니다.
- 위 폴더에 `opio.exe`, `install_Windows.reg` 파일을 복사합니다.
- `install_Windows.reg` 파일을 더블클릭 합니다.

#### CentOS7
터미널을 열고 아래처럼 명령어를 타이핑 합니다.

```bash
$ tcsh install_CentOS7.sh
```

#### macOS
1. 다운로드 받은 파일을 압축풉니다.
1. `opio.app` 파일을 어플리케이션에 복사합니다.
1. opio 명령어는 ~/bin 폴더에 넣습니다.

### Test실행
터미널을 이용해서 실제로 opio가 잘 작동되는지 체크해볼 수 있습니다.

```bash
$ opio opio:///file/path/test.blend
```

### License
BSD 3-Clause License