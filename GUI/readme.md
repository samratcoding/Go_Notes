Install MingW 64
```
https://www.msys2.org/
```
Add Env
```
C:\msys64\ucrt64\bin
# Install Go and C++ toolchains
pacman -S mingw-w64-ucrt-x86_64-{go,gcc,pkg-config}
export GOROOT=/ucrt64/lib/go # Needed only if this is the first time installing Go in MSYS2. Otherwise it would be automatically applied when opening a new Bash terminal.

# Install Qt
pacman -S mingw-w64-ucrt-x86_64-qt5-base # For Qt 5 (UCRT64 GCC toolchain)
pacman -S mingw-w64-ucrt-x86_64-qt6-base # For Qt 6 (UCRT64 GCC toolchain)
pacman -S mingw-w64-clang-x86_64-qt6-base # For Qt 6 (CLANG64 toolchain)

go build -ldflags "-s -w -H windowsgui"
C:\msys64\ucrt64\lib\go\bin
```
```
Manually Set Environment Variables (Recommended)

Press Win + R, type sysdm.cpl, and press Enter.
Go to Advanced → Environment Variables.
Under System Variables, check or add:
GOROOT → C:\Program Files\Go
GOPATH → C:\Users\pc\go
Edit PATH:
Remove C:\msys64\ucrt64\bin\go.exe
Ensure C:\Program Files\Go\bin is present before MSYS2 paths.
```
vs code terminal
```
$env:GOROOT = "C:\Program Files\Go"
$env:GOPATH = "$env:USERPROFILE\go"
$env:PATH += ";C:\Program Files\Go\bin;$env:GOPATH\bin"
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
