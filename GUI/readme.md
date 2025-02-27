Install MingW 64
```
https://www.msys2.org/
```
Run MSYS2 terminal
```
pacman -Syu
pacman -Su
pacman -S mingw-w64-ucrt-x86_64-{go,gcc,pkg-config}
pacman -S mingw-w64-ucrt-x86_64-qt5-base  # For Qt 5
pacman -S mingw-w64-ucrt-x86_64-qt6-base  # For Qt 6
export GOROOT=/ucrt64/lib/go
export PATH=$GOROOT/bin:$PATH
go build -ldflags "-s -w -H windowsgui"
```
Check from VS Code
```
pacman -Q | grep go
```

```
https://github.com/mappu/miqt
```

```bash
miqt-uic -InFile login.ui -OutFile login_ui.go
```


thererecip
```
go get -u github.com/therecipe/qt/cmd/...
```
```
go install github.com/therecipe/qt/cmd/...
or
go install github.com/stephenlyu/goqtuic@latest
goqtuic -ui-file login.ui -go-ui-dir .
```
check
```
ls $env:GOPATH\bin
```
