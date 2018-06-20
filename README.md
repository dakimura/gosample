# gosample
this repository is dakimura's example project to learn golang using ...

### env
Visual Studio Code 1.19.1
Windows 10

## debugger
it seems "delve" is a defacto-standard debugger...

### installation
To install, `go get -u github.com/derekparker/delve/cmd/dlv` failed due to the following error:
```
> go get -u github.com/derekparker/delve/cmd/dlv
# cd ***\projects\go\src\github.com\derekparker\delve; git pull --ff-only
fatal: Not a git repository (or any of the parent directories): .git
package github.com/derekparker/delve/cmd/dlv: exit status 128
```

Finally I ended up installing delve by the following commands at VSCode's terminal (PowerShell):
```
***> git clone https://github.com/derekparker/delve.git
***> go get -v github.com/golang/dep
```

## dependency management
it seems "dep" is an official dependency management tool for golang.

### installation
```
***> 
```


### reference
```
https://github.com/derekparker/delve
https://dev.classmethod.jp/go/dep/
```
