Install MingW 64
```
https://www.msys2.org/
```
Add Env
```
C:\msys64\ucrt64\bin
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
